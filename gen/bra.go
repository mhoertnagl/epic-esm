package gen

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
