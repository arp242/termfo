//go:generate zsh term.h.zsh

package termfo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"zgo.at/termfo/caps"
	"zgo.at/termfo/keys"
)

// Terminfo describes the terminfo database for a single terminal.
type Terminfo struct {
	Name     string   // Main name as listed in the terminfo file.
	Desc     string   // Some textual description.
	Aliases  []string // Aliases for this terminal.
	Location string   // Where it was loaded from; path or "builtin".

	Bools   map[*caps.Cap]struct{} // Boolean capabilities.
	Numbers map[*caps.Cap]int32    // Number capabilities.
	Strings map[*caps.Cap]string   // String capabilities.

	// Capabilities listed in the "extended" section. The values are in the
	// Bools, Numbers, and Strings maps.
	Extended []*caps.Cap

	// The default format uses int16, but the "extended number format" uses
	// int32. This lists the integer size as 2 or 4.
	IntSize int

	// List of keys, as sequence → Key mapping. e.g. "\x1b[OP" → KeyF1.
	//
	// This contains all key_* capabilities, plus a few generated ones for
	// modifier keys and such.
	Keys map[string]keys.Key
}

// New reads the terminfo for term. If term is an empty string then the value of
// the TERM environment variable is used.
//
// It tries to load a terminfo file according to these rules:
//
//   1. Use the path in TERMINFO if it's set and don't search any other
//      locations.
//
//   2. Try built-in ones unless set NO_BUILTIN_TERMINFO is set.
//
//   3. Try ~/.terminfo/ as the database path.
//
//   4. Look in the paths listed in TERMINFO_DIRS.
//
//   5. Look in /lib/terminfo/
//
//   6. Look in /usr/share/terminfo/
//
// These are the same rules as ncurses, except that step 2 was added.
//
// TODO: curses allows setting a different path at compile-time; we can use
// infocmp -D to get this. Probably want to add this as step 7(?)
func New(term string) (*Terminfo, error) {
	if term == "" {
		term = os.Getenv("TERM")
		if term == "" {
			return nil, errors.New("terminfo: TERM not set")
		}
	}
	ti, err := loadTerminfo(term)
	if err != nil {
		return nil, err
	}

	// Add all the keys.
	for o, k := range keys.Keys {
		seq, ok := ti.Strings[o]
		if ok {
			ti.Keys[seq] = k
		}
	}

	// From tcell:
	//
	//   Sadly, xterm handling of keycodes is somewhat erratic. In particular,
	//   different codes are sent depending on application mode is in use or
	//   not, and the entries for many of these are simply absent from terminfo
	//   on many systems. So we insert a number of escape sequences if they are
	//   not already used, in order to have the widest correct usage. Note that
	//   prepareKey will not inject codes if the escape sequence is already
	//   known. We also only do this for terminals that have the application
	//   mode present.
	if _, ok := ti.Strings[caps.KeypadXmit]; ok {
		ti.Keys["\x1b[A"] = keys.Up
		ti.Keys["\x1b[B"] = keys.Down
		ti.Keys["\x1b[C"] = keys.Right
		ti.Keys["\x1b[D"] = keys.Left
		ti.Keys["\x1b[F"] = keys.End
		ti.Keys["\x1b[H"] = keys.Home
		ti.Keys["\x1b[3~"] = keys.Delete
		ti.Keys["\x1b[1~"] = keys.Home
		ti.Keys["\x1b[4~"] = keys.End
		ti.Keys["\x1b[5~"] = keys.PageUp
		ti.Keys["\x1b[6~"] = keys.PageDown
		// Application mode
		ti.Keys["\x1bOA"] = keys.Up
		ti.Keys["\x1bOB"] = keys.Down
		ti.Keys["\x1bOC"] = keys.Right
		ti.Keys["\x1bOD"] = keys.Left
		ti.Keys["\x1bOH"] = keys.Home
	}
	for seq, k := range ti.Keys {
		addModifierKeys(ti, seq, k)
	}
	return ti, nil
}

func (ti Terminfo) String() string {
	return fmt.Sprintf("Terminfo file for %q from %q with %d properties", ti.Name, ti.Location,
		len(ti.Bools)+len(ti.Numbers)+len(ti.Strings))
}

// Supports reports if this terminal supports the given capability.
func (ti Terminfo) Supports(c *caps.Cap) bool {
	if _, ok := ti.Bools[c]; ok {
		return true
	}
	if _, ok := ti.Numbers[c]; ok {
		return true
	}
	if v := ti.Strings[c]; v != "" {
		return true
	}

	return false
}

// Get a capability.
func (ti Terminfo) Get(c *caps.Cap, args ...int) string {
	v, ok := ti.Strings[c]
	if !ok {
		return ""
	}
	return replaceParams(v, args...)
}

func (ti Terminfo) Put(w io.Writer, c *caps.Cap, args ...int) {
	w.Write([]byte(ti.Get(c, args...)))
}

// Event sent by FindKeys.
type Event struct {
	Key keys.Key // Processed key that was pressed.
	Seq []byte   // Unprocessed text; only usedful for debugging really.
	Err error    // Error; only set for read errors.
}

// FindKeys finds all keys in the given reader (usually stdin) and sends them in
// the channel.
//
// Any read error will send an Event with Err set and it will stop reading keys.
func (ti Terminfo) FindKeys(fp io.Reader) <-chan Event {
	var (
		ch   = make(chan Event)
		pbuf []byte
	)
	go func() {
		for {
			buf := make([]byte, 32)
			n, err := fp.Read(buf)
			if err != nil {
				ch <- Event{Err: err}
				break
			}
			buf = buf[:n]
			if pbuf != nil {
				buf = append(pbuf, buf...)
				pbuf = nil
			}

			for {
				k, n := ti.FindKey(buf)
				if n == 0 {
					break
				}

				// Possible the buffer just ran out in the middle of a multibyte
				// character, so try again.
				if k == utf8.RuneError && len(buf) < 4 {
					pbuf = buf
					break
				}

				seq := buf[:n]
				buf = buf[n:]
				ch <- Event{Key: k, Seq: seq}
			}
		}
	}()
	return ch
}

// Find the first valid keypress in s.
//
// Returns the key and number of bytes processed. On errors it will return
// UnknownSequence and the length of the string.
func (ti Terminfo) FindKey(b []byte) (keys.Key, int) {
	// TODO: this only works for ASCII; not entirely sure how non-ASCII input is
	//       done wrt. Control key etc.
	// TODO: doesn't deal with characters consisting of multiple codepoints.
	//       Maybe want to add: https://github.com/arp242/termtext
	//
	// TODO: on my system <C-Tab> sends \E[Z, which isn't recognized(?)
	//       Also: <C-End>
	//       <S-End> is "<Send>"?
	if len(b) == 0 {
		return 0, 0
	}

	// No escape sequence.
	if b[0] != 0x1b {
		// TODO: we probably want to use rivo/uniseg here; otherwise something
		// like:
		//
		//     U+1F3F4 (🏴) U+200D U+2620 (☠️)
		//
		// will be sent as three characters, rather than one. It should be the
		// "pirate flag" emoji.
		//
		// Actually, this kinda sucks because this is where my clever "encode
		// everything in a uint64!"-scheme kind of breaks down, since this can't
		// be represented by that.
		//
		// Perhaps change the return signature to (Key, string, int), and then
		// send a special MultiCodepoint as the Key? I don't know...
		//
		// Or maybe just don't support it. Many applications will work just fine
		// anyway; e.g. if you print text it will just output those three bytes
		// and it's all grand, and modifiers aren't sent with that in the first
		// place.
		r, n := utf8.DecodeRune(b)

		switch {
		case r == 0x7f:
			return keys.Backspace, n
		case r == 0x0d:
			return keys.Enter, n
		case r < 0x1f:
			return keys.Key(r) | 0x20 | 0x40 | keys.Ctrl, n
		case r >= 'A' && r <= 'Z':
			return keys.Key(r) ^ 0x20 | keys.Shift, n
		default:
			return keys.Key(r), n
		}
	}

	// Exact match.
	k, ok := ti.Keys[string(b)]
	if ok {
		return k, len(b)
	}

	// Find first matching.
	for seq, k := range ti.Keys {
		if bytes.HasPrefix(b, []byte(seq)) {
			return k, len(seq)
		}
	}

	// Alt keys are sent as \Ek.
	// TODO: I think this depends on the "mode"? Xterm has settings for it
	// anyway.
	if len(b) == 2 {
		return keys.Key(b[1]) | keys.Alt, 2
	}
	return 0, len(b)
	return keys.UnknownSequence, len(b)
}
