// Code generated by term.h.zsh; DO NOT EDIT.

package keys

import "zgo.at/termfo/caps"

// CursesVersion is the version of curses this data was generated with, as [implementation]-[version].
const CursesVersion = `ncurses-6.5.20240511`

// Keys maps caps.Cap to Key constants
var Keys = map[*caps.Cap]Key{
	caps.TableStrs[55]:  Backspace,
	caps.TableStrs[59]:  Delete,
	caps.TableStrs[61]:  Down,
	caps.TableStrs[66]:  F1,
	caps.TableStrs[67]:  F10,
	caps.TableStrs[68]:  F2,
	caps.TableStrs[69]:  F3,
	caps.TableStrs[70]:  F4,
	caps.TableStrs[71]:  F5,
	caps.TableStrs[72]:  F6,
	caps.TableStrs[73]:  F7,
	caps.TableStrs[74]:  F8,
	caps.TableStrs[75]:  F9,
	caps.TableStrs[76]:  Home,
	caps.TableStrs[77]:  Insert,
	caps.TableStrs[79]:  Left,
	caps.TableStrs[81]:  PageDown,
	caps.TableStrs[82]:  PageUp,
	caps.TableStrs[83]:  Right,
	caps.TableStrs[87]:  Up,
	caps.TableStrs[148]: BackTab,
	caps.TableStrs[164]: End,
	caps.TableStrs[165]: Enter,
	caps.TableStrs[172]: Next,
	caps.TableStrs[173]: Open,
	caps.TableStrs[174]: Options,
	caps.TableStrs[175]: Previous,
	caps.TableStrs[176]: Print,
	caps.TableStrs[177]: Redo,
	caps.TableStrs[178]: Reference,
	caps.TableStrs[179]: Refresh,
	caps.TableStrs[180]: Replace,
	caps.TableStrs[181]: Restart,
	caps.TableStrs[182]: Resume,
	caps.TableStrs[183]: Save,
	caps.TableStrs[184]: Suspend,
	caps.TableStrs[185]: Undo,
	caps.TableStrs[186]: Sbeg,
	caps.TableStrs[187]: Scancel,
	caps.TableStrs[188]: Scommand,
	caps.TableStrs[189]: Scopy,
	caps.TableStrs[190]: Screate,
	caps.TableStrs[191]: Sdc,
	caps.TableStrs[192]: Sdl,
	caps.TableStrs[193]: Select,
	caps.TableStrs[194]: Send,
	caps.TableStrs[195]: Seol,
	caps.TableStrs[196]: Sexit,
	caps.TableStrs[197]: Sfind,
	caps.TableStrs[198]: Shelp,
	caps.TableStrs[199]: Shome,
	caps.TableStrs[200]: Sic,
	caps.TableStrs[201]: Sleft,
	caps.TableStrs[202]: Smessage,
	caps.TableStrs[203]: Smove,
	caps.TableStrs[204]: Snext,
	caps.TableStrs[205]: Soptions,
	caps.TableStrs[206]: Sprevious,
	caps.TableStrs[207]: Sprint,
	caps.TableStrs[208]: Sredo,
	caps.TableStrs[209]: Sreplace,
	caps.TableStrs[210]: Sright,
	caps.TableStrs[211]: Srsume,
	caps.TableStrs[212]: Ssave,
	caps.TableStrs[213]: Ssuspend,
	caps.TableStrs[214]: Sundo,
	caps.TableStrs[216]: F11,
	caps.TableStrs[217]: F12,
	caps.TableStrs[355]: Mouse,
}

// List of all key sequences we know about. This excludes most obscure ones not
// present on modern devices.
const (
	// Special key used to signal errors.
	UnknownSequence Key = iota + (1 << 32)

	Sredo
	Snext
	Backspace
	Scommand
	Shome
	End
	Options
	Save
	Smessage
	Smove
	Refresh
	Restart
	Scancel
	Seol
	Undo
	Ssave
	Sbeg
	F1
	F2
	F3
	F4
	F5
	F6
	Ssuspend
	F7
	F8
	F9
	Shelp
	Scopy
	Left
	Sexit
	Suspend
	Sic
	Right
	Resume
	Reference
	Sundo
	Sfind
	Enter
	Select
	Redo
	Previous
	Delete
	Sright
	Down
	F10
	F11
	F12
	Send
	Up
	Print
	Next
	Home
	Sreplace
	Sdc
	Sleft
	PageDown
	Insert
	Sprint
	Sdl
	Screate
	PageUp
	BackTab
	Replace
	Soptions
	Sprevious
	Srsume
	Mouse
	Open
)

// Names of named key constants.
var keyNames = map[Key]string{
	Sredo:     `Sredo`,
	Snext:     `Snext`,
	Backspace: `Backspace`,
	Scommand:  `Scommand`,
	Shome:     `Shome`,
	End:       `End`,
	Options:   `Options`,
	Save:      `Save`,
	Smessage:  `Smessage`,
	Smove:     `Smove`,
	Refresh:   `Refresh`,
	Restart:   `Restart`,
	Scancel:   `Scancel`,
	Seol:      `Seol`,
	Undo:      `Undo`,
	Ssave:     `Ssave`,
	Sbeg:      `Sbeg`,
	F1:        `F1`,
	F2:        `F2`,
	F3:        `F3`,
	F4:        `F4`,
	F5:        `F5`,
	F6:        `F6`,
	Ssuspend:  `Ssuspend`,
	F7:        `F7`,
	F8:        `F8`,
	F9:        `F9`,
	Shelp:     `Shelp`,
	Scopy:     `Scopy`,
	Left:      `Left`,
	Sexit:     `Sexit`,
	Suspend:   `Suspend`,
	Sic:       `Sic`,
	Right:     `Right`,
	Resume:    `Resume`,
	Reference: `Reference`,
	Sundo:     `Sundo`,
	Sfind:     `Sfind`,
	Enter:     `Enter`,
	Select:    `Select`,
	Redo:      `Redo`,
	Previous:  `Previous`,
	Delete:    `Delete`,
	Sright:    `Sright`,
	Down:      `Down`,
	F10:       `F10`,
	F11:       `F11`,
	F12:       `F12`,
	Send:      `Send`,
	Up:        `Up`,
	Print:     `Print`,
	Next:      `Next`,
	Home:      `Home`,
	Sreplace:  `Sreplace`,
	Sdc:       `Sdc`,
	Sleft:     `Sleft`,
	PageDown:  `PageDown`,
	Insert:    `Insert`,
	Sprint:    `Sprint`,
	Sdl:       `Sdl`,
	Screate:   `Screate`,
	PageUp:    `PageUp`,
	BackTab:   `BackTab`,
	Replace:   `Replace`,
	Soptions:  `Soptions`,
	Sprevious: `Sprevious`,
	Srsume:    `Srsume`,
	Mouse:     `Mouse`,
	Open:      `Open`,
}
