package main

type BraInstruction struct {
	cmd string
	cnd string
	lbl *Label
}

func NewBraInstr(
	cmd interface{},
	cnd interface{},
	lbl interface{}) (*BraInstruction, error) {
	return &BraInstruction{
		asString(cmd, ""),
		asString(cnd, "al"),
		lbl.(*Label)}, nil
}

func (ins *BraInstruction) Generate(g *CodeGen) []uint32 {
	code := g.placeBranchCmd(ins.cmd)
	code |= g.placeBranchAddress(ins.lbl)
	return []uint32{code}
}
