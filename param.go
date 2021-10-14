package termfo

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// TODO: we can actually simplify a lot of this. This uses a two-stage "lex +
// parse", but that's actually not really needed with this stack-based
// mini-language since we don't really need to know all that much about syntax,
// so we can just skip the entire lexing.
//
// Kind of obvious really but I didn't realize this until I implemented it.
// Sometimes you need to implement something at least once to fully understand
// it.
//
// Now I can't really be bothered to replace it. I didn't really *want* to write
// a terminfo implementation, it was just needed for an application. Maybe
// later.
//
// It also doesn't handle the named "dynamic" and "static" variables for now.
// None of the common terminfos seem to use it so we can probably get away with
// it. These are annoying as they persist across escape-string evaluations.
// Something of a misfeature IMO.

type item struct {
	typ itemType
	val string
	pos int
}

type lexer struct {
	input      string
	start, pos int
	items      chan item
	state      stateFn
}

func replaceParams(input string, args ...int) string {
	l := &lexer{input: input, items: make(chan item, 2), state: lexTop}
	var items []item
	for {
		select {
		case item := <-l.items:
			if item.typ == itemEOF {
				return parseParams(items, args...)
			}
			items = append(items, item)
		default:
			l.state = l.state(l)
		}
	}
}

func printParams(input string, args ...int) string {
	l := &lexer{input: input, items: make(chan item, 2), state: lexTop}
	var items []item
	for {
		select {
		case item := <-l.items:
			if item.typ == itemEOF {
				pnl := false
				b := new(strings.Builder)
				for _, c := range items {
					if c.typ == itemIf && !pnl {
						fmt.Fprint(b, "\n        ")
					}
					pnl = false
					fmt.Fprintf(b, "%s(%q)  ", c.typ, c.val)
					if c.typ == itemEndif {
						fmt.Fprint(b, "\n        ")
						pnl = true
					}
				}
				fmt.Fprint(b, "\n\n\n")
				return b.String()
			}
			items = append(items, item)
		default:
			l.state = l.state(l)
		}
	}
}

func (l *lexer) peek() byte {
	if l.pos+1 >= len(l.input) {
		return 0
	}
	return l.input[l.pos+1]
}
func (l *lexer) backup() { l.pos-- }
func (l *lexer) next() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	l.pos++
	return l.input[l.pos-1]
}
func (l *lexer) until(anyOf ...byte) {
	for {
		b := l.next()
		for _, a := range anyOf {
			if b == a {
				return
			}
		}
	}
}

func (l *lexer) emit(typ itemType) {
	l.items <- item{typ: typ, pos: l.start, val: l.input[l.start:l.pos]}
	l.start = l.pos
}

func lexTop(l *lexer) stateFn {
	switch b := l.next(); b {
	default:
		return lexTop
	case eof:
		if l.pos > l.start {
			l.emit(itemStr)
		}
		l.emit(itemEOF)
		return lexTop
	case '%':
		l.backup()
		if l.pos > l.start {
			l.emit(itemStr)
		}
		l.next()
		return lexPercent
	}
}

// This doesn't do a lot of error checking; it sort-of assumes the terminfo is
// sane.
func lexPercent(l *lexer) stateFn {
	switch b := l.next(); b {
	case eof:
		l.emit(itemEOF)
	case '%':
		l.emit(itemPercent)

	// %[[:]flags][width[.precision]][doxXs]
	case 'd', 'o', 'x', 'X', 's', 'c':
		l.emit(itemPrint)
	case ':', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		l.until('d', 'o', 'x', 'X', 's')
		l.emit(itemPrint)
	// %p[1-9]
	case 'p':
		l.next()
		l.emit(itemPush)
	// %P[a-z]
	// %P[A-Z]
	case 'P':
		n := l.next()
		if n >= 'a' && n <= 'z' {
			l.emit(itemSetDynamic)
		} else {
			l.emit(itemSetStatic)
		}
	// %g[a-z]/
	// %g[A-Z]
	// The manpage lists this with a / at the end, but that seems a typo. It's
	// not present in the actual terminfo files.
	case 'g':
		n := l.next()
		if n >= 'a' && n <= 'z' {
			l.emit(itemGetDynamic)
		} else {
			l.emit(itemGetStatic)
		}
	// %'c'
	case '\'':
		l.next()
		l.next()
		l.emit(itemCharConstant)
	// %{nn}
	case '{':
		l.until('}')
		l.emit(itemIntConstant)
	// %l   push strlen(pop)
	case 'l':
		l.emit(itemStrlen)
	// %+, %-, %*, %/, %m
	case '+':
		l.emit(itemAdd)
	case '-':
		l.emit(itemSub)
	case '*':
		l.emit(itemMult)
	case '/':
		l.emit(itemDiv)
	case 'm':
		l.emit(itemMod)
	// %&, %|, %^
	case '&':
		l.emit(itemAnd)
	case '|':
		l.emit(itemOr)
	case '^':
		l.emit(itemXor)
	// %=, %>, %<
	case '=':
		l.emit(itemEq)
	case '>':
		l.emit(itemGt)
	case '<':
		l.emit(itemLt)
	// %A, %O
	case 'A':
		l.emit(itemLogicalAnd)
	case 'O':
		l.emit(itemLogicalOr)
	// %!, %~
	case '!':
		l.emit(itemBang)
	case '~':
		l.emit(itemTilde)
	// %i
	case 'i':
		l.emit(itemIncParams)
	// %? expr %t thenpart %e elsepart %;
	case '?':
		l.emit(itemIf)
	case 't':
		l.emit(itemThen)
	case 'e':
		l.emit(itemElse)
	case ';':
		l.emit(itemEndif)
	}
	return lexTop
}

func parseParams(items []item, args ...int) string {
	params := make([]int, 9)
	for i := range args {
		params[i] = args[i]
	}

	var (
		stack []int
		doInc = 0
		b     = new(strings.Builder)
	)

	push := func(v int) { stack = append(stack, v) }
	pop := func() int {
		if len(stack) == 0 {
			// Increment first two pops even if the stack is empty; a common
			// "trick" is to use:
			//    \E[%i%d;%dR
			// to print 0 for non-ANSI terminals and 1 for ANSI.
			if doInc > 0 {
				doInc--
				return 1
			}
			return 0
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return pop
	}
	pushBool := func(v bool) {
		if v {
			stack = append(stack, 1)
		} else {
			stack = append(stack, 0)
		}
	}

	cond := false
	state := ""
	for _, item := range items {
		//fmt.Printf("%s -> %t\n", state, cond)
		if item.typ == itemEndif {
			state = ""
		}
		if state == "then" && !cond {
			if item.typ == itemElse {
				state = "else"
			}
			continue
		}
		if state == "else" && cond {
			continue
		}

		switch item.typ {
		case itemStr:
			b.WriteString(item.val)
		case itemPercent:
			b.WriteByte('%')
		case itemPush:
			a, _ := strconv.Atoi(item.val[2:])
			push(params[a-1])
		case itemPrint:
			p := pop()
			as := item.val[len(item.val)-1]
			if as == 's' {
				as = 'v'
			}
			fmt.Fprintf(b, "%"+
				strings.TrimLeft(item.val[1:len(item.val)-1], ":")+
				string(as), p)
		// from lib_tparam.c:
		// Increment the first two parameters -- if they are numbers rather than
		// strings.  As a side effect, assign into the stack; if this is
		// termcap, then the stack was populated using the termcap hack above
		// rather than via the terminfo 'p' case.
		case itemIncParams:
			doInc = 2
			params[0]++
			params[1]++
		case itemIntConstant:
			n, _ := strconv.Atoi(item.val[2 : len(item.val)-1])
			push(n)
		case itemCharConstant: // %'c'
			push(int(item.val[2]))
		case itemStrlen:
			push(len(strconv.Itoa(pop())))
		case itemAdd:
			a, b := pop(), pop()
			push(b + a)
		case itemSub:
			a, b := pop(), pop()
			push(b - a)
		case itemMult:
			a, b := pop(), pop()
			push(b * a)
		case itemDiv:
			a, b := pop(), pop()
			push(b / a)
		case itemMod:
			a, b := pop(), pop()
			push(b % a)
		case itemAnd:
			a, b := pop(), pop()
			push(b & a)
		case itemOr:
			a, b := pop(), pop()
			push(b | a)
		case itemXor:
			a, b := pop(), pop()
			push(b ^ a)
		case itemTilde:
			push(^pop())
		case itemLogicalAnd:
			a, b := pop() > 0, pop() > 0
			pushBool(b && a)
		case itemLogicalOr:
			a, b := pop() > 0, pop() > 0
			pushBool(b || a)
		case itemEq:
			a, b := pop(), pop()
			pushBool(b == a)
		case itemGt:
			a, b := pop(), pop()
			pushBool(b > a)
		case itemLt:
			a, b := pop(), pop()
			pushBool(b < a)
		case itemBang:
			pushBool(!(pop() > 0))

		case itemIf, itemEndif:
			// Handled at start.
		case itemElse:
			state = "else"
		case itemThen:
			cond = pop() > 0
			state = "then"
		}
	}
	return b.String()
}

type stateFn func(l *lexer) stateFn

func (s stateFn) String() string {
	name := runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name()
	if i := strings.LastIndexByte(name, '.'); i > -1 {
		name = name[i+1:]
	}
	if s == nil {
		name = "<nil>"
	}
	return name + "()"
}

type itemType int

const (
	itemEOF itemType = iota
	itemStr
	itemPrint
	itemPush
	itemCharConstant
	itemIntConstant
	itemPercent
	itemIncParams
	itemAdd
	itemSub
	itemMult
	itemDiv
	itemMod
	itemAnd
	itemOr
	itemXor
	itemLogicalAnd
	itemLogicalOr
	itemStrlen
	itemIf
	itemThen
	itemElse
	itemEndif
	itemEq
	itemGt
	itemLt
	itemBang
	itemTilde

	// Unhandled:
	itemGetDynamic
	itemGetStatic
	itemSetDynamic
	itemSetStatic
)

const eof = 0

func (i itemType) String() string {
	switch i {
	default:
		return "XXX"
	case itemEOF:
		return "EOF"
	case itemStr:
		return "str"
	case itemPrint:
		return "print"
	case itemPush:
		return "push"
	case itemPercent:
		return "%"
	case itemSetDynamic:
		return "setDyn"
	case itemSetStatic:
		return "setStat"
	case itemGetDynamic:
		return "getDyn"
	case itemGetStatic:
		return "getStat"
	case itemCharConstant:
		return "charConstant"
	case itemIntConstant:
		return "intConstant"
	case itemStrlen:
		return "strlen"
	case itemAdd:
		return "add"
	case itemSub:
		return "sub"
	case itemMult:
		return "mult"
	case itemDiv:
		return "div"
	case itemMod:
		return "mod"
	case itemAnd:
		return "and"
	case itemOr:
		return "or"
	case itemXor:
		return "xor"
	case itemEq:
		return "eq"
	case itemGt:
		return "gt"
	case itemLt:
		return "lt"
	case itemLogicalAnd:
		return "logicalAnd"
	case itemLogicalOr:
		return "logicalOr"
	case itemBang:
		return "bang"
	case itemTilde:
		return "tilde"
	case itemIncParams:
		return "incParams"
	case itemIf:
		return "if"
	case itemThen:
		return "then"
	case itemElse:
		return "else"
	case itemEndif:
		return "endif"
	}
}
