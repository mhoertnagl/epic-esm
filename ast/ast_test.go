package ast

import (
	"testing"
  
  "github.com/mhoertnagl/epic-esm/token"
)

func TestInstruction0(t *testing.T) {
  ins := &Instr{
    Set: false,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r1"},
      {Typ: token.NUM, Literal: "42"},
  	},
  }
	test(t, ins, "_ add c r r n")
}

func TestInstruction1(t *testing.T) {
  ins := &Instr{
    Set: true,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r1"},
      {Typ: token.NUM, Literal: "42"},
  	},
  }
	test(t, ins, "! add c r r n")
}

func TestInstruction2(t *testing.T) {
  ins := &Instr{
    Set: true,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r1"},
      {Typ: token.REG, Literal: "r2"},
  	},
  }
	test(t, ins, "! add c r r r")
}

func TestInstruction3(t *testing.T) {
  ins := &Instr{
    Set: true,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r1"},
      {Typ: token.REG, Literal: "r2"},
      {Typ: token.SOP, Literal: ">>"},
      {Typ: token.NUM, Literal: "3"},
  	},
  }
	test(t, ins, "! add c r r r s n")
}

func TestInstruction4(t *testing.T) {
  ins := &Instr{
    Set: false,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.NUM, Literal: "3"},
      {Typ: token.SOP, Literal: "<<"},
      {Typ: token.NUM, Literal: "3"},
  	},
  }
	test(t, ins, "_ add c r n s n")
}

const msgErr = "Unexpected instruction hash [%s]. Expecting [%s]."

func test(t *testing.T, ins *Instr, exp string) {
  act := ins.ArgsString()
  if act != exp {
    t.Errorf(msgErr, act, exp)
  }
}
