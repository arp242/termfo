package main

import (
	"fmt"
	"sort"
	"strings"

	"zgo.at/termfo"
	"zgo.at/termfo/caps"
)

func termWithCap(name string) {
	c := findCap(name)
	if c == nil {
		fatalf("no cap: %q", name)
	}

	list := make(map[string][]string)
	for _, a := range allInfos("") {
		ti, err := termfo.New(a)
		if err != nil {
			fatalf(a, "\t", err)
			continue
		}

		if v, ok := findCapInTi(ti, c); ok {
			list[v] = append(list[v], a)
		}
	}
	type termWithCap struct {
		cap   string
		terms []string
	}
	order := make([]termWithCap, 0, len(list))
	pad := 0
	for k, v := range list {
		k = fmt.Sprintf("%q", k)
		order = append(order, termWithCap{
			cap:   k,
			terms: dedup(v),
		})
		if len(k) > pad {
			pad = len(k)
		}
	}
	sort.Slice(order, func(i, j int) bool { return len(order[i].terms) > len(order[j].terms) })

	w := 100
	pad += 4
	padf := fmt.Sprintf("%%-%ds", pad)

	fmt.Printf("Capability %q / %q: %s\n\n", c.Short, c.Long, c.Desc)
	for _, o := range order {
		p := o.terms[0]
		l := len(p)
		for _, vv := range o.terms[1:] {
			if l > w {
				p += "\n" + strings.Repeat(" ", pad+7) + vv
				l = len(vv) + 2
			} else {
				p += ", " + vv
				l += len(vv) + 2
			}
		}
		fmt.Printf(padf+" → %3d %s\n", o.cap, len(o.terms), p)
		if len(o.terms) > 10 {
			fmt.Println()
		}
	}
}

func findCap(name string) *caps.Cap {
	for _, c := range caps.TableBools {
		if c.Short == name || c.Long == name {
			return c
		}
	}
	for _, c := range caps.TableNums {
		if c.Short == name || c.Long == name {
			return c
		}
	}
	for _, c := range caps.TableStrs {
		if c.Short == name || c.Long == name {
			return c
		}
	}
	return nil
}

func findCapInTi(ti *termfo.Terminfo, c *caps.Cap) (string, bool) {
	v, ok := ti.Strings[c]
	if ok {
		return v, ok
	}
	v2, ok := ti.Numbers[c]
	if ok {
		return fmt.Sprintf("%v", v2), ok
	}

	v3, ok := ti.Bools[c]
	return fmt.Sprintf("%v", v3), ok
}

// List xterm, xterm-256color, xterm-mono, etc. as xterm*; no need to list them
// all one-by-one  here.
func dedup(list []string) []string {
	dup := make(map[string][]string)
	for _, l := range list {
		i := strings.IndexAny(l, "-+")
		if i > -1 {
			dup[l[:i]] = append(dup[l[:i]], l)
		} else {
			dup[l] = append(dup[l], l)
		}
	}

	type o struct {
		t string
		l int
	}
	order := make([]o, 0, len(dup))
	for k, v := range dup {
		order = append(order, o{k, len(v)})
	}
	sort.Slice(order, func(i, j int) bool { return order[i].l > order[j].l })

	ret := make([]string, 0, len(dup))
	for _, v := range order {
		if v.l == 1 {
			ret = append(ret, v.t)
		} else {
			ret = append(ret, fmt.Sprintf("%s* (%d)", v.t, v.l))
		}
	}
	return ret
}
