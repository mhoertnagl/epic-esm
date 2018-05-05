package gen

type Instruction interface {
	Generate(g *CodeGen) []uint32
}

func (ins *RegInstruction) Generate(g *CodeGen) []uint32 {
	code := g.placeDataCmd(ins.cmd)
	code |= g.placeCnd(ins.cnd)
	code |= g.placeSetBit(ins.set)
	code |= g.placeRd(ins.rd)
	code |= g.placeRa(ins.ra)
	code |= g.placeRb(ins.rb)
	code |= g.placeNumShift(ins.sh)
	return []uint32{code}
}

func (ins *I12Instruction) Generate(g *CodeGen) []uint32 {
	code := g.placeDataCmd(ins.cmd)
	code |= g.placeCnd(ins.cnd)
	code |= g.placeSetBit(ins.set)
	code |= g.placeI12Bit()
	code |= g.placeRd(ins.rd)
	code |= g.placeRa(ins.ra)
	// hängt von der operation ab ob signed oder unsigned
	code |= g.convertSignedNum(ins.num, 4, 12)
	return []uint32{code}
}

func (ins *I16Instruction) Generate(g *CodeGen) []uint32 {
	code := g.placeDataCmd(ins.cmd)
	code |= g.placeCnd(ins.cnd)
	code |= g.placeI16Bit()
	code |= g.placeSetBit(ins.set)
	code |= g.placeRd(ins.rd)
	// hängt von der operation ab ob signed oder unsigned
	code |= g.convertSignedNum(ins.num, 4, 16)
	return []uint32{code}
}

func (ins *MemRegInstruction) Generate(g *CodeGen) []uint32 {
	code := g.placeMemCmd(ins.cmd)
	code |= g.placeCnd(ins.cnd)
	code |= g.placeSetBit(ins.set)
	code |= g.placeRd(ins.rd)
	code |= g.placeRa(ins.ra)
	code |= g.placeRb(ins.rb)
	return []uint32{code}
}

func (ins *MemI12Instruction) Generate(g *CodeGen) []uint32 {
	code := g.placeMemCmd(ins.cmd)
	code |= g.placeCnd(ins.cnd)
	code |= g.placeSetBit(ins.set)
	code |= g.placeRd(ins.rd)
	code |= g.placeRa(ins.ra)
	code |= g.convertSignedNum(ins.num, 4, 12)
	return []uint32{code}
}

func (ins *BraInstruction) Generate(g *CodeGen) []uint32 {
	code := g.placeBranchCmd(ins.cmd)
	code |= g.placeCnd(ins.cnd)
	code |= g.placeBranchAddress(ins.lbl)
	return []uint32{code}
}
