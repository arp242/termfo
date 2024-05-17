//go:build ignore

package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// ######## APPLE → #### Terminal.app → ...
	var (
		terms         = make(map[string]map[string][]string)
		sect, subsect string
	)
	for _, l := range strings.Split(string(f), "\n") {
		if strings.HasPrefix(l, "########") {
			sect = strings.TrimLeft(l, "# ")
			terms[sect] = make(map[string][]string)
			continue
		}
		if strings.HasPrefix(l, "####") && !strings.Contains(l, "TERMINFO FILE CAN BE SPLIT HERE") {
			subsect = strings.TrimLeft(l, "# ")
			terms[sect][subsect] = []string{}
			continue
		}
		if !strings.HasPrefix(l, "\t") && strings.HasSuffix(l, ",") && strings.Contains(l, "|") {
			t := strings.Split(l, "|")
			terms[sect][subsect] = append(terms[sect][subsect], t[:len(t)-1]...)
		}
	}

	// for k, v := range terms {
	// 	fmt.Println(k)
	// 	for k2, v2 := range v {
	// 		fmt.Println("   ", k2)
	// 		fmt.Println("       ", v2)
	// 	}
	// }
	// return

	fmt.Println("package main\n\nvar oldterm = map[string]int{")
	dopr := func(k string, list []string) {
		fmt.Printf("\t// %s\n\t", k)
		for _, v := range list {
			fmt.Printf("%q: 1,", v)
		}
		fmt.Println()
	}
	only := func(list []string, incl ...string) []string {
		n := make([]string, 0, 16)
		for _, v := range list {
			for _, i := range incl {
				if strings.HasPrefix(v, i) {
					n = append(n, v)
					continue
				}
			}
		}
		return n
	}
	for _, k := range ordered(terms) {
		v := terms[k]
		switch k {
		// Pretty much all old hardware terminals.
		case "OLDER TERMINAL TYPES", "OTHER OBSOLETE TYPES", "LCD DISPLAYS", "OBSOLETE VDT TYPES",
			"COMMON TERMINAL TYPES": // Common hardware terminals in 1990
			fmt.Printf("\t// %s\n", k)
			for _, k2 := range ordered(v) {
				dopr(k2, v[k2])
			}

		// All of these are very old and/or for outdated versions.
		case "Non-Unix Consoles", "NON-ANSI TERMINAL EMULATIONS", "OBSOLETE UNIX CONSOLES",
			"COMMERCIAL WORKSTATION CONSOLES", "OBSOLETE PERSONAL-MICRO CONSOLES AND EMULATIONS":
			fmt.Printf("\t// %s\n", k)
			for _, k2 := range ordered(v) {
				dopr(k2, v[k2])
			}
		case "APPLE":
			for _, k2 := range ordered(v) {
				fmt.Printf("\t// %s\n", k)
				switch k2 {
				case "Terminal.app":
					dopr(k2, only(v[k2], "nsterm-build", "nsterm+", "nsterm-7", "nsterm-old"))
				// xnuppc* and darwin* is old PowerPC system console.
				case "iTerm, iTerm2":
					dopr(k2, only(v[k2], "xnuppc", "darwin"))
				}
			}
		case "UNIX VIRTUAL TERMINALS, VIRTUAL CONSOLES, AND TELNET CLIENTS":
			fmt.Printf("\t// %s\n", k)
			for _, k2 := range ordered(v) {
				switch k2 {
				// Old Palm pilot, telnet
				case "NCSA Telnet", "Pilot Pro Palm-Top":
					dopr(k2, v[k2])
				case "Screen":
					dopr(k2, only(v[k2], "screen2", "screen3", "screen4", "screen5")) // Old versions
				}
			}
			dopr("CB UNIX, early 80s", []string{"cbunix", "vremote", "pty"})
		case "X TERMINAL EMULATORS":
			fmt.Printf("\t// %s\n", k)
			for _, k2 := range ordered(v) {
				switch k2 {
				// "Multi Gnome Terminal", hasn't been updated in 20 years
				case "Other GNOME":
					dopr(k2, only(v[k2], "mgt"))
				// Old pre-VTE gnome-terminal, and specific gnome-terminal/VTE versions
				case "GNOME (VTE)":
					dopr(k2, only(v[k2], "gnome",
						"vte-2007", "vte-2008", "vte-2012", "vte-2014", "vte-2017", "vte-2018"))
				case "SIMPLETERM":
					dopr(k2, only(v[k2], "st-0.", "old-st", "simpleterm")) // Old versions
				case "TERMINOLOGY":
					dopr(k2, only(v[k2], "terminology-0", "terminology-1")) // Old versions
				case "MLTERM":
					dopr(k2, only(v[k2], "mlterm3", "mlterm2")) // Old versions
				case "KDE":
					dopr(k2, only(v[k2],
						// Old konsole name (<2001).
						"kvt",
						// Old seemingly obselete konsole variants.
						"konsole-xf"))
				case "XTERM":
					// Old versions
					dopr(k2, only(v[k2], "x10term", "xterm-r5", "xterm-r6", "xterm-old", "xterm-xf86", "xterm-xfree86",
						"xterm-24", "vs100", "xterms", "xterm-new"))
				case "XTERM Features":
					dopr(k2, only(v[k2], "xterm-8bit", "xterm-vt52", "xterm1", "xterm-nic"))
				case "Other XTERM":
					dopr(k2, only(v[k2],
						// Old xterm fork from before xterm supported ECMA-64 colors
						"color_xterm", "cx",
						// Weird terminal with weird overrides.
						"xterm-pcolor",
						// xterm as distributes with late 90s Solaris.
						"xtermc", "xtermm",
					))
				case
					// Old Xfree86 term, or something.
					"EMU",
					// MGR is old Bell Labs thingy
					// https://en.wikipedia.org/wiki/ManaGeR
					// https://hack.org/mc/mgr/
					"MGR",
					// Seems to be some old (non-ANSI) emulator? Can't find much about it.
					"MTERM",
					// Seems to be (very) old version of "Rocket® Terminal Emulator"
					// https://www.rocketsoftware.com/products/rocket-terminal-emulator/rocket-terminal-emulation
					"MVTERM",
					// Old versions of https://github.com/TragicWarrior/vwm
					"VWM",
					// RXVT hasn't been updated in >20 years; no one ships it
					// any more (e.g. on Debian it links to rxvt-unicode).
					"RXVT",
					// Old emulator from 90s with support for Kanji
					// https://ja.wikipedia.org/wiki/Kterm
					"KTERM",
					// HP-UX hpterm terminal emulator; I'm not really sure if
					// this is still in use on modern HP-UX systems. Even if it
					// is: it's pretty obscure.
					"HPTERM":
					dopr(k2, v[k2])
				}
			}
		case "ANSI, UNIX CONSOLE, AND SPECIAL TYPES":
			fmt.Printf("\t// %s\n", k)
			for _, k2 := range ordered(v) {
				switch k2 {
				// Only include vt220; no one really targets vt100, or vt400, or
				// whatnot.
				case "DEC VT100 and compatibles":
					dopr(k2, only(v[k2], "vt100", "vt102", "vt125", "vt13", "vt200",
						"vt220-old", "vt220-8", "vt220d", "vt3", "vt4", "vt5", "dec-"))
				// Old NetBSD system consoles
				case "NetBSD consoles":
					dopr(k2, only(v[k2], "pcvt", "arm100", "x68k", "ofcons", "rcons", "mgterm", "netbsd6"))
				case "ANSI.SYS/ISO 6429/ECMA-48 Capabilities":
					dopr(k2, only(v[k2],
						"ibcs2", // iBCS is some failed Intel standard
					))
				case "ANSI/ECMA-48 terminals and terminal emulators":
					dopr(k2, only(v[k2],
						"ansi77", // 1977 ANSI
					))
				case
					"Specials",
					"DEC VT52",
					"VT100 emulations",
					"SCO consoles", "SGI consoles", "386BSD and BSD/OS Consoles", // Old Unix
					"Atari ST terminals",    // Atari ST (1985)
					"Mach",                  // No one uses this.
					"BeOS",                  // Don't think this is used?
					"DOS ANSI.SYS variants", // Old DOS; not used I think.
					"QNX":                   // QNX 4 (1990); probably not used on modern QNX.
					dopr(k2, v[k2])
				}
			}
		case "DOS/WINDOWS":
			fmt.Printf("\t// %s\n", k)
			for _, k2 := range ordered(v) {
				switch k2 {
				case "Command prompt":
					dopr(k2, only(v[k2], "vt100+", "ms-vt100", "ms-vt-", "vtnt", "vt-utf")) // WinNT 4
				case "PuTTY":
					dopr(k2, only(v[k2], "vt100-putty", "putty-vt100", "putty+fnkeys+vt100", "putty+fnkeys+vt400"))
				}
			}
			// CRT is very old outdated version of SecureCRT
			dopr("CRT is very old outdated version of SecureCRT", []string{"crt", "crt-vt220"})
		}
	}
	fmt.Println("}")
}

func ordered[M ~map[K]V, K cmp.Ordered, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	slices.Sort(r)
	return r
}
