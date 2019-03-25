package parser

import (
	"testing"

	"github.com/mhoertnagl/epic-esm/lexer"
)

func TestStatements(t *testing.T) {
	test(t, "add r0 r1 r2", "add r0 r1 r2")
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
