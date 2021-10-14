package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	"zgo.at/termfo"
	"zgo.at/termfo/caps"
	"zgo.at/termfo/cmd/termfo/internal/term"
	"zgo.at/termfo/keys"
)

func keyscan() {
	ti, err := termfo.New("")
	if err != nil {
		fatalf("%s", err)
	}

	restore, err := term.MakeRaw()
	if err != nil {
		fatalf("%s", err)
	}
	defer restore()

	fmt.Printf("Loaded terminfo for %q\r\n", ti.Name)
	fmt.Print("<C-k> to list key sequences we know about; <C-t> to print terminfo; <C-c> to quit\r\n\n")

	fmt.Printf("%-48s │ %-20s\r\n", "Received", "Sent on channel")
	fmt.Printf("%s┼%s\r\n", strings.Repeat("─", 49), strings.Repeat("─", 50))

	ch := ti.FindKeys(os.Stdin)
	for e := <-ch; ; e = <-ch {
		if e.Err != nil {
			fatalf("keyscan error: %s", e.Err)
		}

		showKey(e)

		switch e.Key {
		case 'c' | keys.Ctrl:
			return
		case 't' | keys.Ctrl:
			more(ti, ch, fmtTerminfo(ti))
		case 'k' | keys.Ctrl:
			more(ti, ch, listKeys(ti))
		}
	}
}

// A very simple pager.
func more(ti *termfo.Terminfo, ch <-chan termfo.Event, s string) {
	if height <= 0 {
		fmt.Print(strings.ReplaceAll(s, "\n", "\r\n"))
		return
	}

	lines := strings.Split(s, "\n")
	for i, l := range lines {
		fmt.Print(l, "\r\n")
		if i > 0 && (i+1)%(height-1) == 0 {
			fmt.Printf("%s--- Displaying %d-%d of %d. Press any key to show more or <C-C> to abort%s\r",
				ti.Strings[caps.EnterStandoutMode],
				i-height+3, i+3, len(lines),
				ti.Strings[caps.ExitStandoutMode])
			if kk := <-ch; kk.Key == 'c'|keys.Ctrl {
				fmt.Print("\r", ti.Strings[caps.ClrEol])
				return
			}
		}
	}
}

func showKey(e termfo.Event) {
	fmt.Printf("Pressed %-40s │ → Key %-12s  0x%02x",
		fmt.Sprintf("%-10s  0x% 2x", prSeq(string(e.Seq)), e.Seq), e.Key, rune(e.Key))
	if e.Key.Shift() {
		fmt.Print("|Shift")
	}
	if e.Key.Ctrl() {
		fmt.Print("|Ctrl")
	}
	if e.Key.Alt() {
		fmt.Print("|Alt")
	}
	if kk := e.Key.WithoutMods(); kk < unicode.MaxRune {
		fmt.Printf("  U+%04X", rune(kk))
	}
	fmt.Print("\r\n")
}

func listKeys(ti *termfo.Terminfo) string {
	all := make([]string, 0, len(ti.Keys))
	for seq, k2 := range ti.Keys {
		all = append(all, fmt.Sprintf("    %-22s → %-20s", prSeq(seq), k2))
	}
	sort.Strings(all)

	b := new(strings.Builder)
	fmt.Fprintf(b, "%d keys; note that not all of these may be supported by your terminal (especially multiple modifiers)\r\n", len(all))

	w := width - 10
	p := 0
	for _, s := range all {
		b.WriteString(s)
		p += len(s)
		if p > w {
			b.WriteString("\r\n")
			p = 0
		}
	}
	b.WriteString("\r\n")
	return b.String()
}

func prSeq(s string) string {
	return strings.ReplaceAll(fmt.Sprintf("%q", s), `\x1b`, `\E`)
}
