package gen

import (
	//"fmt"
	"strconv"
	"strings"

	"github.com/mhoertnagl/epic-esm/ast"
)

type ParamHashs map[string]string

type Block struct {
	Pattern  string
	Template uint32
}

type SymbolValidation func(ctx AsmContext, param string, arg string) bool
type SymbolConversion func(ctx AsmContext, param string, arg string) int32
type BitValidation func(ctx AsmContext, param string, arg int32) bool
type BitConversion func(ctx AsmContext, param string, arg int32) int32

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
		Pattern:  pattern,
		Template: template,
	}
}

func symValID(ctx AsmContext, param string, val string) bool {
	return true
}

func symConvID(ctx AsmContext, param string, val string) int32 {
	//ctx.Error("Missing symbol converter for argument [%s].", val)
	return 0
}

func bitValID(ctx AsmContext, param string, val int32) bool {
	return true
}

func bitConvID(ctx AsmContext, param string, val int32) int32 {
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

func exactMatchValidation() SymbolValidation {
	return func(ctx AsmContext, param string, arg string) bool {
		if param != arg {
			ctx.Error("Unexpected token [%s]. Expected [%s]", arg, param)
			return false
		}
		return true
	}
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
	return func(ctx AsmContext, param string, rx string) int32 {
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
	return func(ctx AsmContext, param string, cnd string) int32 {
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
	return func(ctx AsmContext, param string, sop string) int32 {
		code, ok := shiftOps[sop]
		if !ok {
			ctx.Error("Unrecognized shift operator [%s].", sop)
		}
		return code
	}
}

func numberConversion(min int64, max int64) SymbolConversion {
	return func(ctx AsmContext, param string, num string) int32 {
		i, err := parseNum(num)

		if err != nil {
			ctx.Error("Value [%s] is not a number.", num)
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

func invNumberConversion(max int64) SymbolConversion {
	return func(ctx AsmContext, param string, num string) int32 {
		i, err := parseNum(num)

		if err != nil {
			ctx.Error("Value [%s] is not a number.", num)
		}
		if i < 0 {
			ctx.Error("Unexpected number [%s]. Number must be greater than [%d].", num, 0)
		}
		if i >= max {
			ctx.Error("Unexpected number [%s]. Number must be less than [%d].", num, max)
		}
		return int32(max - i)
	}
}

func parseNum(n string) (int64, error) {
	// strings.HasPrefix
	// if len(n) > 2 && n[0:2] == "0b" {
	// 	return strconv.ParseInt(n[2:], 2, 32)
	// }
	if len(n) > 2 && n[0:2] == "0x" {
		return strconv.ParseInt(n[2:], 16, 64)
	}
	return strconv.ParseInt(n, 10, 64)
}

func branchLabelConversion() SymbolConversion {
	return func(ctx AsmContext, param string, lbl string) int32 {
		sym, ok := ctx.FindSymbol(lbl)
		if !ok {
			ctx.Error("")
			return 0
		}
		return int32(sym.addr - ctx.Ip())
	}
}

const (
	braLen = 25 // bits
	braMin = -(1 << (braLen - 1))
	braMax = 1 << (braLen - 1)
)

func branchDistanceValidation() BitValidation {
	return func(ctx AsmContext, param string, bra int32) bool {
		if bra < braMin || bra >= braMax {
			ctx.Error("Branch distance [%d] too large.", bra)
			return false
		}
		return true
	}
}

func rangeValidation(min int32, max int32) BitValidation {
	return func(ctx AsmContext, param string, val int32) bool {
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
	return func(ctx AsmContext, param string, val int32) int32 {
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

	g.AddParamHash("rd", "r")
	g.AddParamHash("ra", "r")
	g.AddParamHash("rb", "r")
	g.AddParamHash("u5", "n")
	g.AddParamHash("s12", "n")
	g.AddParamHash("u12", "n")
	g.AddParamHash("s16", "n")
	g.AddParamHash("u16", "n")
	g.AddParamHash("@25", "@")

	// For 16-bit immediate instructions that address the upper halfword of a
	// register, e.g. oor r0 0xFEDC << 16.
	// The parser will eject an instruction that will hash to '_ oor r n s n'.
	// This will create a matching block entry for instruction templates ending
	// with '<< 16'. An exact match SymbolValidation will then check the argument
	// and the parameter for exact equivalence.
	g.AddParamHash("<<", "s")
	g.AddParamHash("16", "n")

	// For memory instructions to denote a memory location, e.g.
	// stw r0 r1[-4].
	g.AddSymVal("[", exactMatchValidation())
	g.AddSymVal("]", exactMatchValidation())

	// For 16-bit immediate instructions that address the upper halfword of a
	// register, e.g. oor r0 0xFEDC << 16.
	g.AddSymVal("<<", exactMatchValidation())
	g.AddSymVal("16", exactMatchValidation())

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

	g.AddSymConv("u5", numberConversion(0, 32))
	g.AddBitConv("u5", placementConversion(5, 4))

	// Parameter i5 will subtract the 5-bit number from the least upper bound 32.
	g.AddSymConv("i5", invNumberConversion(32))
	g.AddBitConv("i5", placementConversion(5, 4))

	g.AddSymConv("s12", numberConversion(-4096, 4096))
	g.AddBitConv("s12", placementConversion(12, 4))

	g.AddSymConv("u12", numberConversion(0, 8192))
	g.AddBitConv("u12", placementConversion(12, 4))

	g.AddSymConv("s16", numberConversion(-32768, 32768))
	g.AddBitConv("s16", placementConversion(16, 4))

	g.AddSymConv("u16", numberConversion(0, 65536))
	g.AddBitConv("u16", placementConversion(16, 4))

	g.AddSymConv("@25", branchLabelConversion())
	g.AddBitVal("@25", branchDistanceValidation())
	g.AddBitConv("@25", placementConversion(25, 0))

	// TODO: rda - places the same register at rd and ra

	g.Add("_ add c rd ra rb s u5", 0x00000000)
	g.Add("_ add c rd ra rb", 0x00000000)
	g.Add("_ add c rd ra u12", 0x01000000)
	g.Add("! add c rd ra rb s u5", 0x02000000)
	g.Add("! add c rd ra rb", 0x02000000)
	g.Add("! add c rd ra u12", 0x03000000)

	g.Add("_ sub c rd ra rb s u5", 0x00000001)
	g.Add("_ sub c rd ra rb", 0x00000001)
	g.Add("_ sub c rd ra u12", 0x01000001)
	g.Add("! sub c rd ra rb s u5", 0x02000001)
	g.Add("! sub c rd ra rb", 0x02000001)
	g.Add("! sub c rd ra u12", 0x03000001)

	g.Add("_ mul c rd ra rb s u5", 0x00000002)
	g.Add("_ mul c rd ra rb", 0x00000002)
	g.Add("_ mul c rd ra s12", 0x01000002)
	g.Add("! mul c rd ra rb s u5", 0x02000002)
	g.Add("! mul c rd ra rb", 0x02000002)
	g.Add("! mul c rd ra s12", 0x03000002)

	g.Add("_ div c rd ra rb s u5", 0x00000003)
	g.Add("_ div c rd ra rb", 0x00000003)
	g.Add("_ div c rd ra s12", 0x01000003)
	g.Add("! div c rd ra rb s u5", 0x02000003)
	g.Add("! div c rd ra rb", 0x02000003)
	g.Add("! div c rd ra s12", 0x03000003)

	g.Add("_ and c rd ra rb s u5", 0x00000004)
	g.Add("_ and c rd ra rb", 0x00000004)
	g.Add("_ and c rd ra u12", 0x01000004)
	g.Add("! and c rd ra rb s u5", 0x02000004)
	g.Add("! and c rd ra rb", 0x02000004)
	g.Add("! and c rd ra u12", 0x03000004)

	g.Add("_ oor c rd ra rb s u5", 0x00000005)
	g.Add("_ oor c rd ra rb", 0x00000005)
	g.Add("_ oor c rd ra u12", 0x01000005)
	g.Add("! oor c rd ra rb s u5", 0x02000005)
	g.Add("! oor c rd ra rb", 0x02000005)
	g.Add("! oor c rd ra u12", 0x03000005)

	g.Add("_ xor c rd ra rb s u5", 0x00000006)
	g.Add("_ xor c rd ra rb", 0x00000006)
	g.Add("_ xor c rd ra u12", 0x01000006)
	g.Add("! xor c rd ra rb s u5", 0x02000006)
	g.Add("! xor c rd ra rb", 0x02000006)
	g.Add("! xor c rd ra u12", 0x03000006)

	g.Add("_ nor c rd ra rb s u5", 0x00000007)
	g.Add("_ nor c rd ra rb", 0x00000007)
	g.Add("_ nor c rd ra u12", 0x01000007)
	g.Add("! nor c rd ra rb s u5", 0x02000007)
	g.Add("! nor c rd ra rb", 0x02000007)
	g.Add("! nor c rd ra u12", 0x03000007)

	// TODO: Hier fehlt auch noch was.
	// "adu": 0x00000008,
	// "sbu": 0x00000009,
	// //"mlu": 0x0000000a, multiplikation ist immer signed
	// //"dvu": 0x0000000b, division ist immer signed

	// cmp, cpu and tst do not write any result to the register file. Register rd
	// (bits 23-20) must be 0b0000 to guarantee future extensibility.
	g.Add("_ cmp c ra rb s u5", 0x0000000C)
	g.Add("_ cmp c ra rb", 0x0000000C)
	g.Add("_ cmp c ra s12", 0x0100000C)

	g.Add("_ cpu c ra rb s u5", 0x0000000D)
	g.Add("_ cpu c ra rb", 0x0000000D)
	g.Add("_ cpu c ra u12", 0x0100000D)

	g.Add("_ tst c ra rb s u5", 0x0000000E)
	g.Add("_ tst c ra rb", 0x0000000E)
	g.Add("_ tst c ra u12", 0x0100000E)

	// TODO: Unsigned or signed?
	g.Add("_ mov c rd rb s u5", 0x0000000F)
	g.Add("_ mov c rd rb", 0x0000000F)
	g.Add("_ mov c rd u12", 0x0100000F)
	g.Add("! mov c rd rb s u5", 0x0200000F)
	g.Add("! mov c rd rb", 0x0200000F)
	g.Add("! mov c rd u12", 0x0300000F)

	// These dedicated shift instructions are mere move instructions in disguise.
	g.Add("_ sll c rd rb u5", 0x0000000F)
	g.Add("! sll c rd rb u5", 0x0200000F)

	g.Add("_ srl c rd rb u5", 0x0000020F)
	g.Add("! srl c rd rb u5", 0x0200020F)

	g.Add("_ sra c rd rb u5", 0x0000040F)
	g.Add("! sra c rd rb u5", 0x0200040F)

	g.Add("_ rol c rd rb u5", 0x0000060F)
	g.Add("! rol c rd rb u5", 0x0200060F)

	// Rotate right is a pseudo instruction defined as:
	//   ror rd rb x := rol rd rb (32 - x)
	g.Add("_ ror c rd rb i5", 0x0000060F)
	g.Add("! ror c rd rb i5", 0x0200060F)

	g.Add("_ add c rd u16 << 16", 0x21000000)
	g.Add("_ add c rd u16", 0x20000000)
	g.Add("! add c rd u16 << 16", 0x23000000)
	g.Add("! add c rd u16", 0x22000000)
	// TODO: Hier fehlt doch was :)
	g.Add("_ mov c rd u16 << 16", 0x2100000F)
	g.Add("_ mov c rd u16", 0x2000000F)
	g.Add("! mov c rd u16 << 16", 0x2300000F)
	g.Add("! mov c rd u16", 0x2200000F)

	g.Add("_ stw c rd ra [ rb s u5 ]", 0x40000000)
	g.Add("_ stw c rd ra [ rb ]", 0x40000000)
	g.Add("_ stw c rd ra [ s12 ]", 0x41000000)
	g.Add("! stw c rd ra [ rb s u5 ]", 0x42000000)
	g.Add("! stw c rd ra [ rb ]", 0x42000000)
	g.Add("! stw c rd ra [ s12 ]", 0x43000000)

	g.Add("_ ldw c rd ra [ rb s u5 ]", 0x40000001)
	g.Add("_ ldw c rd ra [ rb ]", 0x40000001)
	g.Add("_ ldw c rd ra [ s12 ]", 0x41000001)
	g.Add("! ldw c rd ra [ rb s u5 ]", 0x42000001)
	g.Add("! ldw c rd ra [ rb ]", 0x42000001)
	g.Add("! ldw c rd ra [ s12 ]", 0x43000001)

	g.Add("_ bra c @25", 0xE0000000)
	g.Add("_ brl c @25", 0xE2000000)

	return g
}

func (g *CodeGen) Generate(ins *ast.Instr) uint32 {
	as := ins.ArgsString()
	blk, ok := g.blocks[as]
	if !ok {
		g.ctx.Error("Unsupported instruction [%s] with arg string [%s].", ins, as)
		return 0
	}
	code := blk.Template
	// TODO: Split the pattern once and for all in the Add(string, uint32) method.
	params := strings.Split(blk.Pattern, " ")
	// TODO: params length test. Actually this cannot happen but one never knows.
	// Skip set bit (! or _) and command.
	// Encode the condition flag.
	// TODO: Add condition to the list of arguments in the parser?
	code |= g.generateParam(params[2], ins.Cond)
	for idx, param := range params[3:] {
		arg := ins.Args[idx].Literal
		code |= g.generateParam(param, arg)
	}
	return code
}

func (g *CodeGen) generateParam(param string, arg string) uint32 {
	g.GetSymVal(param)(g.ctx, param, arg)
	bits := g.GetSymConv(param)(g.ctx, param, arg)
	g.GetBitVal(param)(g.ctx, param, bits)
	bits = g.GetBitConv(param)(g.ctx, param, bits)
	return uint32(bits)
}
