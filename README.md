termfo is a terminfo library for Go.

It also has a little `termfo` commandline tool to list, search, and print
various terminfo things that are not so easy to get with `infocmp` and/or are
formatted a bit nicer.

Import at `zgo.at/termfo`; API docs: https://godocs.io/zgo.at/termfo

Current status: should be (mostly) usable and complete, but not widely tested
yet and there are a few rough edges here and there. Also the API may change;
specifically, I might rename some capability or key constants. For now I want to
focus on the application I *wanted* to write rather than all this stuff :-)

Usage
-----
*Note: you may want to read the "Some background and concepts" section below
first if you're not already familiar with the basics of terminals and/or
terminfo, which explains a bit of background that may be useful.*

---

First create a new `termfo.Terminfo` instance; the parameter is the terminal to
load; it will use the `TERM` environment variable if it's empty, which is what
you want >99% of the time:

```go
ti, err := termfo.New("")
```

Capabilities have three types: bool, number, and string; you can get them from
the `Bools`, `Numbers`, and `Strings` maps:

```go
s, ok := ti.Strings[caps.EnterItalicsMode]
if ok {
    fmt.Println(s, "Italic text!", ti.GetString(caps.ExitItalicsMode))
}

n, ok := ti.Numbers[cap.MaxColors]     // "max_colors" ("colors")
if ok {
    fmt.Printf("Supports %d colours\n")
}

_, ok := ti.Bools[caps.AutoRightMargin] // "am"; note this only lists values 
                                        // present, so it's always true.
if ok {
    fmt.Println("Has am")
}
```

The capabilities themselves are in the `termfo/caps` subpackage as pointers to
the `caps.Cap` struct, which also contains the short name (e.g. `sitm` for
`enter_italics_mode`) and the description from terminfo(5). If you have an
impressive unix beard and managed to memorize all the short codes then you can
use the `scaps` package:

```
s, ok := ti.Strings[scaps.Sitm]
if ok {
    fmt.Println(s, "Italic text!", ti.GetString(scaps.Ritm))
}
```

`sitm` instead of `enter_italics_mode` is just obscure, but having the mapping
is useful at times, even if only to make it easier to find out what something
does from looking at constants in C code.

To add parameters use the `Get()` method:

```go
ti.Get(caps.ParmDeleteLine, 5)   // Delete 5 lines
```

There is also `Put()` to write it, and `Supports()` to check if it's supported.

NOTE: this part of the API still sucks a bit; some of the capabilities at least
indicate they accept parameters with `Parm`, but some don't, and omitting the
argument will send nonsense (this should typed). Not yet 100% sure what a nice
API would look like. Part of the problem is that terminfo files can add
user-defined extended attributes.

### Keys
There is some additional processing for keys; the most common way to use this is
through the `Terminfo.FindKeys()` method; for example:

```go
ti, _ := termfo.New("")

ch := ti.FindKeys(os.Stdin)
for e := <-ch; ; e = <-ch {
    fmt.Println("Pressed", e.Key)
}
```

This will keep scanning for keys in stdin. Note that you'll need to put the
terminal in "raw mode" by sending the appropriate ioctls to send keys without
having to press Enter. I recommend using the `golang.org/x/term` package for
this; it's not included here as it pulls in x/sys which is ~8.5M and a somewhat
large dependency. You can also use `syscall`. See `internal/term` for an example
of that.

Keys are represented as a `Key`, which is an uint64. The lower 32 bits are used
as a regular rune, and the remaining 32 for some other information like modifier
keys. The upshot of this is that you can now use a single value to test for all
combinations:

    switch Key(0x61) {
    case 'a':                           // 'a' w/o modifiers
    case 'a' | keys.Ctrl:               // 'a' with control
    case 'a' | keys.Ctrl | keys.Shift:  // 'a' with shift and control

    case keys.Up:                       // Arrow up
    case keys.Up | keys.Ctrl:           // Arrow up with control
    }

Note that keys are always sent as lower-case; use `'a' | keys.Shift` to test for
upper-case, and control characters are always sent as `'a' | keys.Ctrl` rather
than 0x01.

### Mouse support
There is no direct support for this (yet), mostly because I simply don't need
it.

<!--
TODO: Need to actually work on this; the current method adds way too much data.

Using embed to add a terminfo database is probably better than compiling them as
Go files. The "common" list added 1.5M to the binary, which is comparatively a lot.

Compiling in terminfo definitions
---------------------------------
Use e.g. `go run ./cmd/termfo build xterm*` to generate a Go file with all
`xterm*` terminfos present on your system. Add this somewhere in your
application. This will call `termfo.SetBuiltin()`, after this `New()` will use
it.

The special name `%common` uses most common terminals in use today.

The upshot of this is that you don't need a `/usr/share/terminfo` on the system,
and it's a wee bit faster as it won't have to read anything from disk (although
this should be more than fast enough really).
-->

Updating
--------
Various terminfo data in is generated from the ncurses source with `term.h.zsh`.
This requires the ncurses source tree.

This requires zsh, awk, and gofmt.


Some background and concepts
----------------------------
A "terminfo" file is essentially a key/value database to tell applications about
the properties of the terminal the user is using.

To understand why this is needed you need to understand that terminals – and
applications that run inside them – are completely text-based. If you press the
`a` key then the terminal will send exactly the letter `a` the the application,
and nothing more. There is no such thing as a "key down" and "key up event",
it's just the byte for `a` that's sent. Special keys like F1, arrow keys, etc.
are actually multiple characters, usually starting with the 0x1b character (the
escape character, which I'll write as `\E` henceforth); for example on my system
F1 sends `\EOP` and the arrow up key sends `\EOA`.

Similarly, all *output* from applications to a terminal are also pure text. To
do something more than just "display text" we again need to use escape
sequences. For example `\E[1m` will make the text bold (until reset), or `\E[2J`
will clear the screen. Some may also send back data; for example `\E[6n` to get
the cursor position.

The reason it all works like this is because "terminals" were originally actual
devices with a screen and keyboard, connected over a serial port to a computer,
and this was the only way to send *any* data. Early version from the 60s often
had a printer rather than a screen. 

What you're using today is more accurately described as a *terminal emulator*;
that is, a program that *emulates* one of those physical devices. Back in those
days computers were very expensive (hundreds of thousands of dollars), and
terminals were comparatively cheap (though still expensive, usually several
*thousand* dollars in today's money!) I wrote a little bit more about the
history at https://bestasciitable.com, which also includes some pictures.

Now, the problem is that not every terminal (or "terminal emulator", but you can
use them interchangeably) may agree what the escape sequence is to make text
bold, or what the "F1 key" looks like. There is nothing "special" about `\E[1m`;
it's just that most terminals agree that this is the sequence to make text bold,
but it could also have been `\Ebold` or even just `BOLD` if you wanted (but that
would make it impossible to write the text "BOLD", but you could if you wanted
to).

In the past there were dozens of brands and many different terminal devices,
many of which had widely different escape sequences and logic, so people created
databases to record all of this, and an application could work with multiple
terminal devices rather than the one the author of the program was using. There
are actually two solutions for this: `termcap` and `terminfo`, both more or less
similar (terminfo has more features). Creating two different standards because
there are too many standards ... classic. These days, systems almost exclusively
use terminfo, although termcap compatibility is still provided by some systems.
There have historically been a few different implementations of terminfo, but
the one used almost universally today is the one that's part of ncurses,
maintained by Thomas Dickey who also maintains xterm. terminfo is part of POSIX
(as is curses), termcap is not.

terminfo files are usually stored in `/usr/share/terminfo`; the files in there
are "compiled" binary files. I guess parsing text was too expensive in 1981, and
the binary format stuck (including 16bit alignment for old 16bit aligned systems
like the PDP-11 by the way).

Today a lot has been standardized and converged; ECMA-48 and "Xterm-style escape
sequences" are what almost all (if not all) commonly used terminals use. This is
why you can get away with just using `printf '\x1b[1mBOLD!\x1b[0m\n'` in your
shell scripts and not worry too much about looking up the correct terminfo
properties. The *True Right Way™* to do this is still to look up the terminfo
entries though, and if you get beyond some of the basics like bold text this is
still needed. There are still several "styles" of doing some things for some
more advanced control codes (such as RGB colours, mouse support) and recognition
of "special" keys ("backspace sends delete" is a common one).

You can send this from the shell with the `tput` command, for example:

    % printf "$(tput bold)Bold!$(tput sgr0) Not bold\n"

`sgr0` is "set graphic reset" (I think?) and resets all graphical attributes.
The names for these things range from "a bit obscure" or "an indecipherable set
of letters with no obvious meaning" – party like it's UNIX. You also have long
names (`exit_attribute_mode` for `sgr0`) but `tput` doesn't recognize them.

You can see a list of the terminfo entries for your current terminal with
`infocmp -1Lx` (`-L` to use long names, `-1` to print one per line, and `-x` to
print extended data too). You can compare them too:

    % infocmp -1Lx xterm-256color iterm2
    comparing booleans.
        backspaces_with_bs: T:F.
        can_change: T:F.

    comparing strings.
        clear_screen: '\E[H\E[2J', '\E[H\E[J'.
        cursor_normal: '\E[?12l\E[?25h', '\E[?25h'.
        cursor_visible: '\E[?12;25h', NULL.

There are actually many more differences between `xterm-256color` and `iterm2`,
but I'm not going to show them all here. Note how `clear_screen` is slightly
different.

Aside from a simple key→value mapping, terminfo entries can also have
parameters. For example `parm_delete_line` (or `dl`) is `\E[%p1%dM`. The way
this works is with a small stack-based language; `%p1` pushes the first
parameter on the stack, and `%d` pops a value and prints it as a number. So with
`dl 5` we end up with `\E[5M`.

There are all sorts of things you can do with this, like conditionals and
arithmetic. This is useful because some may accept a RGB colour as hex in the
range 0x00-0xff, whereas others may want it as a decimal in the range 0-255.
Stuff like that. You don't really need to worry about this because the only
people writing these files are authors or terminal applications (or people who
write terminfo libraries). But it's fun to know that terminfo files are
Turing-complete, no?

So this is the short version on how terminals work, and what the point of
terminfo is :-) There's more to tell here (as well as another way to control
terminals, with the `ioctl()` syscall) but I'll tell that bedtime story next
week, but only if you behave and don't do anything naughty!


Others
------
Some other Go terminfo implementations I found:

- https://github.com/nsf/go-termbox  
  Very incomplete.

- https://github.com/gdamore/tcell  
  Incomplete, relies on `infocmp` for non-builtin terminals (yikes), and I
  didn't care much for the API.

- https://github.com/eiannone/keyboard  
  Copied from go-termbox, even more incomplete.

- https://github.com/charithe/terminfo  
  This actually seems pretty decent, although it lacks the key scanning part. If
  I had known about it I would have forked/used it. Ah well...

- https://github.com/xo/terminfo  
  Fork of the above; adds a lot of ncurses-y TUI stuff.

Some of these other packages (such as termbox and tcell) also do much more than
just dealing with terminfo. This package is intended to *only* support doing
useful things with terminfo and not much more. A big advantage is that it's a
lot easier to use in simpler CLI apps that are not full-blown TUIs.
