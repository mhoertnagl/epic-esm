package main

import "fmt"

var dataInstructions = map[string]uint32{
	"add": 0x00000000,
	"sll": 0x00000000,
	"tst": 0x00000000,
	"stw": 0x00000000,
}

var branchInstructions = map[string]uint32{
	"bra": 0x00000000,
	"brl": 0x00000000,
}

var registers = map[string]uint32{
	"r0":  0,
	"r1":  0,
	"r2":  0,
	"r3":  0,
	"r4":  0,
	"r5":  0,
	"r6":  0,
	"r7":  0,
	"r8":  0,
	"r9":  0,
	"r10": 0,
	"r11": 0,
	"r12": 0,
	"r13": 0,
	"r14": 0,
	"r15": 0,
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
	case *ImmInstruction:
		code = g.genImmInstruction(node.(*ImmInstruction))
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
	code, ok := dataInstructions[ins.cmd.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd.cmd)
	}
	rd, ok := registers[ins.rd.name]
	if !ok {
		g.Error("unrecognized destination register [%s]", ins.rd.name)
	}
	code |= place(int64(rd), 20, 4)
	ra, ok := registers[ins.ra.name]
	if !ok {
		g.Error("unrecognized source A register [%s]", ins.ra.name)
	}
	code |= place(int64(ra), 16, 4)
	rb, ok := registers[ins.rb.name]
	if !ok {
		g.Error("unrecognized source B register [%s]", ins.rb.name)
	}
	code |= place(int64(rb), 12, 4)
	return code
}

func (g *CodeGen) genImmInstruction(ins *ImmInstruction) uint32 {
	code, ok := dataInstructions[ins.cmd.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd.cmd)
	}
	rd, ok := registers[ins.rd.name]
	if !ok {
		g.Error("unrecognized destination register [%s]", ins.rd.name)
	}
	code |= place(int64(rd), 20, 4)
	ra, ok := registers[ins.ra.name]
	if !ok {
		g.Error("unrecognized source A register [%s]", ins.ra.name)
	}
	code |= place(int64(ra), 16, 4)
	code |= g.convertSignedNum(ins.num.value, 4, 12)
	return code
}

func (g *CodeGen) genBraInstruction(ins *BraInstruction) uint32 {
	code, ok := branchInstructions[ins.cmd.cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", ins.cmd.cmd)
	}
	sym, ok := g.st.Find(ins.lbl.name)
	if !ok {
		g.Error("Reference to undefined symbol [%s].", ins.lbl.name)
	}
	code |= g.convertAddr(sym.addr)
	return code
}
