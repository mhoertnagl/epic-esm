package gen

import (
  //"fmt"
  "strings"
	"strconv"
  
  "github.com/mhoertnagl/epic-esm/ast"
)

type ParamHashs map[string]string

type Block struct {
  Pattern string
  Template uint32
}

type SymbolValidation func(ctx AsmContext, val string) bool
type SymbolConversion func(ctx AsmContext, val string) int32
type BitValidation func(ctx AsmContext, val int32) bool
type BitConversion func(ctx AsmContext, val int32) int32

type Blocks map[string]*Block
type SymbolValidations map[string]SymbolValidation
type SymbolConversions map[string]SymbolConversion
type BitValidations map[string]BitValidation
type BitConversions map[string]BitConversion

type CodeGen struct {
  ctx        AsmContext
  paramHashs ParamHashs
  blocks     Blocks
  symVals    SymbolValidations
  symConvs   SymbolConversions
  bitVals    BitValidations
  bitConvs   BitConversions
}

func (g *CodeGen) AddParamHash(param string, hash string) {
  g.paramHashs[param] = hash
}

func (g *CodeGen) Add(pattern string, template uint32) {
  hashs := []string{}
  params := strings.Split(pattern, " ")
  for _, p := range params {
    h, ok := g.paramHashs[p]
    if ok {
      hashs = append(hashs, h)
    } else {
      // If there is no parameter hash defined, use the original parameter value
      // instead. 
      hashs = append(hashs, p)
    }
  }
  hash := strings.Join(hashs, " ")
  g.blocks[hash] = &Block{ 
    Pattern: pattern, 
    Template: template, 
  }
}

func symValID(ctx AsmContext, val string) bool { 
  return true 
}

func symConvID(ctx AsmContext, val string) int32 { 
  ctx.Error("Missing symbol converter for argument [%s].", val)
  return 0 
}

func bitValID(ctx AsmContext, val int32) bool { 
  return true 
}

func bitConvID(ctx AsmContext, val int32) int32 { 
  //ctx.Error("Missing bit converter for argument [%s].", val)
  return 0 
}

func (g *CodeGen) AddSymVal(param string, fun SymbolValidation) {
  g.symVals[param] = fun
}

func (g *CodeGen) GetSymVal(param string) SymbolValidation {
  f, ok := g.symVals[param]
  if !ok {
    return symValID
  }
  return f
}

func (g *CodeGen) AddSymConv(param string, fun SymbolConversion) {
  g.symConvs[param] = fun
}

func (g *CodeGen) GetSymConv(param string) SymbolConversion {
  f, ok := g.symConvs[param]
  if !ok {
    return symConvID
  }
  return f
}

func (g *CodeGen) AddBitVal(param string, fun BitValidation) {
  g.bitVals[param] = fun
}

func (g *CodeGen) GetBitVal(param string) BitValidation {
  f, ok := g.bitVals[param]
  if !ok {
    return bitValID
  }
  return f
}

func (g *CodeGen) AddBitConv(param string, fun BitConversion) {
  g.bitConvs[param] = fun
}

func (g *CodeGen) GetBitConv(param string) BitConversion {
  f, ok := g.bitConvs[param]
  if !ok {
    return bitConvID
  }
  return f
}

var registers = map[string]int32{
	"r0":  0,
	"r1":  1,
	"r2":  2,
	"r3":  3,
	"r4":  4,
	"r5":  5,
	"r6":  6,
	"r7":  7,
	"r8":  8,
	"r9":  9,
	"r10": 10,
	"r11": 11,
	"r12": 12,
	"r13": 13,
	"r14": 14,
	"r15": 15,
	"sp":  13,
	"rp":  14,
	"ip":  15,
}

func registerNameConversion() SymbolConversion {
  return func (ctx AsmContext, rx string) int32 {
    	reg, ok := registers[rx]
    	if !ok {
    		ctx.Error("Unrecognized register [%s].", rx)
    	}
      return reg
  }
}

var conditions = map[string]int32{
	"nv": 0,
	"eq": 1,
	"lt": 2,
	"le": 3,
	"gt": 4,
	"ge": 5,
	"ne": 6,
	"al": 7,
}

func conditionConversion() SymbolConversion {
  return func (ctx AsmContext, cnd string) int32 {
    	code, ok := conditions[cnd]
    	if !ok {
    		ctx.Error("Unrecognized condition flag [%s].", cnd)
    	}
      return code
  }
}

var shiftOps = map[string]int32{
	"<<":  0,
	">>":  1,
	">>>": 2,
	"<<>": 3,
	"<>>": 3,
}

func sopConversion() SymbolConversion {
  return func (ctx AsmContext, sop string) int32 {
    	code, ok := shiftOps[sop]
    	if !ok {
    		ctx.Error("Unrecognized shift operator [%s].", sop)
    	}
      return code
  }
}

func numberConversion(min int64, max int64) SymbolConversion {
  return func (ctx AsmContext, num string) int32 {
    i, err := parseNum(num)

  	if err != nil {
  		ctx.Error("Number [%s] too long.", num)
  	}
  	if i < min {
  		ctx.Error("Unexpected number [%s]. Number must be greater than [%d].", num, min)
  	}
  	if i >= max {
  		ctx.Error("Unexpected number [%s]. Number must be less than [%d].", num, max)
  	}
    return int32(i)
  }
}

func parseNum(n string) (int64, error) {
	// strings.HasPrefix
	if len(n) > 2 && n[0:2] == "0b" {
		return strconv.ParseInt(n[2:], 2, 32)
	}
	if len(n) > 2 && n[0:2] == "0x" {
		return strconv.ParseInt(n[2:], 16, 32)
	}
	return strconv.ParseInt(n, 10, 32)
}

func branchLabelConversion() SymbolConversion {
  return func (ctx AsmContext, lbl string) int32 {
    sym, ok := ctx.FindSymbol(lbl)
    if !ok {
      ctx.Error("")
      return 0
    }
    return int32(sym.addr - ctx.Ip());
  }
}

const (
	BRA_MIN = -(1 << 24)
	BRA_MAX = 1 << 24
)

func branchDistanceValidation() BitValidation {
  return func (ctx AsmContext, bra int32) bool {
    if bra < BRA_MIN || bra >= BRA_MAX {
    	ctx.Error("Branch distance [%d] too large.", bra)
      return false
    }
    return true
  }
}


func rangeValidation(min int32, max int32) BitValidation {
  return func (ctx AsmContext, val int32) bool {
    if val < min {
      ctx.Error("")
      return false
    }
    if val > max {
      ctx.Error("")
      return false      
    }
    return true
  }
}

func placementConversion(p uint8, s uint8) BitConversion {
	return func (ctx AsmContext, val int32) int32 {
    return (val & ((1 << p) - 1)) << s
  }
}

func NewCodeGen(ctx AsmContext) *CodeGen {
 	g := &CodeGen{
    ctx, 
    ParamHashs{},
    Blocks{},
    SymbolValidations{},
    SymbolConversions{},
    BitValidations{},
    BitConversions{},
  }
  
  g.AddParamHash("rd",  "r")
  g.AddParamHash("ra",  "r")
  g.AddParamHash("rb",  "r")
  g.AddParamHash("u5",  "n")
  g.AddParamHash("s12", "n")
  g.AddParamHash("u12", "n")
  g.AddParamHash("s16", "n")
  g.AddParamHash("u16", "n")
  g.AddParamHash("@25", "@")
  
  g.AddSymConv("c", conditionConversion())
  g.AddBitConv("c", placementConversion(3, 26))
  
  g.AddSymConv("rd", registerNameConversion())
  g.AddBitConv("rd", placementConversion(4, 20))

  g.AddSymConv("ra", registerNameConversion())
  g.AddBitConv("ra", placementConversion(4, 16))
  
  g.AddSymConv("rb", registerNameConversion())
  g.AddBitConv("rb", placementConversion(4, 12))
  
  g.AddSymConv("s", sopConversion())
  g.AddBitConv("s", placementConversion(2, 9))
  
  g.AddSymConv("u5", numberConversion(0, 31))
  g.AddBitConv("u5", placementConversion(2, 4))
  
  g.AddSymConv("s12", numberConversion(-4096, 4095))
  g.AddBitConv("s12", placementConversion(12, 4))
  
  g.AddSymConv("u12", numberConversion(0, 8192))
  g.AddBitConv("u12", placementConversion(12, 4))
  
  g.AddSymConv("s16", numberConversion(-32768, 32767))
  g.AddBitConv("s16", placementConversion(16, 4))
  
  g.AddSymConv("u16", numberConversion(0, 65535))
  g.AddBitConv("u16", placementConversion(16, 4))
  
  g.AddSymConv("@25", branchLabelConversion())
  g.AddBitVal("@25", branchDistanceValidation())
  g.AddBitConv("@25", placementConversion(25, 0))
  
  g.Add("_ add c rd ra rb s u5",     0x00000000)
  g.Add("_ add c rd ra rb",          0x00000000)
  g.Add("_ add c rd ra u12",         0x01000000)
  g.Add("_ add c rd u16",            0x20000000)
  g.Add("! add c rd ra rb s u5",     0x02000000)
  g.Add("! add c rd ra rb",          0x02000000)
  g.Add("! add c rd ra u12",         0x03000000)
  g.Add("! add c rd u16",            0x22000000)

  g.Add("_ sub c rd ra rb s u5",     0x00000001)
  g.Add("_ sub c rd ra rb",          0x00000001)
  g.Add("_ sub c rd ra u12",         0x01000001)
  g.Add("_ sub c rd u16",            0x20000001)
  g.Add("! sub c rd ra rb s u5",     0x02000001)
  g.Add("! sub c rd ra rb",          0x02000001)
  g.Add("! sub c rd ra u12",         0x03000001)
  g.Add("! sub c rd u16",            0x22000001)
  
  // "mul": 0x00000002,
  // "div": 0x00000003,

  // "and": 0x00000004,
  // "oor": 0x00000005,
  // "xor": 0x00000006,
  // "nor": 0x00000007,

  // "adu": 0x00000008,
  // "sbu": 0x00000009,
  // //"mlu": 0x0000000a, multiplikation ist immer signed
  // //"dvu": 0x0000000b, division ist immer signed

  g.Add("_ cmp c ra rb s u5",        0x0000000c)
  g.Add("_ cmp c ra rb",             0x0000000c)
  g.Add("_ cmp c ra s12",            0x0100000c)
  g.Add("! cmp c ra rb s u5",        0x0200000c)
  g.Add("! cmp c ra rb",             0x0200000c)
  g.Add("! cmp c ra s12",            0x0300000c)
  // "cmp": 0x0000000c,
  
  // "cpu": 0x0000000d,
  // "tst": 0x0000000e,
  // "mov": 0x0000000f,
  
  // "sll": 0x0000000f,
  // "srl": 0x0000020f,
  // "sra": 0x0000040f,
  // "rol": 0x0000060f,
  // "ror": 0x0000060f,
  
  g.Add("_ stw c rd ra [ rb s u5 ]", 0x40000000)
  g.Add("_ stw c rd ra [ rb ]",      0x40000000)
  g.Add("_ stw c rd ra [ s12 ]",     0x41000000)
  g.Add("! stw c rd ra [ rb s u5 ]", 0x42000000)
  g.Add("! stw c rd ra [ rb ]",      0x42000000)
  g.Add("! stw c rd ra [ s12 ]",     0x43000000)
  
  g.Add("_ ldw c rd ra [ rb s u5 ]", 0x40000001)
  g.Add("_ ldw c rd ra [ rb ]",      0x40000001)
  g.Add("_ ldw c rd ra [ s12 ]",     0x41000001)
  g.Add("! ldw c rd ra [ rb s u5 ]", 0x42000001)
  g.Add("! ldw c rd ra [ rb ]",      0x42000001)
  g.Add("! ldw c rd ra [ s12 ]",     0x43000001)
  
  g.Add("_ bra c @25",               0xE0000000)
  g.Add("_ brl c @25",               0xE2000000)
    
  return g
}

func (g *CodeGen) Generate(ins *ast.Instr) uint32 {
  as := ins.ArgsString()
  blk, ok := g.blocks[as]
  if !ok {
    g.ctx.Error("Unsupported arguments [%s] for command [%s].", as, ins.Cmd)
    return 0
  }
  code := blk.Template
  params := strings.Split(blk.Pattern, " ")
  // Skip set bit (! or _) and command.
  // Encode the condition flag.
  // TODO: params length test.
  code |= g.generateParam(params[2], ins.Cond)
  for idx, param := range params[3:] {    
    arg := ins.Args[idx].Literal
    code |= g.generateParam(param, arg)
  }
  return code
}

func (g *CodeGen) generateParam(param string, arg string) uint32 {
    g.GetSymVal(param)(g.ctx, arg)
    bits := g.GetSymConv(param)(g.ctx, arg)
    g.GetBitVal(param)(g.ctx, bits)
    bits = g.GetBitConv(param)(g.ctx, bits)
    return uint32(bits)
}
