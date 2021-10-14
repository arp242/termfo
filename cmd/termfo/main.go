// Program termfo prints information about the terminfo database.
package main

import (
	"fmt"
	"os"

	"zgo.at/termfo/cmd/termfo/internal/term"
)

const usage = `Usage: termfo [cmd]

Display information about terminfo files.

Commands:
    ls-term   [glob]           List terminals/terminfos.
    ls-cap    [glob]           List capabilities.
    show      [term] [term...] Show terminfo for term; defaults to TERM if none given.
                               Use "all" to list all terminfo files.
    find-cap  [cap cap..]      Show all terminals with these capabilities.
    keyscan                    Scan for keys.

If [glob] is omitted for the ls commands then it will just everything we know
about.
`

// build     [pkg] [terms..]  Generate a Go file for package [pkg] with terminals to compile in.
//                            Use '%common' for a list of common terminals in use today.

func fatalf(f string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", a...)
	os.Exit(1)
}

var (
	width, height = term.Size()
	isTerm        = term.IsTerminal()
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print(usage)
		os.Exit(0)
	}

	switch os.Args[1] {
	default:
		fatalf("unknown command: %q", os.Args[1])
	case "help", "-h", "-help", "--help":
		fmt.Print(usage)
		os.Exit(0)
	case "ls-term":
		m := ""
		if len(os.Args) > 2 {
			m = os.Args[2]
		}
		lsTerm(m)
	case "ls-cap":
		m := ""
		if len(os.Args) > 2 {
			m = os.Args[2]
		}
		lsCap(m)
	case "show":
		if len(os.Args) > 2 && os.Args[2] == "all" {
			show(allInfos("*")...)
		} else if len(os.Args) > 2 {
			show(os.Args[2:]...)
		} else {
			show(os.Getenv("TERM"))
		}
	case "find-cap":
		if len(os.Args) <= 2 {
			fatalf("need a cap name")
		}
		termWithCap(os.Args[2])
	//case "build":
	//	if len(os.Args) <= 2 {
	//		fatalf("need a package name to use")
	//	}
	//	if len(os.Args) <= 3 {
	//		fatalf("need at least one terminal name")
	//	}
	//	build(os.Args[2], os.Args[3:]...)
	case "keyscan":
		keyscan()
	}
}
