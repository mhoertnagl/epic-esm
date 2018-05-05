package ast

type MemRegInstruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	rb  string
}

func NewMemRegInstr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	rb interface{}) (*MemRegInstruction, error) {
	return &MemRegInstruction{
		set != nil,
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		asString(rb, "")}, nil
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
