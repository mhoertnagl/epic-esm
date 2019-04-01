package parser

import (
	"testing"

	"github.com/mhoertnagl/epic-esm/lexer"
)

func TestDataErrors(t *testing.T) {
  test(t, "add 42", "ERROR")
  test(t, "add sub", "ERROR")
  test(t, "add r0 sub", "ERROR")
  test(t, "add r0 r4 r4 r7", "ERROR")
  test(t, "add r0 r4 r4 >>>", "ERROR")
  test(t, "add r0 r4 r4 >>> r8", "ERROR")
}

func TestData3Reg(t *testing.T) {
	test(t, "add r0 r1 r2", "add al r0 r1 r2 << 0")
  test(t, "! subeq r0 r1 r2", "! sub eq r0 r1 r2 << 0")
  
  test(t, "mullt r0 r2", "mul lt r0 r0 r2 << 0")
  
  test(t, "add r0 r2 r2 >> 2", "add al r0 r2 r2 >> 2")
}

func TestDataI12(t *testing.T) {
  test(t, "add r0 42", "add al r0 42")
}

func test(t *testing.T, input string, expected string) {
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)
	root := parser.Parse()
	actual := root.String()
	if actual != expected {
		t.Errorf("Expected [%s] but got [%s].", expected, actual)
	}
}
