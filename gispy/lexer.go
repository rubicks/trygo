package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pos int

type Item struct {
	Type  ItemType
	Pos   Pos
	Value string
}

type ItemType int

const (
	ItemError ItemType = iota
	ItemEOF

	ItemLeftParen
	ItemRightParen
	ItemLeftVect
	ItemRightVect

	ItemIdent
	ItemString
	ItemChar
	ItemFloat
	ItemInt
	ItemComplex

	ItemQuote
	ItemQuasiQuote
	ItemUnquote
	ItemUnquoteSplice
)

const EOF = -1

type stateFn func(*Lexer) stateFn

type Lexer struct {
	Name    string
	Input   string
	State   stateFn `json:"-"`
	Pos     Pos
	Start   Pos
	Width   Pos
	Lastpos Pos
	Items   chan Item `json:"-"`

	Parendepth int
	Vectdepth  int
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.Pos) >= len(l.Input) {
		l.Width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Width = Pos(w)
	l.Pos += l.Width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.Pos -= l.Width
}

// emit passes an Item back to the client.
func (l *Lexer) emit(t ItemType) {
	l.Items <- Item{t, l.Start, l.Input[l.Start:l.Pos]}
	l.Start = l.Pos
}

func (l *Lexer) ignore() {
	l.Start = l.Pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.Items <- Item{ItemError, l.Start, fmt.Sprintf(format, args...)}
	return nil
}

func (l *Lexer) NextItem() Item {
	item := <-l.Items
	l.Lastpos = item.Pos
	return item
}

func Lex(name, input string) *Lexer {
	l := &Lexer{
		Name:  name,
		Input: input,
		Items: make(chan Item),
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for l.State = lexWhitespace; l.State != nil; {
		l.State = l.State(l)
	}
	close(l.Items)
}

func lexLeftVect(l *Lexer) stateFn {
	l.emit(ItemLeftVect)

	return lexWhitespace
}

func lexRightVect(l *Lexer) stateFn {
	l.emit(ItemRightVect)

	return lexWhitespace
}

// lexes an open parenthesis
func lexLeftParen(l *Lexer) stateFn {
	l.emit(ItemLeftParen)

	return lexWhitespace
}

func lexWhitespace(l *Lexer) stateFn {
	for r := l.next(); isSpace(r) || r == '\n'; l.next() {
		r = l.peek()
	}
	l.backup()
	l.ignore()

	switch r := l.next(); {
	case r == EOF:
		l.emit(ItemEOF)
		return nil
	case r == '(':
		return lexLeftParen
	case r == ')':
		return lexRightParen
	case r == '[':
		return lexLeftVect
	case r == ']':
		return lexRightVect
	case r == '"':
		return lexString
	case r == '+' || r == '-' || ('0' <= r && r <= '9'):
		return lexNumber
	case r == ';':
		return lexComment
	case isAlphaNumeric(r):
		return lexIdentifier
	default:
		panic(fmt.Sprintf("don't know what to do with: %q", r))
	}
}

func lexString(l *Lexer) stateFn {
	for r := l.next(); r != '"'; r = l.next() {
		if r == '\\' {
			r = l.next()
		}
		if r == EOF {
			return l.errorf("unterminated quoted string")
		}
	}
	l.emit(ItemString)
	return lexWhitespace
}

func lexIdentifier(l *Lexer) stateFn {
	for r := l.next(); isAlphaNumeric(r); r = l.next() {
	}
	l.backup()

	l.emit(ItemIdent)
	return lexWhitespace
}

// lex a close parenthesis
func lexRightParen(l *Lexer) stateFn {
	l.emit(ItemRightParen)

	return lexWhitespace
}

// lex a comment, comment delimiter is known to be already read
func lexComment(l *Lexer) stateFn {
	i := strings.Index(l.Input[l.Pos:], "\n")
	l.Pos += Pos(i)
	l.ignore()
	return lexWhitespace
}

func lexNumber(l *Lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.Input[l.Start:l.Pos])
	}

	if l.Start+1 == l.Pos {
		return lexIdentifier
	}

	if sign := l.peek(); sign == '+' || sign == '-' {
		// Complex: 1+2i. No spaces, must end in 'i'.
		if !l.scanNumber() || l.Input[l.Pos-1] != 'i' {
			return l.errorf("bad number syntax: %q", l.Input[l.Start:l.Pos])
		}
		l.emit(ItemComplex)
	} else if strings.ContainsRune(l.Input[l.Start:l.Pos], '.') {
		l.emit(ItemFloat)
	} else {
		l.emit(ItemInt)
	}

	return lexWhitespace
}

func (l *Lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if r := l.peek(); isAlphaNumeric(r) {
		l.next()
		return false
	}
	return true
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is a valid rune for an identifier.
func isAlphaNumeric(r rune) bool {
	return r == '>' || r == '<' || r == '=' || r == '-' || r == '+' || r == '*' || r == '&' || r == '_' || r == '/' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func debug(msg string) {
	fmt.Println(msg)
}
