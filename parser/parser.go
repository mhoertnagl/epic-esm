package parser

import (
	"fmt"

	"github.com/mhoertnagl/epic-esm/lexer"
	"github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/gen"
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

func (p *Parser) expectNextLiteral(exp string) bool {
	if p.nxtToken.Literal == exp {
		p.next()
		return true
	}
	p.errorNextLiteral(exp)
	return false
}

func (p *Parser) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) errorNode() Node {
	return &Err{}
}

func (p *Parser) error(format string, a ...interface{}) {
	p.errors = append(p.errors, fmt.Sprintf(format, a...))
}

func (p *Parser) errorNext(exp token.TokenType) {
	p.error("Expected token [%s] but got [%s].", exp, p.nxtToken.Typ)
}

func (p *Parser) errorNextLiteral(exp string) {
	p.error("Expected literal [%s] but got [%s].", exp, p.nxtToken.Literal)
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

// func (p *Parser) parseInstruction() Node {
//   ins := &Instr{}
//   if p.curTokenIs(token.SET) {
//     ins.Set = true
//     p.next()
//   }
//   if p.curTokenIs(token.ID) {
//     c := p.curToken.Literal
//     if len(c) == 5 {
//       ins.Cmd = c[:3]
//       ins.Cond = c[3:]
//     } else {
//       ins.Cmd = c
//       ins.Cond = "al"      
//     }
//     p.next()
//   } else {
//     // Error: expecting assembler command.
//     return &Err{}
//   }
//   for !p.curTokenIs(token.EOF) {
//     ins.Args = append(ins.Args, p.curToken)
//     p.next()
//   }
//   return ins
// }

func (p *Parser) parseInstruction() Node {
  set := false
  cmd := ""
  cond := "al"

  if p.curTokenIs(token.SET) {
    set = true
    p.next()
  }
  
  if p.curTokenIs(token.ID) {
    c := p.curToken.Literal
    if len(c) == 5 {
      cmd = c[:3]
      cond = c[3:]
    } else {
      cmd = c
      cond = "al"
    }
  } else {
    // Error: expecting assembler command.
    return p.errorNode()
  }
  
  if gen.IsDataInstruction(cmd) {
    return p.parseDataInstr(set, cmd, cond)
  }
  
  p.error("Unexpected command [%s]", cmd)
  return p.errorNode()
}

func (p *Parser) parseDataInstr(set bool, cmd string, cond string) Node {
  if !p.expectNext(token.REG) {
    return p.errorNode()
  }
  rd := p.curToken.Literal
  if p.nxtTokenIs(token.NUM) {
    return p.parseI16Instr(set, cmd, cond, rd)
  }
  if !p.expectNext(token.REG) {
    return p.errorNode()
  }
  ra := p.curToken.Literal
  if p.nxtTokenIs(token.NUM) {
    return p.parseI12Instr(set, cmd, cond, rd, ra)
  }
  if !p.expectNext(token.REG) {
    return p.errorNode()
  }
  rb := p.curToken.Literal
  //if !p.expectNext(token.EOF) {
  sh := p.parseShift()
  //}  
  return &RegInstruction{
    Set: set,
    Cmd: cmd,
    Cnd: cond,
    Rd: rd,
    Ra: ra,
    Rb: rb,
    Sh: sh,
  }
}

func (p *Parser) parseI16Instr(set bool, cmd string, cond string, rd string) Node {
  p.next()
  num := p.curToken
  up := false
  if p.nxtTokenIs(token.SLL) {
    p.next()
    if !p.expectNextLiteral("16") {
      return p.errorNode()
    }
    up = true
  }
  return &I16Instruction{
    Set: set,
    Cmd: cmd,
    Cnd: cond,
    Rd: rd,
    Num: num.Literal,
    Up: up,
  }
}

func (p *Parser) parseI12Instr(set bool, cmd string, cond string, rd string, ra string) Node {
  p.next()
  num := p.curToken
  return &I12Instruction{
    Set: set,
    Cmd: cmd,
    Cnd: cond,
    Rd: rd,
    Ra: ra,
    Num: num.Literal,
  }
}

func (p *Parser) parseShift() *NumShift {
  sh := &NumShift{}
  if !p.nxtTokenIs(token.EOF) {
    p.next()
    sh.Cmd = p.curToken.Literal
    if !gen.IsShiftOp(sh.Cmd) {
      p.error("Expecting shift operator.")
      return p.nullShift()
    }
    if !p.expectNext(token.NUM) {
      return p.nullShift()
    }
    sh.Num = p.curToken.Literal
  }
  return p.nullShift()
}

func (p *Parser) nullShift() *NumShift {
  return &NumShift{ Cmd: "sll", Num: "0" }
}
