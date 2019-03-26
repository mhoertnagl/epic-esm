package lexer

import (
	"testing"

	"github.com/mhoertnagl/epic-esm/token"
)

func TestSingleInstruction0(t *testing.T) {
	input := "addeq r0 r1 42"

	tokens := []token.Token{
		{token.ID, "addeq"},
    {token.REG, "r0"},
    {token.REG, "r1"},
    {token.NUM, "42"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction1(t *testing.T) {
	input := "add r0 r0 42 << 1"

	tokens := []token.Token{
		{token.ID, "add"},
    {token.REG, "r0"},
    {token.REG, "r0"},
    {token.NUM, "42"},
    {token.SLL, "<<"},
    {token.NUM, "1"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction2(t *testing.T) {
	input := "add r0 r0 42 >> 1"

	tokens := []token.Token{
		{token.ID, "add"},
    {token.REG, "r0"},
    {token.REG, "r0"},
    {token.NUM, "42"},
    {token.SRL, ">>"},
    {token.NUM, "1"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction3(t *testing.T) {
	input := "add r0 r0 42 >>> 1"

	tokens := []token.Token{
		{token.ID, "add"},
    {token.REG, "r0"},
    {token.REG, "r0"},
    {token.NUM, "42"},
    {token.SRA, ">>>"},
    {token.NUM, "1"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction4(t *testing.T) {
	input := "add r0 r0 42 <<> 1"

	tokens := []token.Token{
		{token.ID, "add"},
    {token.REG, "r0"},
    {token.REG, "r0"},
    {token.NUM, "42"},
    {token.ROL, "<<>"},
    {token.NUM, "1"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction5(t *testing.T) {
	input := "add r0 r0 42 <>> 1"

	tokens := []token.Token{
		{token.ID, "add"},
    {token.REG, "r0"},
    {token.REG, "r0"},
    {token.NUM, "42"},
    {token.ROR, "<>>"},
    {token.NUM, "1"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction6(t *testing.T) {
	input := "stw r0 r1[r2]"

	tokens := []token.Token{
		{token.ID, "stw"},
    {token.REG, "r0"},
    {token.REG, "r1"},
    {token.LBRK, "["},
    {token.REG, "r2"},
    {token.RBRK, "]"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction7(t *testing.T) {
	input := "! stw r0 r1[42]"

	tokens := []token.Token{
    {token.SET, "!"},
		{token.ID, "stw"},
    {token.REG, "r0"},
    {token.REG, "r1"},
    {token.LBRK, "["},
    {token.NUM, "42"},
    {token.RBRK, "]"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleInstruction8(t *testing.T) {
	input := "brlgt @test.L0"

	tokens := []token.Token{
		{token.ID, "brlgt"},
    {token.LBL, "@test.L0"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleLineComment0(t *testing.T) {
	input := "// Comment."

	tokens := []token.Token{
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

func TestSingleLineComment1(t *testing.T) {
	input := "inv r15 // Comment."

	tokens := []token.Token{
    {token.ID, "inv"},
    {token.REG, "r15"},
		{token.EOF, string(0)},
	}
	test(t, input, tokens)
}

const msgErrUnexpType = "%d: Unexpected token type [%s]. Expecting [%s]."
const msgErrUnexpLiteral = "%d: Unexpected token literal [%s]. Expecting [%s]."

func test(t *testing.T, input string, tokens []token.Token) {
	l := NewLexer(input)
	for i, e := range tokens {
		a := l.Next()
		if a.Typ != e.Typ {
			t.Errorf(msgErrUnexpType, i, a.Typ, e.Typ)
		}
		if a.Literal != e.Literal {
			t.Errorf(msgErrUnexpLiteral, i, a.Literal, e.Literal)
		}
		t.Logf("%d: %s", i, a.Literal)
	}
}
