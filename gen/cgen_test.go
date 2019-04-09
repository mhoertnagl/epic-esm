package gen

import (
  "fmt"
	"testing"
  
  "github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/ast"
)

// 0000 0
// 0001 1
// 0010 2
// 0011 3
// 0100 4
// 0101 5
// 0110 6
// 0111 7
// 1000 8
// 1001 9
// 1010 A
// 1011 B
// 1100 C
// 1101 D
// 1110 E
// 1111 F

func TestI16_0(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "nv",
    Args: []token.Token{
      {Typ: token.REG, Literal: "rp"},
      {Typ:token.NUM, Literal: fmt.Sprintf("%d", 0xFFFF)},
  	},
  }
  // 0010 0000 1110 1111  1111 1111 1111 0000
	test(t, ins, 0x20EFFFF0)
}

func TestI16_1(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "nv",
    Args: []token.Token{
      {Typ: token.REG, Literal: "rp"},
      {Typ:token.NUM, Literal: fmt.Sprintf("%d", 0xFFFF)},
      {Typ: token.SOP, Literal: "<<"},
      {Typ: token.NUM, Literal: "16"},
  	},
  }
  // 0010 0001 1110 1111  1111 1111 1111 0000
	test(t, ins, 0x21EFFFF0)
}

func TestI16_2(t *testing.T) {
  ins := &ast.Instr{
    Set: true,
    Cmd: "add",
    Cond: "nv",
    Args: []token.Token{
      {Typ: token.REG, Literal: "rp"},
      {Typ:token.NUM, Literal: fmt.Sprintf("%d", 0xFFFF)},
      {Typ: token.SOP, Literal: "<<"},
      {Typ: token.NUM, Literal: "16"},
  	},
  }
  // 0010 0011 1110 1111  1111 1111 1111 0000
	test(t, ins, 0x23EFFFF0)
}

func TestInstruction0(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ:token.REG, Literal: "r1"},
      {Typ:token.NUM, Literal: "42"},
  	},
  }
  // 0001 1101 0000 0001  0000 0010 1010 0000
	test(t, ins, 0x1D0102A0)
}

func TestInstruction00(t *testing.T) {
  ins := &ast.Instr{
    Set: true,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ:token.REG, Literal: "r1"},
      {Typ:token.NUM, Literal: "42"},
  	},
  }
  // 0001 1111 0000 0001  0000 0010 1010 0000
	test(t, ins, 0x1F0102A0)
}

func TestInstruction1(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r2"},
      {Typ:token.REG, Literal: "r3"},
      {Typ:token.REG, Literal: "r4"},
      {Typ:token.SOP, Literal: ">>>"},
      {Typ:token.NUM, Literal: "3"},
  	},
  }
  // 0001 1100 0010 0011  0100 0100 0011 0000
	test(t, ins, 0x1C234430)
}

const msgErr = "Unexpected code [%08X]. Expecting code [%08X]."

func test(t *testing.T, ins *ast.Instr, exp uint32) {
  st := NewSymbolTable()
  ctx := NewAsmContext("test.esm", st)
  g := NewCodeGen(ctx)
  code := g.Generate(ins)
  if code != exp {
    t.Errorf(msgErr, code, exp)
  }
}
