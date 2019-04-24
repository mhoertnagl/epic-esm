package gen

import (
  "fmt"
  
  "github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/ast"
)

// TODO: Add argument sanity checks.
//       Perhaps like cgen? Probably overkill.

type Expansion func(ctx AsmContext, ins *ast.Instr) ast.Instrs

type Expansions map[string]Expansion

type InstrGen struct {
  ctx  AsmContext
  exps Expansions
}

func NewInstrGen(ctx  AsmContext) *InstrGen {
  g := &InstrGen{
    ctx: ctx,
    exps: Expansions{},
  }
  
  g.Add("nop", g.nopExp)
  g.Add("clr", g.clrExp)
  g.Add("inv", g.invExp)
  g.Add("neg", g.negExp)
  g.Add("ldc", g.ldcExp)
  g.Add("lda", g.ldaExp)
  g.Add("ret", g.retExp)
  
  return g
}

func (g *InstrGen) Add(cmd string, exp Expansion) {
  g.exps[cmd] = exp
}

func (g *InstrGen) GetExpanson(cmd string) Expansion {
  exp, ok := g.exps[cmd]
  if ok {
    return exp
  }
  return expID
}

func (g *InstrGen) Generate(ins *ast.Instr) ast.Instrs {
  return g.GetExpanson(ins.Cmd)(g.ctx, ins)
}

func expID(ctx AsmContext, ins *ast.Instr) ast.Instrs{ 
  return ast.Instrs{ ins }
}

func (g *InstrGen) instr(set bool, cmd string, cond string, args ...token.Token) *ast.Instr {
  return &ast.Instr{
    Set: set,
    Cmd: cmd,
    Cond: cond,
    Args: args,
  }
}

func (g *InstrGen) nopExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  // TODO: Define as constant.
  // The nop instruction (addnv r0 r0 r0) is all zero. 
  return ast.Instrs { 
    g.instr(false, "add", "nv", reg("r0"), reg("r0"), reg("r0")),
  }
}

func (g *InstrGen) clrExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  // TODO: Type constraints like in cgen.
  rd := ins.Args[0].Literal
  return g.clr(ins.Set, ins.Cond, rd)
}

func (g *InstrGen) clr(set bool, cond string, rd string) ast.Instrs {
  return ast.Instrs { 
    g.instr(set, "xor", cond, reg(rd), reg(rd), reg(rd)),
  }
}

func (g *InstrGen) invExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  // TODO: Type constraints like in cgen.
  rd := ins.Args[0].Literal
  ra := ins.Args[1].Literal
  return ast.Instrs { 
    g.instr(ins.Set, "nor", ins.Cond, reg(rd), reg(ra), reg(ra)),
  }
}

func (g *InstrGen) negExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  // TODO: Type constraints like in cgen.
  rd := ins.Args[0].Literal
  ra := ins.Args[1].Literal
  return ast.Instrs { 
    g.instr(ins.Set, "mul", ins.Cond, reg(rd), reg(ra), numi(-1)),
  }
}

func (g *InstrGen) retExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  return ast.Instrs { 
    g.instr(ins.Set, "mov", ins.Cond, reg("ip"), reg("rp")),
  }
}

func (g *InstrGen) ldaExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  rd := ins.Args[0].Literal
  lbl := ins.Args[1].Literal
  return g.lda(ins.Set, ins.Cond, rd, lbl)
}

func (g *InstrGen) lda(set bool, cond string, rd string, lbl string) ast.Instrs {
  sym, ok := g.ctx.FindSymbol(lbl)
  if !ok {
    g.ctx.Error("Undefinded symbol [%s].", lbl)
  }
  
  nu := sym.addr >> 16
  nl := sym.addr & 0xFFFF
  
  return ast.Instrs { 
    g.instr(false, "ldh", cond, reg(rd), numu(nu), sop("<<"), numu(16)),
    g.instr(set, "ldh", cond, reg(rd), numu(nl)),
  }
}

func (g *InstrGen) ldcExp(ctx AsmContext, ins *ast.Instr) ast.Instrs {
  rd := ins.Args[0].Literal
  ns := ins.Args[1].Literal
  n, err := parseNum(ns)
  if err != nil {
    g.ctx.Error("Is not a number [%s].", ns)
  }
  return g.ldc(ins.Set, ins.Cond, rd, uint32(n))
}

func (g *InstrGen) ldc(set bool, cond string, rd string, n uint32) ast.Instrs {
  instrs := ast.Instrs{}
  
  if n == 0 {
    return g.clr(set, cond, rd)
  }

  nu := n >> 16
  nl := n & 0xFFFF
    
  if (nu > 0) {
    upper := g.instr(false, "ldh", cond, reg(rd), numu(nu), sop("<<"), numu(16))
    instrs = append(instrs, upper)
  }
  
  if (nl > 0) {
    lower := g.instr(set, "ldh", cond, reg(rd), numu(nl))
    instrs = append(instrs, lower)
  }
  return instrs
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
