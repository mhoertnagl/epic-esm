package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const eof = rune(0)

type tokenType int

const (
	ERROR tokenType = iota
	EOF
	COMMAND
	REGISTER
	NUMBER
	LBRACKET
	RBRACKET
	SYMBOL
	COMMENT
)

// Token provides a set of attributes for each scanned token.
type Token struct {
	typ    tokenType
	lexeme string
	lineNo int
	chrPos int
}

func (t Token) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("[%3d:%2d] ", t.lineNo, t.chrPos))
	switch t.typ {
	case ERROR:
		buf.WriteString("ERROR   ")
	case EOF:
		buf.WriteString("EOF     ")
	case COMMAND:
		buf.WriteString("COMMAND ")
	case REGISTER:
		buf.WriteString("REGISTER")
	case NUMBER:
		buf.WriteString("NUMBER  ")
	case SYMBOL:
		buf.WriteString("SYMBOL  ")
	case COMMENT:
		buf.WriteString("COMMENT ")
	}
	buf.WriteString(fmt.Sprintf(" [%q]", t.lexeme))
	return buf.String()
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isBinDigit(c rune) bool {
	return c == '0' || c == '1'
}

func isHexDigit(c rune) bool {
	return isDigit(c) || 'a' <= c && c <= 'f'
}

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\n' //|| c == '\r' || c == '\t'
}

func isLowerAlpha(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func isUpperAlpha(c rune) bool {
	return 'A' <= c && c <= 'Z'
}

func isAlpha(c rune) bool {
	return isLowerAlpha(c) || isUpperAlpha(c) || c == '_' || c == '$'
}

func isAlphaNum(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

type Lexer struct {
	rd     *bufio.Reader
	buf    string
	tokens chan Token
	lineNo int
	chrPos int
}

func (l *Lexer) read() rune {
	r, _, err := l.rd.ReadRune()
	if err != nil {
		return eof
	}
	if r == '\n' {
		l.lineNo++
		l.chrPos = 1
	}
	l.buf = l.buf + string(r)
	return r
}

func (l *Lexer) unread() {
	l.rd.UnreadRune()
	lexemeLen := len(l.buf) - 1
	if l.buf[lexemeLen] == '\n' {
		l.lineNo--
	}
	l.buf = l.buf[:lexemeLen]
}

func (l *Lexer) peek() rune {
	r, _, err := l.rd.ReadRune()
	l.rd.UnreadRune()
	if err != nil {
		return eof
	}
	return r
}

func (l *Lexer) ignore() {
	if r, _, _ := l.rd.ReadRune(); r == '\n' {
		l.lineNo++
		l.chrPos = 1
	} else if r == ' ' {
		l.chrPos++
	}
}

func (l *Lexer) emit(typ tokenType) {
	l.tokens <- Token{
		typ:    typ,
		lexeme: l.buf,
		lineNo: l.lineNo,
		chrPos: l.chrPos,
	}
	l.drop()
}

func (l *Lexer) error(format string, a ...interface{}) {
	l.tokens <- Token{
		typ:    ERROR,
		lexeme: fmt.Sprintf(format, a...),
		lineNo: l.lineNo,
		chrPos: l.chrPos,
	}
	l.drop()
}

func (l *Lexer) drop() {
	l.chrPos += len(l.buf)
	l.buf = ""
}

type LexingPredicate func(rune) bool

func (l *Lexer) accept(p LexingPredicate, msg string) bool {
	r := l.read()
	if p(r) {
		return true
	}
	l.unread()
	l.error("Unexpected [%q]. Expecting %s.", r, msg)
	return false
}

func (l *Lexer) acceptOptional(p LexingPredicate) bool {
	if p(l.read()) {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptZeroOrMore(p LexingPredicate) bool {
	for {
		r := l.read()
		if !p(r) || r == eof {
			break
		}
	}
	l.unread()
	return true
}

func (l *Lexer) acceptOneOrMore(p LexingPredicate, msg string) {
	l.accept(p, msg)
	l.acceptZeroOrMore(p)
}

func not(p LexingPredicate) LexingPredicate {
	return func(r rune) bool {
		return !p(r)
	}
}

// or, and?

func any(v string) LexingPredicate {
	return func(r rune) bool {
		return strings.IndexRune(v, r) >= 0
	}
}

func chr(c rune) LexingPredicate {
	return func(r rune) bool {
		return c == r
	}
}

// //[^\n]*\n
func (l *Lexer) lexComment() {
	l.accept(chr('/'), "[/]")
	l.accept(chr('/'), "[/]")
	l.acceptZeroOrMore(not(chr('\n')))
	//l.emit(COMMENT)
	l.drop()
}

// [a-z]+
func (l *Lexer) lexCommand() {
	l.acceptOneOrMore(isLowerAlpha, "lower case letter")
	l.emit(COMMAND)
}

// %[0-9]+
func (l *Lexer) lexRegister() {
	l.accept(chr('%'), "[%]")
	l.acceptOneOrMore(isDigit, "a decimal digit")
	l.emit(REGISTER)
}

// (+|-)?(([0-9]+)|(0x[0-9a-f]+))
// (+|-)?(0([0-9]*)|(x[0-9a-f]+))|([1-9][0-9]*)
func (l *Lexer) lexNumber() {
	l.acceptOptional(any("+-"))
	if l.acceptOptional(chr('0')) {
		if l.acceptOptional(chr('x')) {
			l.acceptOneOrMore(isHexDigit, "at least one hexadecimal digit")
		} else {
			l.acceptZeroOrMore(isDigit)
		}
	} else {
		l.acceptOneOrMore(isDigit, "at least one decimal digit")
	}
	l.emit(NUMBER)
}

// @[a-zA-Z0-9]+
func (l *Lexer) lexSymbol() {
	l.accept(chr('@'), "[@]")
	l.acceptOneOrMore(isAlphaNum, "letter or number")
	l.emit(SYMBOL)
}

// state functions?
func (l *Lexer) lex() {
	for {
		r := l.peek()
		if r == eof {
			break
		} else if isWhitespace(r) {
			l.ignore()
		} else if r == '/' {
			l.lexComment()
		} else if isLowerAlpha(r) {
			l.lexCommand()
		} else if r == '%' {
			l.lexRegister()
		} else if isDigit(r) {
			l.lexNumber()
		} else if r == '@' {
			l.lexSymbol()
		} else {
			l.error("Unexpected [%q].", r)
			l.ignore()
		}
	}
	l.emit(EOF)
	close(l.tokens)
}

func NewLexer(rd io.Reader, size int) *Lexer {
	l := &Lexer{
		rd:     bufio.NewReaderSize(rd, size),
		tokens: make(chan Token),
		lineNo: 1,
		chrPos: 1,
	}
	go l.lex()
	return l
}
