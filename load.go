package termfo

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"zgo.at/termfo/caps"
	"zgo.at/termfo/keys"
)

const (
	headerMagic    = 0o0432
	headerMagicExt = 0o01036
	headerSize     = 12
)

// Loaders is a list of additional loader callbacks.
//
// See the documentation on New() for the exact loading order.
//
// Note that access to this isn't synchronized; it usually shouldn't be needed.
//var Loaders []func(string) *Terminfo

func loadTerminfo(term string) (*Terminfo, error) {
	if term == "" {
		return nil, errors.New("TERM not set")
	}

	ti, fp, err := findTerminfo(term)
	if err != nil {
		return nil, fmt.Errorf("terminfo: %w", err)
	}
	if ti != nil {
		return ti, nil
	}

	defer fp.Close()
	return readTi(fp)
}

// See doc on New() for loading order.
func findTerminfo(term string) (*Terminfo, *os.File, error) {
	if terminfo := os.Getenv("TERMINFO"); terminfo != "" {
		return fromPath(term, terminfo)
	}

	if _, ok := os.LookupEnv("NO_BUILTIN_TERMINFO"); !ok {
		// for _, l := range Loaders {
		// 	if ti := l(term); ti != nil {
		// 		return ti, nil, nil
		// 	}
		// }
	}

	if h := os.Getenv("HOME"); h != "" {
		if _, fp, err := fromPath(term, h+"/.terminfo"); err == nil {
			return nil, fp, nil
		}
	}

	if dirs := os.Getenv("TERMINFO_DIRS"); dirs != "" {
		for _, dir := range strings.Split(dirs, ":") {
			if dir == "" {
				dir = "/usr/share/terminfo"
			}
			if _, fp, err := fromPath(term, dir); err == nil {
				return nil, fp, nil
			}
		}
	}

	if _, fp, err := fromPath(term, "/lib/terminfo"); err == nil {
		return nil, fp, nil
	}
	return fromPath(term, "/usr/share/terminfo")
}

func fromPath(term, path string) (*Terminfo, *os.File, error) {
	if _, err := os.Open(path); err != nil {
		return nil, nil, err
	}

	fp, err := os.Open(path + "/" + term[0:1] + "/" + term) // e.g. x/xterm
	if err == nil {
		return nil, fp, nil
	}

	// 68/xterm; as used on Darwin/macOS.
	fp, err = os.Open(path + "/" + hex.EncodeToString([]byte(term[:1])) + "/" + term)
	return nil, fp, err
}

func readTi(fp *os.File) (*Terminfo, error) {
	// Read the header.
	var header struct{ Magic, SizeNames, CountBools, CountNums, StrOffets, SizeTbl int16 }
	if err := binary.Read(fp, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("terminfo: reading header: %w", err)
	}

	// The regular format has 16bit numbers, the "extended number format" is
	// 32bits. It looks like tic will only compile them with 32bit numbers if
	// needed, so both are common.
	intSize := int16(2)
	switch header.Magic {
	case headerMagic:
	case headerMagicExt:
		intSize = 4
	default:
		return nil, fmt.Errorf("terminfo: unexpected magic number in header: 0o%o", header.Magic)
	}

	tiData := struct {
		names   []byte
		bools   []bool
		align   []byte
		nums    []byte // Can be 16 or 32 bit, will convert later.
		strOffs []int16
		strTbl  []byte
	}{
		make([]byte, header.SizeNames),
		make([]bool, header.CountBools),
		make([]byte, align(header.SizeNames+header.CountBools)),
		make([]byte, header.CountNums*intSize),
		make([]int16, header.StrOffets),
		make([]byte, header.SizeTbl),
	}
	err := readM(fp, &tiData.names, &tiData.bools, &tiData.align, &tiData.nums, &tiData.strOffs, &tiData.strTbl)
	if err != nil {
		return nil, fmt.Errorf("terminfo: reading data: %w", err)
	}

	// Terminal names separated by "|", with the last entry being the
	// description. Ends with NUL.
	snames := strings.Split(string(tiData.names[:len(tiData.names)-1]), "|")
	ti := &Terminfo{
		Name:     snames[0],
		Desc:     snames[len(snames)-1],
		Bools:    make(map[*caps.Cap]struct{}, 8),
		Numbers:  make(map[*caps.Cap]int32, 8),
		Strings:  make(map[*caps.Cap]string, 32),
		Keys:     make(map[string]keys.Key, len(keys.Keys)),
		IntSize:  int(intSize),
		Location: fp.Name(),
	}
	if len(snames) > 2 {
		ti.Aliases = snames[1 : len(snames)-1]
	}

	// Booleans are one byte per value.
	for i, b := range tiData.bools {
		if b {
			ti.Bools[caps.TableBools[i]] = struct{}{}
		}
	}
	// Numbers can be 16 or 32bits, depending on the header. -1 means it's not
	// present in the file.
	for i, n := range toNum(tiData.nums, int(intSize)) {
		if n > -1 {
			ti.Numbers[caps.TableNums[i]] = n
		}
	}
	// strOffs are offsets to an entry in strTbl; the table entries are ended by
	// NULL bytes. -1 means the entry is missing.
	for i, s := range tiData.strOffs {
		if s > -1 {
			ti.Strings[caps.TableStrs[i]] = string(tiData.strTbl[s : int(s)+bytes.IndexByte(tiData.strTbl[s:], 0)])
		}
	}

	// The "extended storage format" has another header after the string table,
	// which may or may not be present.

	if tell, _ := fp.Seek(0, io.SeekCurrent); tell%2 != 0 {
		fp.Read(make([]byte, 1))
	}

	var extHeader struct{ CountBools, CountNums, CountStrs, UsedStrs, SizeTbl int16 }
	if err := binary.Read(fp, binary.LittleEndian, &extHeader); err != nil {
		if errors.Is(err, io.EOF) { // No header: no problem.
			return ti, nil
		}
		return nil, fmt.Errorf("terminfo: reading extended header: %w", err)
	}
	extData := struct {
		bools   []bool
		align   []byte
		nums    []byte
		strOffs []int16
		strTbl  []byte
	}{
		make([]bool, extHeader.CountBools),
		make([]byte, align(extHeader.CountBools)),
		make([]byte, extHeader.CountNums*intSize),
		make([]int16, extHeader.UsedStrs),
		make([]byte, extHeader.SizeTbl),
	}
	if err := readM(fp, &extData.bools, &extData.align, &extData.nums, &extData.strOffs, &extData.strTbl); err != nil {
		return nil, fmt.Errorf("terminfo: reading extended data: %w", err)
	}

	// The strings table includes both string values and the names of the
	// extended capablities; CountStrs is the number of string values, UsedStrs
	// is total number of strings.
	startNames := -1
	extStrs := make([]string, 0, extHeader.CountStrs)
	for i := int16(0); i < extHeader.CountStrs; i++ {
		s := extData.strOffs[i]
		if s > -1 {
			e := int(s) + bytes.IndexByte(extData.strTbl[s:], 0)
			startNames = e
			extStrs = append(extStrs, string(extData.strTbl[s:e]))
		}
	}

	startNames++
	ti.Extended = make([]*caps.Cap, extHeader.UsedStrs-extHeader.CountStrs)
	for i := int16(0); i < extHeader.UsedStrs-extHeader.CountStrs; i++ {
		s := extData.strOffs[i+extHeader.CountStrs] + int16(startNames)
		e := int(s) + bytes.IndexByte(extData.strTbl[s:], 0)
		name := string(extData.strTbl[s:e])

		var c *caps.Cap
		// TODO: it list AX and G0 in the file, but infocmp lists it as OTbs and
		// OTpt? Hmm. Not sure where it gets that from.
		for _, v := range caps.TableStrs {
			if v.Short == name {
				c = v
				break
			}
		}
		if c == nil {
			for _, v := range caps.TableNums {
				c = v
			}
		}
		if c == nil {
			for _, v := range caps.TableBools {
				c = v
			}
		}
		if c == nil {
			c = &caps.Cap{Short: name, Long: name, Desc: "extended user-defined"}
		}
		ti.Extended[i] = c
	}

	// Don't need to check the value here, as it's never false or -1.
	for i := range extData.bools {
		ti.Bools[ti.Extended[i]] = struct{}{}
	}
	for i, n := range toNum(extData.nums, int(intSize)) {
		ti.Numbers[ti.Extended[i+len(extData.bools)]] = n
	}
	for i, s := range extStrs {
		ti.Strings[ti.Extended[i+len(extData.bools)+len(extData.nums)/int(intSize)]] = s
	}
	return ti, nil
}

// From term(5): "Between the boolean section and the number section, a null
// byte will be inserted, if necessary, to ensure that the number section begins
// on an even byte (this is a relic of the PDP-11's word-addressed architecture,
// originally designed in to avoid IOT traps induced by addressing a word on an
// odd byte boundary). All short integers are aligned on a short word boundary."
func align(n int16) int {
	if n%2 != 0 {
		return 1
	}
	return 0
}

func readM(fp *os.File, data ...any) error {
	for _, d := range data {
		if err := binary.Read(fp, binary.LittleEndian, d); err != nil {
			return err
		}
	}
	return nil
}

func toNum(read []byte, intSize int) []int32 {
	nums := make([]int32, 0, len(read)/intSize)
	for i := 0; i < len(read); i += intSize {
		n := int32(read[i]) | int32(read[i+1])<<8
		if intSize == 4 {
			n |= int32(read[i+2])<<16 | int32(read[i+3])<<24
		} else if n == 65535 { // -1 in int16; we need to add them as it's all offset based.
			n = -1
		}
		nums = append(nums, n)
	}
	return nums
}

// This adds "PC-style function keys" modifiers, as Xterm does it. When a
// modifier is used the character after the CSI is replaced with a modifier code
// or inserted before the final ~. For example (CSI prefix omitted):
//
//	            F1     F5     Up
//	Regular     OP   15~      OA
//	Ctrl      1;5P   15;5~  1;5A
//	Shift     1;2P   15;2~  1;2A
//	Alt       1;3P   15;3~  1;3A
//
// Modifier codes:
//
//	2   Shift
//	3   Alt
//	4   Shift + Alt
//	5   Ctrl
//	6   Shift + Ctrl
//	7   Alt + Ctrl
//	8   Shift + Alt + Ctrl
//
// We don't do anything with meta.
//
// You tell me why it works like this... My guess that in 19verylongago it was
// easier to do some bit banging like this on a very simple terminal (by the
// standard of the last 30 years anyway), and now we're still stuck with this.
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.pdf
//
// Anyway, pre-compute a list here so it's easier to check later. It adds about
// 0.08ms startup time, which isn't too bad, and on the upside you'll save up to
// a whopping 0.04ms latency on evey key stroke.
func addModifierKeys(ti *Terminfo, seq string, k keys.Key) {
	switch {
	case strings.HasPrefix(seq, "\x1b[") && seq[len(seq)-1] == '~':
		noTilde := seq[:len(seq)-1]
		ti.Keys[noTilde+";2~"] = k | keys.Shift
		ti.Keys[noTilde+";3~"] = k | keys.Alt
		ti.Keys[noTilde+";4~"] = k | keys.Shift | keys.Alt
		ti.Keys[noTilde+";5~"] = k | keys.Ctrl
		ti.Keys[noTilde+";6~"] = k | keys.Shift | keys.Ctrl
		ti.Keys[noTilde+";7~"] = k | keys.Ctrl | keys.Alt
		ti.Keys[noTilde+";8~"] = k | keys.Shift | keys.Ctrl | keys.Alt
	case strings.HasPrefix(seq, "\x1bO") && len(seq) == 3:
		noCSI := seq[2:]
		ti.Keys["\x1b[1;2"+noCSI] = k | keys.Shift
		ti.Keys["\x1b[1;3"+noCSI] = k | keys.Alt
		ti.Keys["\x1b[1;4"+noCSI] = k | keys.Shift | keys.Alt
		ti.Keys["\x1b[1;5"+noCSI] = k | keys.Ctrl
		ti.Keys["\x1b[1;6"+noCSI] = k | keys.Shift | keys.Ctrl
		ti.Keys["\x1b[1;7"+noCSI] = k | keys.Ctrl | keys.Alt
		ti.Keys["\x1b[1;8"+noCSI] = k | keys.Shift | keys.Ctrl | keys.Alt
	}
}
