package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"zgo.at/termfo/caps"
)

func lsCap(match string) {
	all := append(append(caps.TableBools, caps.TableNums...), caps.TableStrs...)
	sort.Slice(all, func(i, j int) bool { return all[i].Long < all[j].Long })

	for _, c := range all {
		if match != "" {
			m1, _ := filepath.Match(match, c.Short)
			m2, _ := filepath.Match(match, c.Long)
			if !m1 && !m2 {
				continue
			}
		}

		fmt.Printf("%-30s %-12s %s", c.Long, c.Short, c.Desc)
		fmt.Println()
	}
}
