package gen

import (
  "fmt"
  
  "github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/ast"
)

type Generator func(ins *ast.Instr) []uint32

type InstrGen struct {
  ctx  AsmContext
  cgen *CodeGen
}

// TODO: Builder pattern?

func NewInstrGen(ctx  AsmContext) *InstrGen {
  g := &InstrGen{
    ctx: ctx,
    cgen: NewCodeGen(ctx),
  }
  return g
}

func (g *InstrGen) Generate(ins *ast.Instr) []uint32 {
  return []uint32{}
}

func (g *InstrGen) gen(set bool, cmd string, cond string, args ...token.Token) []uint32 {
  ins := &ast.Instr{
    Set: set,
    Cmd: cmd,
    Cond: cond,
    Args: args,
  }
  return g.Generate(ins)
}

func (g *InstrGen) genOne(set bool, cmd string, cond string, args ...token.Token) uint32 {
  ins := &ast.Instr{
    Set: set,
    Cmd: cmd,
    Cond: cond,
    Args: args,
  }
  return g.cgen.Generate(ins)
}

func (g *InstrGen) nop() uint32 {
  // The nop instruction (addnv r0 r0 r0) is all zero. 
  return 0
}

func (g *InstrGen) clr(set bool, cond string, rd string) uint32 {
  return g.genOne(set, "xor", cond, reg(rd), reg(rd), reg(rd))
}

func (g *InstrGen) inv(set bool, cond string, rd string) uint32 {
  return g.genOne(set, "nor", cond, reg(rd), reg(rd), reg(rd))
}

func (g *InstrGen) neg(set bool, cond string, rd string) uint32 {
  return g.neg2(set, cond, rd, rd)
}

func (g *InstrGen) neg2(set bool, cond string, rd string, ra string) uint32 {
  return g.genOne(set, "mul", cond, reg(rd), reg(ra), numi(-1))
}

func (g *InstrGen) ret(set bool, cond string) uint32 {
  return g.ret1(set, cond, "rp")
}

func (g *InstrGen) ret1(set bool, cond string, ra string) uint32 {
  return g.genOne(set, "mov", cond, reg("ip"), reg(ra))
}

func (g *InstrGen) ldaGenerator(ins *ast.Instr) []uint32 {
  rd := ins.Args[0].Literal
  lbl := ins.Args[1].Literal
  return g.lda(ins.Set, ins.Cond, rd, lbl)
}

func (g *InstrGen) lda(set bool, cond string, rd string, lbl string) []uint32 {
  sym, ok := g.ctx.FindSymbol(lbl)
  if !ok {
    g.ctx.Error("Undefinded symbol [%s].", lbl)
  }
  return g.ldc(set, cond, rd, sym.addr)
}

func (g *InstrGen) ldcGenerator(ins *ast.Instr) []uint32 {
  rd := ins.Args[0].Literal
  ns := ins.Args[1].Literal
  n, err := parseNum(ns)
  if err != nil {
    g.ctx.Error("Is not a number [%s].", ns)
  }
  return g.ldc(ins.Set, ins.Cond, rd, uint32(n))
}

func (g *InstrGen) ldc(set bool, cond string, rd string, n uint32) []uint32 {
  codes := []uint32{}
  
  if n == 0 {
    // code := g.genOne(set, "clr", cond, reg(rd))
    // code := g.genOne(set, "xor", cond, reg(rd), reg(rd), reg(rd))
    code := g.clr(set, cond, rd)
    codes = append(codes, code)
    return codes
  }
  
  // if n > 0 {
  //   return codes
  // }
  
  nu := n >> 16
  if (nu > 0) {
    upper := g.genOne(false, "ldc", cond, reg(rd), numu(nu), sop("<<"), numu(16))
    codes = append(codes, upper)
  }
  
  nl := n & 0xFFFF
  if (nl > 0) {
    lower := g.genOne(set, "ldc", cond, reg(rd), numu(nl))
    codes = append(codes, lower)
  }
  return codes
}

func reg(r string) token.Token {
  return token.Token{ Typ: token.REG, Literal: r }
}

func sop(s string) token.Token {
  return token.Token{ Typ: token.SOP, Literal: s }
}

func numi(n int32) token.Token {
  return token.Token{ Typ: token.NUM, Literal: fmt.Sprintf("%d", n) }
}

func numu(n uint32) token.Token {
  return token.Token{ Typ: token.NUM, Literal: fmt.Sprintf("%d", n) }
}

func nums(n string) token.Token {
  return token.Token{ Typ: token.NUM, Literal: n }
}
