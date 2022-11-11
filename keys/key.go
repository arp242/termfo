package keys

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Modifiers keys.
const (
	Shift = 1 << 63
	Ctrl  = 1 << 62
	Alt   = 1 << 61

	modmask = Shift | Ctrl | Alt
)

// Common control sequences better known by their name than letter+Ctrl
// combination.
const (
	Tab    = 'i' | Ctrl
	Escape = '[' | Ctrl
)

// Key represents a keypress. This is formatted as follows:
//
//   - First 32 bits   → rune (int32)
//   - Next 16 bits    → Named key constant.
//   - Bits 49-61      → Currently unused.
//
// And the last three bits are flags for modifier keys:
//
//   - bit 62          → Alt
//   - bit 63          → Ctrl
//   - bit 64          → Shift
//
// The upshot of this is that you can now use a single value to test for all
// combinations:
//
//	switch Key(0x61) {
//	case 'a':                           // 'a' w/o modifiers
//	case 'a' | keys.Ctrl:               // 'a' with control
//	case 'a' | keys.Ctrl | keys.Shift:  // 'a' with shift and control
//
//	case keys.Up:                       // Arrow up
//	case keys.Up | keys.Ctrl:           // Arrow up with control
//	}
//
// Which is nicer than using two or three different variables to signal various
// things.
//
// Letters are always in lower-case; the keys.Shift to test for upper case.
//
// Control sequences (0x00-0x1f, 0x1f), are used sent as key + Ctrl. So this:
//
//	switch k {
//	case 0x09:
//	}
//
// Won't work. you need to use:
//
//	switch k {
//	case 'i' | key.Ctrl:
//	}
//
// Or better yet:
//
//	ti := termfo.New("")
//
//	...
//
//	switch k {
//	case ti.Keys[keys.Tab]:
//	}
//
// This just makes more sense, because people press "<C-a>" not "Start of
// heading".
//
// It's better to use the variables from the terminfo, in case it's something
// different. Especially with things like Shift and Ctrl modifiers this can
// differ.
//
// Note that support for multiple modifier keys is flaky across terminals.
type Key uint64

// Shift reports if the Shift modifier is set.
func (k Key) Shift() bool { return k&Shift != 0 }

// Ctrl reports if the Ctrl modifier is set.
func (k Key) Ctrl() bool { return k&Ctrl != 0 }

// Alt reports if the Alt modifier is set.
func (k Key) Alt() bool { return k&Alt != 0 }

// WithoutMods returns a copy of k without any modifier keys set.
func (k Key) WithoutMods() Key { return k &^ modmask }

// Valid reports if this key is valid.
func (k Key) Valid() bool { return k&^modmask <= 1<<31 || k.Named() }

// Named reports if this is a named key.
func (k Key) Named() bool {
	_, ok := keyNames[k&^modmask]
	return ok
}

// Name gets the key name. This doesn't show if any modifiers are set; use
// String() for that.
func (k Key) Name() string {
	k &^= modmask

	n, ok := keyNames[k]
	if ok {
		return n
	}
	if !k.Valid() {
		return fmt.Sprintf("Unknown key: 0x%x", uint64(k))
	}
	if rune(k) == utf8.RuneError {
		return fmt.Sprintf("Invalid UTF-8: 0x%x", rune(k))
	}

	// TODO: maybe also other spaces like nbsp etc?
	switch k {
	case ' ':
		return "Space"
	case '\t':
		return "Tab"
	case Escape:
		return "Esc"
	}
	return fmt.Sprintf("%c", rune(k))
}

func (k Key) String() string {
	var (
		hasMod = k.Ctrl() || k.Shift() || k.Alt()
		name   = k.Name()
		named  = utf8.RuneCountInString(name) > 1
		b      strings.Builder
	)

	b.Grow(8)

	if hasMod || named {
		b.WriteRune('<')
	}

	if k.Shift() {
		b.WriteString("S-")
	}
	if k.Alt() {
		b.WriteString("A-")
	}

	switch k {
	case Tab:
		b.WriteString("Tab")
	case Escape:
		b.WriteString("Esc")
	default:
		if k.Ctrl() {
			b.WriteString("C-")
		}
		b.WriteString(k.Name())
	}

	if hasMod || named {
		b.WriteRune('>')
	}
	return b.String()
}
