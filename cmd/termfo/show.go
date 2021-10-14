package main

import (
	"fmt"
	"sort"
	"strings"

	"zgo.at/termfo"
	"zgo.at/termfo/caps"
)

func show(terms ...string) {
	for _, t := range terms {
		ti, err := termfo.New(t)
		if err != nil {
			fatalf("%s", err)
		}
		fmt.Println(fmtTerminfo(ti))
	}
}

func fmtTerminfo(ti *termfo.Terminfo) string {
	all := append(append(append(caps.TableBools, caps.TableNums...), caps.TableStrs...), ti.Extended...)
	sort.Slice(all, func(i, j int) bool { return all[i].Long < all[j].Long })

	// Highlight escape codes and such; makes it easier to read.
	hi := func(s string) string { return s }
	if isTerm {
		hi = func(s string) string {
			r := make([]byte, 0, len(s))
			resetAfter := 0
			for i, c := range []byte(s) {
				if c == '\\' {
					r = append(r, "\x1b[34m"...)
					resetAfter = 1
					if len(s) > i+1 && s[i+1] == 'x' {
						resetAfter = 3
					}
				}

				r = append(r, c)
				if resetAfter > -1 {
					resetAfter--
					if resetAfter == -1 {
						r = append(r, "\x1b[0m"...)
					}
				}
			}
			return string(r)
		}
	}

	b := new(strings.Builder)
	b.Grow(16384)
	fmt.Fprintf(b, "Loaded from %s\n", ti.Location)
	a := ""
	if len(ti.Aliases) > 0 {
		a = " (aliases: " + strings.Join(ti.Aliases, ", ") + ")"
	}
	fmt.Fprintf(b, "%s%s – %s\n\n", ti.Name, a, ti.Desc)

	fmt.Fprintf(b, "%-8s │ %-26s │ %-26s │ Description\n", "Short", "Long", "Value")
	fmt.Fprintf(b, "%s┼%s┼%s┼%s\n",
		strings.Repeat("─", 9),
		strings.Repeat("─", 28),
		strings.Repeat("─", 28),
		strings.Repeat("─", 22),
	)
	for _, k := range all {
		var (
			val string
		)
		if _, ok := ti.Bools[k]; ok {
			val = "True"
		} else if v, ok := ti.Numbers[k]; ok {
			val = fmt.Sprintf("#%d", v)
		} else if v, ok := ti.Strings[k]; ok {
			val = strings.ReplaceAll(fmt.Sprintf("%#v", v), `\x1b`, `\E`)
			val = val[1 : len(val)-1]
		}

		if val != "" {
			var overflow string
			if isTerm && len(val) > 26 {
				overflow = val[26:]
				val = val[:26]
			}

			// TODO: if it overflows at %\np then the "p" isn't highlighted
			// (this is also why that reset is in there).
			reset := ""
			if isTerm {
				reset = "\x1b[0m"
			}
			fmt.Fprintf(b, "%-8s │ %-26s │ %s%s │ %s\n", k.Short, k.Long, hi(fmt.Sprintf("%-26s", val)), reset, k.Desc)
			for p, overflow := first(overflow, 26); p != ""; p, overflow = first(overflow, 26) {
				fmt.Fprintf(b, "%-37s │ %s │\n", " ", hi(fmt.Sprintf("%-26s", p)))
			}
		}
	}

	// Mostly because GNU cut is stuck in 1990:
	//
	//   % termfo show|cat | cut -d '│' -f3
	//   cut: the delimiter must be a single character
	if !isTerm {
		return strings.ReplaceAll(b.String(), "│", "|")
	}
	return b.String()
}

func first(s string, n int) (string, string) {
	if len(s) > n {
		return s[:n], s[n:]
	}
	return s, ""
}
