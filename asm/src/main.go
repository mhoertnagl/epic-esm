package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
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
	number int32
	lineNo int
	chrPos int
}

func (t *Token) String() string {
	var buf bytes.Buffer
	switch t.typ {
	case ERROR:
		buf.WriteString("ERROR")
	case EOF:
		buf.WriteString("EOF")
	case COMMAND:
		buf.WriteString("COMMAND")
	case REGISTER:
		buf.WriteString("REGISTER")
	case NUMBER:
		buf.WriteString("NUMBER")
	case SYMBOL:
		buf.WriteString("SYMBOL")
	case COMMENT:
		buf.WriteString("COMMENT")
	}
	buf.WriteString(" lexeme=[")
	buf.WriteString(strings.Replace(t.lexeme, "\n", "\\n", -1))
	buf.WriteString(fmt.Sprintf("] num=[%d] line=[%d] pos=[%d]", t.number, t.lineNo, t.chrPos))
	return buf.String()
}

// TODO: Channel für die Ausgabe von Tokens verwenden?

type Lexer struct {
	rd *bufio.Reader
	//buf    *bytes.Buffer
	buf    string
	bufLen int
	lineNo int
	chrPos int
}

func NewLexer(rd io.Reader, size int) *Lexer {
	return &Lexer{
		rd:     bufio.NewReaderSize(rd, size),
		buf:    "", //new(bytes.Buffer),
		lineNo: 1,
		chrPos: 0,
	}
}

func (l *Lexer) emit(typ tokenType) *Token {
	//return &Token{typ, l.buf.String(), 0, l.lineNo, l.chrPos}
	return &Token{typ, l.buf, 0, l.lineNo, l.chrPos}
}

// func (l *Lexer) error(msg string) *Token {
// 	return &Token{ERROR, msg, 0, l.lineNo, l.chrPos}
// }

func (l *Lexer) read() rune {
	r, _, err := l.rd.ReadRune()
	if err != nil {
		return eof
	}
	if r == '\n' {
		l.lineNo++
		l.chrPos = 0
	} else {
		l.chrPos++
	}
	//l.buf.WriteRune(r)
	//fmt.Printf("Buffer Write [%q] + [%q]\n", l.buf, r)
	l.buf = l.buf + string(r)
	l.bufLen++
	//fmt.Printf("Buffer Write [%q](%d) pos=[%d]\n", l.buf, l.bufLen, l.chrPos)
	return r
}

func (l *Lexer) unread() {
	l.rd.UnreadRune()
	if l.buf[len(l.buf)-1] == '\n' {
		l.lineNo--
	}
	l.chrPos--
	//l.buf.ReadRune()
	//fmt.Printf("Buffer Unwrite [%q] last is [%q]\n", l.buf, l.buf[len(l.buf)-1])
	l.buf = l.buf[:len(l.buf)-1]
	l.bufLen--
	//fmt.Printf("Buffer Unwrite [%q]\n", l.buf)
}

func (l *Lexer) peek() rune {
	r := l.read()
	l.unread()
	return r
}

func (l *Lexer) backup(n int) {
	for i := 0; i < n; i++ {
		l.unread()
	}
}

func (l *Lexer) oneOf(v string) bool {
	r := l.read()
	//fmt.Printf("[%q] oneOf [%q]\n", r, v)
	if strings.IndexRune(v, r) >= 0 {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) manyOf(v string) bool {
	for l.oneOf(v) {
	}
	return true
}

func (l *Lexer) atLeastOneOf(v string) bool {
	if l.oneOf(v) {
		return l.manyOf(v)
	}
	return false
}

func (l *Lexer) acceptSeq(s string) bool {
	for i, r := range s {
		//fmt.Printf("[%q] in seq [%s][%d]?\n", r, s, i)
		if !l.acceptRune(r) {
			l.backup(i)
			//fmt.Printf(" No\n")
			return false
		}
		//fmt.Printf(" Yes\n")
	}
	return true
}

func (l *Lexer) acceptRune(r rune) bool {
	c := l.read()
	//fmt.Printf("acceptRune c=[%q] ? r=[%q]\n", c, r)
	if c == r {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptFunc(p func(rune) bool) bool {
	if p(l.read()) {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptFuncSeq(p func(rune) bool) bool {
	for l.acceptFunc(p) {
	}
	return false
}

func (l *Lexer) acceptUntilRune(r rune) bool {
	for {
		c := l.read()
		if c == r || c == eof {
			break
		}
	}
	l.unread()
	return false
}

func (l *Lexer) acceptUntilAny(v string) bool {
	for {
		c := l.read()
		if strings.IndexRune(v, c) >= 0 || c == eof {
			break
		}
	}
	l.unread()
	return false
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isBinDigit(c rune) bool {
	return c == '0' || c == '1'
}

func isHexDigit(c rune) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f'
}

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t'
}

func isLowerAlpha(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func isAlpha(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c == '$'
}

func isAlphaNum(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

func (l *Lexer) Reset() {
	// for i := 0; i < l.bufLen; i++ {
	// 	l.rd.UnreadRune()
	// }
	//l.buf.Reset()
	l.buf = ""
	l.bufLen = 0
}

// Next returns the next token in the stream.
func (l *Lexer) Next() *Token {
	l.Reset()
	// c := l.read()
	// switch c {
	// case eof:
	// 	return l.emit(EOF)
	// case '[':
	// 	return l.emit(LBRACKET)
	// case ']':
	// 	return l.emit(RBRACKET)
	// case '/':
	// 	return l.scanSingleLineComment()
	// case '%':
	// 	return l.scanRegister()
	// case '@':
	// 	return l.scanSymbol()
	// }
	//
	// if isWhitespace(c) {
	// 	return l.Next()
	// }
	//
	// if isDigit(c) {
	// 	d := l.read()
	// 	if d == 'x' {
	// 		return l.scanHexNumber()
	// 	} else if isDigit(d) {
	// 		return l.scanDecNumber()
	// 	}
	// 	return l.emit(ERROR)
	// }
	//
	// if isLowerAlpha(c) {
	// 	return l.scanCommand()
	// }
	if l.acceptRune(eof) {
		return l.emit(EOF)
	}
	if l.atLeastOneOf(" \t\r\n") {
		//l.scanWhitespace()
		return l.Next()
	}
	if l.acceptRune('[') {
		return l.emit(LBRACKET)
	}
	if l.acceptRune(']') {
		return l.emit(RBRACKET)
	}
	if l.acceptSeq("//") {
		return l.scanSingleLineComment()
	}
	if l.acceptSeq("0x") {
		return l.scanHexNumber()
	}
	if l.acceptFunc(isDigit) {
		return l.scanDecNumber()
	}
	if l.acceptRune('@') {
		return l.scanSymbol()
	}
	if l.acceptRune('%') {
		return l.scanRegister()
	}
	if l.acceptFunc(isLowerAlpha) {
		return l.scanCommand()
	}
	return l.emit(ERROR)
}

// func (l *Lexer) scanWhitespace() *Token {
// 	l.atLeastOneOf(" \t\r\n")
// 	return l.Next()
// }

func (l *Lexer) scanSingleLineComment() *Token {
	l.acceptUntilAny("\n\r")
	return l.emit(COMMENT)
}

func (l *Lexer) scanHexNumber() *Token {
	l.acceptFuncSeq(isHexDigit)
	t := l.emit(NUMBER)
	n, _ := strconv.ParseInt(t.lexeme[2:], 16, 32)
	t.number = int32(n)
	return t
}

func (l *Lexer) scanDecNumber() *Token {
	l.acceptFuncSeq(isDigit)
	t := l.emit(NUMBER)
	n, _ := strconv.ParseInt(t.lexeme, 10, 32)
	t.number = int32(n)
	return t
}

func (l *Lexer) scanRegister() *Token {
	l.acceptFuncSeq(isAlphaNum)
	return l.emit(REGISTER)
}

func (l *Lexer) scanSymbol() *Token {
	l.acceptFuncSeq(isAlphaNum)
	return l.emit(SYMBOL)
}

func (l *Lexer) scanCommand() *Token {
	l.acceptFuncSeq(isLowerAlpha)
	return l.emit(COMMAND)
}

var text = `
// Toller Test
// ------------
@L0
  add   %0 %1 %2    // Test
  sll   %0 %0 2
  // Noch ein Kommentar
  tst   %0 %0 0xff
  brlgt @L0
`

func main() {
	rd := strings.NewReader(text)
	l := NewLexer(rd, 2048)
	for {
		t := l.Next()
		fmt.Println(t)
		if t.typ == EOF {
			fmt.Println("Done.")
			break
		}
	}
}
