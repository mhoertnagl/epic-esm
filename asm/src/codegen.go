package main

import "fmt"

var dataInstructions = map[string]uint32{
	"add": 0x00000000,
	"sub": 0x00000001,
	"mul": 0x00000002,
	"div": 0x00000003,

	"and": 0x00000004,
	"oor": 0x00000005,
	"xor": 0x00000006,
	"nor": 0x00000007,

	//"tst": 0x00000008,
	"cmp": 0x00000009,
	//"add": 0x0000000a,
	//"add": 0x0000000b,

	"sll": 0x0000000c,
	"rol": 0x0000000d,
	"srl": 0x0000000e,
	"sra": 0x0000000f,
}

var immInstructions = map[string]uint32{
	"add": 0x00000000,
	"sub": 0x00000001,
	"mul": 0x00000002,
	"div": 0x00000003,

	"and": 0x00000004,
	"oor": 0x00000005,
	"xor": 0x00000006,
	"nor": 0x00000007,

	//"tst": 0x00000008,
	"cmp": 0x00000009,
	//"add": 0x0000000a,
	//"add": 0x0000000b,

	"ldc": 0x0000000c,
	//"rol": 0x0000000d,
	//"srl": 0x0000000e,
	//"sra": 0x0000000f,
}

var branchInstructions = map[string]uint32{
	"bra": 0xe0000000,
	"brl": 0xe2000000,
}

var registers = map[string]uint32{
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

var conditions = map[string]uint32{
	"nv": 0,
	"eq": 1,
	"lt": 2,
	"le": 3,
	"gt": 4,
	"ge": 5,
	"ne": 6,
	"al": 7,
}

type CodeGen struct {
	filename string
	st       SymbolTable
	ip       uint32
	lineNo   uint32
}

func NewCodeGen(filename string, st SymbolTable) *CodeGen {
	return &CodeGen{filename, st, 0, 1}
}

func (g *CodeGen) Error(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s [%d] ERROR: %s\n", g.filename, g.lineNo, msg)
}

func (g *CodeGen) Generate(node interface{}) (uint32, bool) {
	code := uint32(0)
	ok := false
	switch node.(type) {
	case *RegInstruction:
		code = g.genRegInstruction(node.(*RegInstruction))
		ok = true
		g.ip++
		break
	case *I12Instruction:
		code = g.genI12Instruction(node.(*I12Instruction))
		ok = true
		g.ip++
		break
	case *I16Instruction:
		code = g.genI16Instruction(node.(*I16Instruction))
		ok = true
		g.ip++
		break
	case *BraInstruction:
		code = g.genBraInstruction(node.(*BraInstruction))
		ok = true
		g.ip++
		break
	default:
		break
	}
	g.lineNo++
	return code, ok
}

func (g *CodeGen) genRegInstruction(ins *RegInstruction) uint32 {
	code, ok := dataInstructions[ins.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd)
	}
	if ins.set {
		code |= place(1, 1, 25)
	}
	rd, ok := registers[ins.rd]
	if !ok {
		g.Error("unrecognized destination register [%s]", ins.rd)
	}
	code |= place(int64(rd), 20, 4)
	ra, ok := registers[ins.ra]
	if !ok {
		g.Error("unrecognized source A register [%s]", ins.ra)
	}
	code |= place(int64(ra), 16, 4)
	rb, ok := registers[ins.rb]
	if !ok {
		g.Error("unrecognized source B register [%s]", ins.rb)
	}
	code |= place(int64(rb), 12, 4)
	return code
}

func (g *CodeGen) genI12Instruction(ins *I12Instruction) uint32 {
	code, ok := dataInstructions[ins.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd)
	}
	if ins.set {
		code |= place(1, 1, 25)
	}
	rd, ok := registers[ins.rd]
	if !ok {
		g.Error("unrecognized destination register [%s]", ins.rd)
	}
	code |= place(int64(rd), 20, 4)
	ra, ok := registers[ins.ra]
	if !ok {
		g.Error("unrecognized source A register [%s]", ins.ra)
	}
	code |= place(int64(ra), 16, 4)
	code |= g.convertSignedNum(ins.num.value, 4, 12)
	return code
}

func (g *CodeGen) genI16Instruction(ins *I16Instruction) uint32 {
	code, ok := immInstructions[ins.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd)
	}
	if ins.set {
		code |= place(1, 1, 25)
	}
	rd, ok := registers[ins.rd]
	if !ok {
		g.Error("unrecognized destination register [%s]", ins.rd)
	}
	code |= place(int64(rd), 20, 4)
	code |= g.convertSignedNum(ins.num.value, 4, 16)
	return code
}

func (g *CodeGen) genBraInstruction(ins *BraInstruction) uint32 {
	code, ok := branchInstructions[ins.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd)
	}
	sym, ok := g.st.Find(ins.lbl.name)
	if !ok {
		g.Error("Reference to undefined symbol [%s].", ins.lbl.name)
	}
	code |= g.convertAddr(sym.addr)
	return code
}
