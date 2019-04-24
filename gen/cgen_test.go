package gen

import (
  //"fmt"
  "strings"
  "strconv"
	"testing"
  
  "github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/ast"
  "github.com/mhoertnagl/epic-esm/lexer"
  "github.com/mhoertnagl/epic-esm/parser"
)

func TestI16_0(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "nv",
    Args: []token.Token{
      {Typ: token.REG, Literal: "rp"},
      {Typ: token.NUM, Literal: "0xFFFF"},
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
      {Typ: token.NUM, Literal: "0xFFFF"},
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
      {Typ: token.NUM, Literal: "0xFFFF"},
      {Typ: token.SOP, Literal: "<<"},
      {Typ: token.NUM, Literal: "16"},
  	},
  }
  // 0010 0011 1110 1111  1111 1111 1111 0000
	test(t, ins, 0x23EFFFF0)
}

func TestNop(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "nv",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r0"},
  	},
  }
  // 0000 0000 0000 0000  0000 0000 0000 0000
	test(t, ins, 0x00000000)
}

// TODO: Test different execution conditions nv - al.

func TestInstruction0(t *testing.T) {
  ins := &ast.Instr{
    Set: false,
    Cmd: "add",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.REG, Literal: "r0"},
      {Typ: token.REG, Literal: "r1"},
      {Typ: token.NUM, Literal: "42"},
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
      {Typ: token.REG, Literal: "r1"},
      {Typ: token.NUM, Literal: "42"},
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
      {Typ: token.REG, Literal: "r3"},
      {Typ: token.REG, Literal: "r4"},
      {Typ: token.SOP, Literal: ">>>"},
      {Typ: token.NUM, Literal: "3"},
  	},
  }
  // 0001 1100 0010 0011  0100 0100 0011 0000
	test(t, ins, 0x1C234430)
}

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

func TestBra(t *testing.T) {
  ctx := NewAsmContext("test.esm")
  sym(ctx, "@0", 0x01EDCBA9)

  ins := &ast.Instr{
    Set: false,
    Cmd: "bra",
    Cond: "al",
    Args: []token.Token{
      {Typ: token.LBL, Literal: "@0"},
  	},
  }
  // 1111 1101 1110 1101  1100 1011 1010 1001
	testc(t, ctx, ins, 0xFDEDCBA9)
}

func TestConditions(t *testing.T) {
  ts(t, "addnv r0 r0 r0",  "0000 0000 0000 0000  0000 0000 0000 0000")
  ts(t, "addeq r0 r0 r0",  "0000 0100 0000 0000  0000 0000 0000 0000")
  ts(t, "addlt r0 r0 r0",  "0000 1000 0000 0000  0000 0000 0000 0000")
  ts(t, "addle r0 r0 r0",  "0000 1100 0000 0000  0000 0000 0000 0000")
  ts(t, "addgt r0 r0 r0",  "0001 0000 0000 0000  0000 0000 0000 0000")
  ts(t, "addge r0 r0 r0",  "0001 0100 0000 0000  0000 0000 0000 0000")
  ts(t, "addne r0 r0 r0",  "0001 1000 0000 0000  0000 0000 0000 0000")
  ts(t, "addal r0 r0 r0",  "0001 1100 0000 0000  0000 0000 0000 0000")
  ts(t, "add   r0 r0 r0",  "0001 1100 0000 0000  0000 0000 0000 0000")
}  

// ts(t, "addnv rp 0xFFFF", "0010 0000 1110 1111  1111 1111 1111 0000")

// TODO: Test branch instructions with @lbl > IP > 0.
// TODO: Test branch distance to large error.
// TODO: Test branch with IP > @lbl > 0

// TODO: Test Branch and Link.

const msgErr = "Unexpected code [%08X]. Expecting code [%08X]."

func test(t *testing.T, ins *ast.Instr, exp uint32) {
  testc(t, NewAsmContext("test.esm"), ins, exp)
}

func testc(t *testing.T, ctx AsmContext, ins *ast.Instr, exp uint32) {
  g := NewCodeGen(ctx)
  code := g.Generate(ins)
  if code != exp {
    t.Errorf(msgErr, code, exp)
  }
}

func sym(ctx AsmContext, name string, addr uint32) {
  ip := ctx.Ip()
  ctx.ResetIp()
  ctx.IncrementIp(addr)
  ctx.AddSymbol(name)
  ctx.ResetIp()
  ctx.IncrementIp(ip)
}

const msgErr2 = "[%s]: Unexpected code [%08X]. Expecting code [%08X]."

func ts(t *testing.T, ins string, exp string) {
  tsc(t, NewAsmContext("test.esm"), ins, exp)
}

func tsc(t *testing.T, ctx AsmContext, ins string, exp string) {
  l := lexer.NewLexer(ins)
  p := parser.NewParser(l)
  g := NewCodeGen(ctx)
  i := p.Parse()
  e := toNum(t, exp)
  switch ii := i.(type) {
  case *ast.Instr: 
    c := g.Generate(ii)
    if c != e {
      t.Errorf(msgErr2, ins, c, e)
    }  
    break
  default:
    t.Errorf("[%s] is not an instruction.", ins) 
    break
  }
}

func toNum(t *testing.T, exp string) uint32 {
  ns := strings.Replace(exp, " ", "", 8)
  n, err := strconv.ParseInt(ns, 2, 64)
  if err != nil {
    panic(err)
  }
  return uint32(n)
}
