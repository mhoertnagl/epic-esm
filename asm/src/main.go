package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

var text = `
// Toller Test
// ------------
@L0
  add   r0 r1 r2    // Test
  sll   r0 r0 2
  // Noch ein Kommentar
  tst   r0 r0 0
  brlgt @L0
`

const eof = rune(0)

type tokenType int

const (
	ERROR tokenType = iota
	NOOP
	EOF
	WS
	COMMAND
	COMMAND_CONDITION
	REGISTER
	NUMBER
	SYMBOL
	COMMENT
)

//
// type item struct {
// 	typ itemType // Type, such as itemNumber.
// 	val string   // Value, such as "23.2".
// }

// Token provides a set of attributes for each scanned token.
// And i dont know what else to write.
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
	case NOOP:
		buf.WriteString("NOOP")
	case EOF:
		buf.WriteString("EOF")
	case WS:
		buf.WriteString("WS")
	case COMMAND:
		buf.WriteString("COMMAND")
	case COMMAND_CONDITION:
		buf.WriteString("COMMAND_CONDITION")
	case REGISTER:
		buf.WriteString("REGISTER")
	case NUMBER:
		buf.WriteString("NUMBER")
	case SYMBOL:
		buf.WriteString("SYMBOL")
	case COMMENT:
		buf.WriteString("SINGLE_LINE_COMMENT")
	}
	buf.WriteString(" lexeme=[")
	buf.WriteString(strings.Replace(t.lexeme, "\n", "\\n", -1))
	buf.WriteString(fmt.Sprintf("] num=[%d] line=[] pos=[]", t.number))
	return buf.String()
}

// TODO: Peek statt read.
// TODO: Liest den nächsten Character. Fügt ihn dem Buffer hinzu wenn p die
//       rune akzeptiert. wenn nicht sollte ein error generiert werden.
//       Methode (l *Lexer) Accept(p func(rune)bool) rune
// TODO: Buffer in Scanner integrieren. mit read char hinzufügen. Mit unread
//       zeichen wieder aus bufer entfernen.

type Lexer struct {
	reader *bufio.Reader
	lineNo int
	chrPos int
}

func NewLexer(reader *bufio.Reader) *Lexer {
	return &Lexer{reader, 0, 0}
}

func (l *Lexer) emit(typ tokenType) *Token {
	return &Token{typ, "", 0, l.lineNo, l.chrPos}
}

func (l *Lexer) error(msg string) *Token {
	return &Token{ERROR, msg, 0, l.lineNo, l.chrPos}
}

func (l *Lexer) noop() *Token {
	return &Token{NOOP, "", 0, 0, 0}
}

func (l *Lexer) read() rune {
	c, _, err := l.reader.ReadRune()
	if err != nil {
		return eof
	}
	return c
}

func (l *Lexer) unread() {
	l.reader.UnreadRune()
}

func (l *Lexer) acceptAny(v string) bool {
	if strings.IndexRune(v, l.read()) >= 0 {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptAnySeq(v string) bool {
	for l.acceptAny(v) {
	}
	l.unread()
	return false
}

func (l *Lexer) acceptSeq(s string) bool {
	for _, r := range s {
		if !l.acceptRune(r) {
			break
		}
	}
	l.unread()
	return false
}

func (l *Lexer) acceptRune(r rune) bool {
	if l.read() == r {
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
	l.unread()
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

func (l *Lexer) acceptOneOf(a []string) bool {
	if strings.IndexRune(v, l.read()) >= 0 {
		return true
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

// Next returns the next token in the stream.
func (l *Lexer) Next() *Token {
	if l.acceptRune(eof) {
		return l.emit(EOF)
	}
	if l.acceptRune('\n') {
		l.lineNo++
		return l.Next()
	}
	if l.acceptRune('\r') {
		return l.Next()
	}
	if l.acceptAny(" \t") {
		l.scanWhitespace()
	}
	if l.acceptSeq("//") {
		l.scanSingleLineComment()
	}
	if l.acceptSeq("0x") {
		l.scanHexNumber()
	}
	if l.acceptFunc(isDigit) {
		l.scanDecNumber()
	}
	if l.acceptRune('@') {
		l.scanSymbol()
	}
	if l.acceptFunc(isLowerAlpha) {
		l.scanComand()
	}
	return nil
}

func (l *Lexer) scanWhitespace() *Token {
	l.acceptAnySeq(" \t")
	return l.Next()
}

func (l *Lexer) scanSingleLineComment() *Token {
	l.acceptUntilAny("\n\r")
	return l.emit(COMMENT)
}

func (l *Lexer) scanHexNumber() *Token {
	l.acceptFuncSeq(isHexDigit)
	return l.emit(NUMBER)
}

func (l *Lexer) scanDecNumber() *Token {
	l.acceptFuncSeq(isDigit)
	return l.emit(NUMBER)
}

func (l *Lexer) scanSymbol() *Token {
	l.acceptFuncSeq(isAlphaNum)
	return l.emit(SYMBOL)
}

func main() {
	reader := bufio.NewReader(strings.NewReader(text))
	l := NewLexer(reader)
	fmt.Println(l.Next())
	fmt.Println(l.Next())
	fmt.Println(l.Next())
	fmt.Println(l.Next())
	fmt.Println(l.Next())
	fmt.Println(l.Next())
	fmt.Println(l.Next())
}
