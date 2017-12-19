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

type tokentype int

const (
	ERROR tokentype = iota
	NOOP
	EOF
	WS
	COMMAND
	COMMAND_CONDITION
	REGISTER
	NUMBER
	SYMBOL
	SINGLE_LINE_COMMENT
	MULTI_LINE_COMMENT
)

type Token struct {
	typ    tokentype
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
	case SINGLE_LINE_COMMENT:
		buf.WriteString("SINGLE_LINE_COMMENT")
	case MULTI_LINE_COMMENT:
		buf.WriteString("MULTI_LINE_COMMENT")
	}
	buf.WriteString(" lexeme=[")
	buf.WriteString(strings.Replace(t.lexeme, "\n", "\\n", -1))
	buf.WriteString(fmt.Sprintf("] num=[%d] line=[] pos=[]", t.number))
	return buf.String()
}

// TODO: Peek statt read.
// TODO: Liest den nächsten Character. Fügt ihn dem Buffer hinzu wenn p die
//       rune akzeptiert. wenn nicht sollte ein error generiert werden.
//       Methode (scanner *EsmScanner) Accept(p func(rune)bool) rune
// TODO: Buffer in Scanner integrieren. mit read char hinzufügen. Mit unread
//       zeichen wieder aus bufer entfernen.

type EsmScanner struct {
	commands map[string]bool
	reader   *bufio.Reader
	lineNo   int
	chrPos   int
}

func NewEsmScanner(reader *bufio.Reader) *EsmScanner {
	s := &EsmScanner{make(map[string]bool), reader, 0, 0}
	s.commands["add"] = true
	return s
}

func (scanner *EsmScanner) newToken(typ tokentype) *Token {
	return &Token{typ, "", 0, scanner.lineNo, scanner.chrPos}
}

func (scanner *EsmScanner) error(msg string) *Token {
	return &Token{ERROR, msg, 0, scanner.lineNo, scanner.chrPos}
}

func (scanner *EsmScanner) noop() *Token {
	return &Token{NOOP, "", 0, 0, 0}
}

func (scanner *EsmScanner) read() rune {
	c, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return c
}

func (scanner *EsmScanner) unread() {
	scanner.reader.UnreadRune()
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isBinDigit(c rune) bool {
	return c == '0' || c == '9'
}

func isHexDigit(c rune) bool {
	return 'a' <= c && c <= 'f'
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

func (scanner *EsmScanner) Next() *Token {
	c := scanner.read()
	//fmt.Printf("%q\n", c)
	if c == eof {
		return scanner.newToken(EOF)
	}
	if isWhitespace(c) {
		scanner.unread()
		return scanner.scanWhitespace()
	}
	if c == '/' {
		return scanner.scanComment()
	}
	if isDigit(c) {
		return scanner.scanNumber(c)
	}
	if c == '@' {
		return scanner.scanSymbol()
	}
	if isLowerAlpha(c) {
		return scanner.scanCommand(c)
	}
	return nil
}

func (scanner *EsmScanner) scanWhitespace() *Token {
	var buf bytes.Buffer
	for {
		c := scanner.read()
		if isWhitespace(c) {
			buf.WriteString(fmt.Sprintf("%q", c))
		} else {
			scanner.unread()
			break
		}
	}
	return &Token{WS, buf.String(), 0, 0, 0}
}

func (scanner *EsmScanner) scanBinNumber() int32 {
	var n int32
	for {
		c := scanner.read()
		if isBinDigit(c) {
			n = 2*n + (int32)(c-rune('0'))
		} else {
			scanner.unread()
			break
		}
	}
	return n
}

func (scanner *EsmScanner) scanHexNumber() int32 {
	var n int32
	for {
		c := scanner.read()
		if isDigit(c) {
			n = 16*n + (int32)(c-rune('0'))
		} else if isHexDigit(c) {
			n = 16*n + (int32)(c-rune('a')+10)
		} else {
			scanner.unread()
			break
		}
	}
	return n
}

func (scanner *EsmScanner) scanDecNumber() int32 {
	var n int32
	for {
		c := scanner.read()
		if isDigit(c) {
			n = 10*n + (int32)(c-rune('0'))
		} else {
			scanner.unread()
			break
		}
	}
	return n
}

func (scanner *EsmScanner) scanNumber(c rune) *Token {
	lineNo := scanner.lineNo
	chrPos := scanner.chrPos
	t := &Token{NUMBER, "", c, lineNo, chrPos}
	//c := scanner.read()
	if c == '0' {
		l := scanner.read()
		if l == 'b' {
			t.number = scanner.scanBinNumber()
			return t
		} else if l == 'x' {
			t.number = scanner.scanHexNumber()
			return t
		} else if isDigit(l) {
			t.number = scanner.scanDecNumber()
			return t
		}
	} else if isDigit(c) {
		t.number = scanner.scanDecNumber()
		return t
	} else {
		scanner.unread()
	}
	return scanner.noop()
}

func (scanner *EsmScanner) scanSymbol() *Token {
	var buf bytes.Buffer
	buf.WriteRune('@')
	c := scanner.read()
	if isAlpha(c) {
		buf.WriteRune(c)
	} else {
		return scanner.error("Invalid symbol character. Expecting letter.")
	}
	for {
		c := scanner.read()
		if isAlphaNum(c) {
			buf.WriteRune(c)
		} else {
			scanner.unread()
			return &Token{SYMBOL, buf.String(), 0, 0, 0}
		}
	}
}

func (scanner *EsmScanner) scanCommand(c rune) *Token {
	var buf bytes.Buffer

	//c := scanner.read()
	if isLowerAlpha(c) {
		buf.WriteRune(c)
	} else {
		return scanner.error(fmt.Sprintf("Invalid character [%q]. Expecting letter.", c))
	}

	c = scanner.read()
	if isLowerAlpha(c) {
		buf.WriteRune(c)
	} else {
		return scanner.error(fmt.Sprintf("Invalid character [%q]. Expecting letter.", c))
	}

	c = scanner.read()
	if isLowerAlpha(c) {
		buf.WriteRune(c)
	} else {
		return scanner.error(fmt.Sprintf("Invalid character [%q]. Expecting letter.", c))
	}

	cmd := buf.String()
	if scanner.commands[cmd] {
		return &Token{COMMAND, cmd, 0, 0, 0}
	}
	return scanner.error(fmt.Sprintf("Unrecognized command [%s].", cmd))
}

func (scanner *EsmScanner) scanComment() *Token {
	lineNo := scanner.lineNo
	chrPos := scanner.chrPos
	var buf bytes.Buffer
	//c := scanner.read()
	//if c == '/' {
	buf.WriteRune('/')
	l := scanner.read()
	if l == '/' {
		buf.WriteRune(l)
		for {
			x := scanner.read()
			if x == '\n' || x == eof {
				buf.WriteRune('\n')
				break
			} else {
				buf.WriteRune(x)
			}
		}
		return &Token{SINGLE_LINE_COMMENT, buf.String(), 0, lineNo, chrPos}
	} else if l == '*' {
		buf.WriteRune(l)
		for {
			x := scanner.read()
			buf.WriteRune(x)
			if x == '*' {
				y := scanner.read()
				buf.WriteRune(y)
				if y == '/' {
					return &Token{MULTI_LINE_COMMENT, buf.String(), 0, lineNo, chrPos}
				} else if x == eof {
					return scanner.error("Unexpected end of file. Expecting end of multi line comment '/'.")
				}
			} else if x == eof {
				return scanner.error("Unexpected end of file. Expecting end of multi line comment '*/'.")
			}
		}
	} else {
		return scanner.error("Invalid input '?'. Expecting '/' or '*'")
	}
	//}
	//return scanner.error("Expecting comment.")
}

/**
 * Test
 */
func main() {
	reader := bufio.NewReader(strings.NewReader(text))
	scanner := NewEsmScanner(reader)
	fmt.Println(scanner.Next())
	fmt.Println(scanner.Next())
	fmt.Println(scanner.Next())
	fmt.Println(scanner.Next())
	fmt.Println(scanner.Next())
	fmt.Println(scanner.Next())
	fmt.Println(scanner.Next())
}
