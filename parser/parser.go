package parser

import (
	"fmt"

	"github.com/mhoertnagl/epic-esm/lexer"
	"github.com/mhoertnagl/epic-esm/token"
)

type Parser struct {
	lexer          *lexer.Lexer
	curToken       token.Token
	nxtToken       token.Token
	errors         []string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: []string{},
	}

	// Sets the parsers current and next tokens.
	p.next()
	p.next()
	return p
}

func (p *Parser) debug(prefix string) {
  fmt.Printf("%s: Current: %s | Next: %s\n", prefix, p.curToken, p.nxtToken)
}

func (p *Parser) next() {
	p.curToken = p.nxtToken
	p.nxtToken = p.lexer.Next()
}

func (p *Parser) curTokenIs(exp token.TokenType) bool {
	return p.curToken.Typ == exp
}

func (p *Parser) nxtTokenIs(exp token.TokenType) bool {
	return p.nxtToken.Typ == exp
}

func (p *Parser) expectNext(exp token.TokenType) bool {
	if p.nxtTokenIs(exp) {
		p.next()
		return true
	}
	p.errorNext(exp)
	return false
}

func (p *Parser) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) error(format string, a ...interface{}) {
	p.errors = append(p.errors, fmt.Sprintf(format, a...))
}

func (p *Parser) errorNext(exp token.TokenType) {
	p.error("Expected token [%s] but got [%s].", exp, p.nxtToken.Typ)
}

func (p *Parser) Parse() Node {
  switch p.curToken.Typ {
  case token.SET:
    return p.parseInstruction()
  case token.ID:
    return p.parseInstruction()
  }
	return &Err{}
}

func (p *Parser) parseInstruction() Node {
  ins := &Instr{}
  if p.curTokenIs(token.SET) {
    ins.Set = true
    p.next()
  }
  if p.curTokenIs(token.ID) {
    c := p.curToken.Literal
    if len(c) == 5 {
      ins.Cmd = c[:3]
      ins.Cond = c[3:]
    } else {
      ins.Cmd = c
      ins.Cond = "al"      
    }
    p.next()
  } else {
    // Error: expecting assembler command.
    return nil
  }
  for !p.curTokenIs(token.EOF) {
    ins.Args = append(ins.Args, p.curToken)
    p.next()
  }
  return ins
}
