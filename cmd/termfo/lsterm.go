package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"zgo.at/termfo"
)

func lsTerm(match string) {
	fmt.Printf("%-20s │ %-42s │ %s\n", "Name", "Data", "Aliases")
	fmt.Printf("%s┼%s┼%s\n",
		strings.Repeat("─", 21),
		strings.Repeat("─", 44),
		strings.Repeat("─", 30))

	for _, a := range allInfos(match) {
		if m, _ := filepath.Match(match, a); match != "" && !m {
			continue
		}
		ti, err := termfo.New(a)
		if err != nil {
			fatalf("error reading %q: %s", a, err)
		}

		//fmt.Printf("%-20s %-40s │ %s\n", ti.Name, strings.Join(ti.Aliases, ", "), ti.Desc)
		fmt.Printf("%-20s │ bool: %2d num: %2d str: %3d key: %3d ext: %2d │ %s",
			ti.Name, len(ti.Bools), len(ti.Numbers), len(ti.Strings), len(ti.Keys), len(ti.Extended),
			strings.Join(ti.Aliases, ", "))

		fmt.Println()
	}
}

func allInfos(match string) []string {
	if _, err := os.Open("/usr/share/terminfo"); os.IsNotExist(err) {
		fatalf("/usr/share/terminfo doesn't exist")
	}

	all, err := filepath.Glob("/usr/share/terminfo/*/*")
	if err != nil {
		fatalf("reading \"/usr/share/terminfo\" %s", err)
	}

	all2 := make([]string, 0, len(all))
	for _, a := range all {
		a = filepath.Base(a)
		if !strings.ContainsRune(a, '.') {
			all2 = append(all2, a)
		}
	}
	return all2
}
