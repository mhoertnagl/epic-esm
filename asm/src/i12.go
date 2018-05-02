package main

type I12Instruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	num string
}

func NewI12Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	num interface{}) (*I12Instruction, error) {
	return &I12Instruction{
		set != nil,
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		asString(num, "")}, nil
}

func (ins *I12Instruction) Generate(g *CodeGen) []uint32 {
	code := g.placeDataCmd(ins.cmd)
	code |= g.placeSetBit(ins.set)
	code |= g.placeI12Bit()
	code |= g.placeRd(ins.rd)
	code |= g.placeRa(ins.ra)
	// hängt von der operation ab ob signed oder unsigned
	code |= g.convertSignedNum(ins.num, 4, 12)
	return []uint32{code}
}
