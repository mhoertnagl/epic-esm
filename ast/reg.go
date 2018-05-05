package ast

//import "github.com/mhoertnagl/epic-esm/gen"

type RegInstruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	rb  string
	sh  *NumShift
}

func NewRegInstr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	rb interface{},
	sh interface{}) (*RegInstruction, error) {
	ins := &RegInstruction{
		set != nil,
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		asString(rb, ""),
		nil}
	if sh != nil {
		ins.sh = sh.(*NumShift)
	}
	return ins, nil
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
