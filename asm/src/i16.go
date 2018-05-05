package main

type I16Instruction struct {
	set bool
	cmd string
	cnd string
	up  bool
	rd  string
	num string
}

func NewI16Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	up interface{},
	rd interface{},
	num interface{}) (*I16Instruction, error) {
	return &I16Instruction{
		set != nil,
		asString(cmd, ""),
		asString(cnd, "al"),
		up != nil,
		asString(rd, ""),
		asString(num, "")}, nil
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
