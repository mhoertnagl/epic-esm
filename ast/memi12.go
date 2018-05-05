package ast

type MemI12Instruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	num string
}

func NewMemI12Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	num interface{}) (*MemI12Instruction, error) {
	return &MemI12Instruction{
		set != nil,
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		asString(num, "")}, nil
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
