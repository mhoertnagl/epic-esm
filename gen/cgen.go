package gen

import (
	"fmt"
	"strconv"
)

type CodeGen struct {
	filename string
	st       SymbolTable
	ip       uint32
	lineNo   uint32
}

func NewCodeGen(filename string, st SymbolTable) *CodeGen {
	return &CodeGen{filename, st, 0, 1}
}

func (g *CodeGen) GetIp() uint32 {
	return g.ip
}

func (g *CodeGen) Error(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s [%d] ERROR: %s\n", g.filename, g.lineNo, msg)
}

func (g *CodeGen) Generate(ins Instruction) []uint32 {
	codes := ins.Generate(g)
	g.ip += uint32(len(codes))
	g.lineNo++
	return codes
}

func (g *CodeGen) placeDataCmd(cmd string) uint32 {
	code, ok := dataInstructions[cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", cmd)
	}
	return code
}

func (g *CodeGen) placeMemCmd(cmd string) uint32 {
	code, ok := memInstructions[cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", cmd)
	}
	return code
}

func (g *CodeGen) placeBranchCmd(cmd string) uint32 {
	code, ok := branchInstructions[cmd]
	if !ok {
		g.Error("Unrecognized instruction [%s].", cmd)
	}
	return code
}

func (g *CodeGen) placeCnd(cnd string) uint32 {
	code, ok := conditions[cnd]
	if !ok {
		g.Error("Unrecognized condition flag [%s].", cnd)
	}
	return g.place(int64(code), 26, 3)
}

func (g *CodeGen) placeSetBit(set bool) uint32 {
	if set {
		return g.place(1, 25, 1)
	}
	return 0
}

func (g *CodeGen) placeI16Bit() uint32 {
	return g.place(1, 29, 1)
}

func (g *CodeGen) placeI12Bit() uint32 {
	return g.place(1, 24, 1)
}

func (g *CodeGen) placeRd(rdName string) uint32 {
	rd, ok := registers[rdName]
	if !ok {
		g.Error("unrecognized destination register [%s]", rdName)
	}
	return g.place(int64(rd), 20, 4)
}

func (g *CodeGen) placeRa(raName string) uint32 {
	ra, ok := registers[raName]
	if !ok {
		g.Error("unrecognized source A register [%s]", raName)
	}
	return g.place(int64(ra), 16, 4)
}

func (g *CodeGen) placeRb(rbName string) uint32 {
	rb, ok := registers[rbName]
	if !ok {
		g.Error("unrecognized source B register [%s]", rbName)
	}
	return g.place(int64(rb), 12, 4)
}

func (g *CodeGen) placeNumShift(sh *NumShift) uint32 {
	if sh == nil {
		return 0
	}
	code := g.placeShiftOp(sh.cmd)
	// Turns a Rotate Right (<>>) into a Rotate Left (<<>). The following
	// identity holds for all cases: x <>> n <--> x <<> (32 - n)
	if sh.cmd == "<>>" || sh.cmd == "ror" {
		shft := g.convertUnsignedNum(sh.num, 0, 5)
		code |= g.place(int64(32-shft), 4, 5)
	} else {
		code |= g.convertUnsignedNum(sh.num, 4, 5)
	}
	return code
}

func (g *CodeGen) placeShiftOp(cmd string) uint32 {
	sop, ok := shiftOps[cmd]
	if !ok {
		g.Error("unrecognized shift operator [%s]", cmd)
	}
	return g.place(int64(sop), 2, 9)
}

func (g *CodeGen) placeBranchAddress(lbl *Label) uint32 {
	sym, ok := g.st.Find(lbl.Name)
	if !ok {
		g.Error("Reference to undefined symbol [%s].", lbl.Name)
	}
	return g.convertAddr(sym.addr)
}

func (g *CodeGen) convertSignedNum(n string, s uint8, p uint8) uint32 {
	return g.convertNum(n, s, p, -(1 << p), 1<<p)
}

func (g *CodeGen) convertUnsignedNum(n string, s uint8, p uint8) uint32 {
	return g.convertNum(n, s, p, 0, 1<<p)
}

func (g *CodeGen) convertNum(n string, s uint8, p uint8, min int64, max int64) uint32 {
	i, err := g.parseNum(n)

	if err != nil {
		g.Error("Number [%s] too long.", n)
	}
	if i < min {
		g.Error("Unexpected number [%s]. Number must be greater than [%d].", n, min)
	}
	if i >= max {
		g.Error("Unexpected number [%s]. Number must be less than [%d]", n, max)
	}
	return g.place(i, s, p)
}

func (g *CodeGen) convertAddr(addr uint32) uint32 {
	bra := int64(addr - g.ip)
	if bra < BRA_MIN || bra >= BRA_MAX {
		g.Error("Branch distance [%d] too large.", bra)
	}
	return g.place(bra, 0, 25)
}

func (g *CodeGen) parseNum(n string) (int64, error) {
	// strings.HasPrefix
	if len(n) > 2 && n[0:2] == "0b" {
		return strconv.ParseInt(n[2:], 2, 32)
	}
	if len(n) > 2 && n[0:2] == "0x" {
		return strconv.ParseInt(n[2:], 16, 32)
	}
	return strconv.ParseInt(n, 10, 32)
}

func (g *CodeGen) place(i int64, s uint8, p uint8) uint32 {
	return uint32((i & ((1 << p) - 1)) << s)
}
