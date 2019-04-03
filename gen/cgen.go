package gen

import (
  "strings"
	"strconv"
  
  "github.com/mhoertnagl/epic-esm/ast"
)

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
  //code     uint32
  ctx      AsmContext
  blocks   Blocks
  symVals  SymbolValidations
  symConvs SymbolConversions
  bitVals  BitValidations
  bitConvs BitConversions
}

func (g *CodeGen) Add(pattern string, template uint32) {
  
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

func (g *CodeGen) AddSymVal(param string, f SymbolValidation) {
  g.symVals[param] = f
}

func (g *CodeGen) GetSymVal(param string) SymbolValidation {
  f, ok := g.symVals[param]
  if !ok {
    return symValID
  }
  return f
}

func (g *CodeGen) AddSymConv(param string, f SymbolConversion) {
  g.symConvs[param] = f
}

func (g *CodeGen) GetSymConv(param string) SymbolConversion {
  f, ok := g.symConvs[param]
  if !ok {
    return symConvID
  }
  return f
}

func (g *CodeGen) AddBitVal(param string, f BitValidation) {
  g.bitVals[param] = f
}

func (g *CodeGen) GetBitVal(param string) BitValidation {
  f, ok := g.bitVals[param]
  if !ok {
    return bitValID
  }
  return f
}

func (g *CodeGen) AddBitConv(param string, f BitConversion) {
  g.bitConvs[param] = f
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
    //0, 
    ctx, 
    Blocks{},
    SymbolValidations{},
    SymbolConversions{},
    BitValidations{},
    BitConversions{},
  }
  
  // rd -> r
  // ra -> r 
  // rb -> r 
  // u12 -> n
  // s12 -> n 
  // @26 -> @ 
  // usw.
  
  g.AddSymConv("c", conditionConversion())
  g.AddBitConv("c", placementConversion(3, 26))
  
  g.AddSymConv("rd", registerNameConversion())
  g.AddBitConv("rd", placementConversion(4, 20))

  g.AddSymConv("ra", registerNameConversion())
  g.AddBitConv("ra", placementConversion(4, 16))
  
  g.AddSymConv("rb", registerNameConversion())
  g.AddBitConv("rb", placementConversion(4, 12))
  
  g.AddSymConv("s12", numberConversion(-4095, 4096))
  g.AddBitConv("s12", placementConversion(12, 4))
  
  g.AddSymConv("u12", numberConversion(0, 8192))
  g.AddBitConv("u12", placementConversion(12, 4))
  
  g.AddSymConv("@25", branchLabelConversion())
  g.AddBitVal("@25", branchDistanceValidation())
  g.AddBitConv("@25", placementConversion(25, 0))
  
  g.Add("_ add c rd ra rb",  0x00000000)
  g.Add("_ add c rd ra u12", 0x01000000)
  g.Add("! add c rd ra rb",  0x02000000)
  g.Add("! add c rd ra u12", 0x03000000)
  
  g.Add("_ bra c @25",       0xe0000000)
  g.Add("_ brl c @25",       0xe2000000)
    
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
  for idx, param := range params[2:] {
    sym := ins.Args[idx - 2].Literal
    g.GetSymVal(param)(g.ctx, sym)
    bits := g.GetSymConv(param)(g.ctx, sym)
    g.GetBitVal(param)(g.ctx, bits)
    bits = g.GetBitConv(param)(g.ctx, bits)
    code |= uint32(bits)
  }
  return code
}

// func (g *CodeGen) Emit() uint32 {
//   return g.code
// }
// 
// func (g *CodeGen) PlaceDataCmd(cmd string) {
// 	code, ok := dataInstructions[cmd]
// 	if !ok {
// 		g.ctx.Error("Unrecognized instruction [%s].", cmd)
// 	}
// 	g.code |= code
// }
// 
// func (g *CodeGen) PlaceMemCmd(cmd string) {
// 	code, ok := memInstructions[cmd]
// 	if !ok {
// 		g.ctx.Error("Unrecognized instruction [%s].", cmd)
// 	}
// 	g.code |= code
// }
// 
// func (g *CodeGen) PlaceBranchCmd(cmd string) {
// 	code, ok := branchInstructions[cmd]
// 	if !ok {
// 		g.ctx.Error("Unrecognized instruction [%s].", cmd)
// 	}
// 	g.code |= code
// }
// 
// func (g *CodeGen) PlaceCnd(cnd string) {
// 	code, ok := conditions[cnd]
// 	if !ok {
// 		g.ctx.Error("Unrecognized condition flag [%s].", cnd)
// 	}
// 	g.code |= g.place(int64(code), 26, 3)
// }
// 
// func (g *CodeGen) PlaceSetBit(set bool) {  
// 	if set {
// 		g.code |= g.place(1, 25, 1)
// 	}
// }
// 
// func (g *CodeGen) PlaceI16Bit() {
// 	g.code |= g.place(1, 29, 1)
// }
// 
// func (g *CodeGen) PlaceI12Bit() {
// 	g.code |= g.place(1, 24, 1)
// }
// 
// func (g *CodeGen) PlaceRd(rdName string) {
// 	rd, ok := registers[rdName]
// 	if !ok {
// 		g.ctx.Error("unrecognized destination register [%s]", rdName)
// 	}
// 	g.code |= g.place(int64(rd), 20, 4)
// }
// 
// func (g *CodeGen) PlaceRa(raName string) {
// 	ra, ok := registers[raName]
// 	if !ok {
// 		g.ctx.Error("unrecognized source A register [%s]", raName)
// 	}
// 	g.code |= g.place(int64(ra), 16, 4)
// }
// 
// func (g *CodeGen) PlaceRb(rbName string) {
// 	rb, ok := registers[rbName]
// 	if !ok {
// 		g.ctx.Error("unrecognized source B register [%s]", rbName)
// 	}
// 	g.code |= g.place(int64(rb), 12, 4)
// }

// func (g *CodeGen) PlaceNumShift(sh *NumShift) uint32 {
// 	if sh == nil {
// 		return 0
// 	}
// 	code := g.placeShiftOp(sh.cmd)
// 	// Turns a Rotate Right (<>>) into a Rotate Left (<<>). The following
// 	// identity holds for all cases: x <>> n <--> x <<> (32 - n)
// 	if sh.cmd == "<>>" || sh.cmd == "ror" {
// 		shft := g.convertUnsignedNum(sh.num, 0, 5)
// 		code |= g.place(int64(32-shft), 4, 5)
// 	} else {
// 		code |= g.convertUnsignedNum(sh.num, 4, 5)
// 	}
// 	return code
// }
// 
// func (g *CodeGen) placeShiftOp(cmd string) uint32 {
// 	sop, ok := shiftOps[cmd]
// 	if !ok {
// 		g.Error("unrecognized shift operator [%s]", cmd)
// 	}
// 	return g.place(int64(sop), 2, 9)
// }
// 
// func (g *CodeGen) PlaceBranchAddress(label string) {
// 	sym, ok := g.ctx.FindSymbol(label)
// 	if !ok {
// 		g.ctx.Error("Reference to undefined symbol [%s].", label)
// 	}
// 	g.code |= g.convertAddr(sym.addr)
// }
// 
// func (g *CodeGen) convertSignedNum(n string, s uint8, p uint8) uint32 {
// 	return g.convertNum(n, s, p, -(1 << p), 1 << p)
// }
// 
// func (g *CodeGen) convertUnsignedNum(n string, s uint8, p uint8) uint32 {
// 	return g.convertNum(n, s, p, 0, 1 << p)
// }
// 
// func (g *CodeGen) convertNum(n string, s uint8, p uint8, min int64, max int64) uint32 {
// 	i, err := g.parseNum(n)
// 
// 	if err != nil {
// 		g.ctx.Error("Number [%s] too long.", n)
// 	}
// 	if i < min {
// 		g.ctx.Error("Unexpected number [%s]. Number must be greater than [%d].", n, min)
// 	}
// 	if i >= max {
// 		g.ctx.Error("Unexpected number [%s]. Number must be less than [%d]", n, max)
// 	}
// 	return g.place(i, s, p)
// }
// 
// func (g *CodeGen) convertAddr(addr uint32) uint32 {
// 	bra := int64(addr - g.ctx.Ip())
// 	if bra < BRA_MIN || bra >= BRA_MAX {
// 		g.ctx.Error("Branch distance [%d] too large.", bra)
// 	}
// 	return g.place(bra, 0, 25)
// }
// 
// func (g *CodeGen) parseNum(n string) (int64, error) {
// 	// strings.HasPrefix
// 	if len(n) > 2 && n[0:2] == "0b" {
// 		return strconv.ParseInt(n[2:], 2, 32)
// 	}
// 	if len(n) > 2 && n[0:2] == "0x" {
// 		return strconv.ParseInt(n[2:], 16, 32)
// 	}
// 	return strconv.ParseInt(n, 10, 32)
// }
// 
// func (g *CodeGen) place(i int64, s uint8, p uint8) uint32 {
// 	return uint32((i & ((1 << p) - 1)) << s)
// }
