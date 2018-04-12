package main

type Comment struct{}

type RegInstruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	rb  string
}

type I12Instruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	num *Number
}

type I16Instruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	num *Number
}

type BraInstruction struct {
	cmd string
	cnd string
	lbl *Label
}

type Number struct {
	base  int
	value string
}

type Label struct {
	name string
}

func NewNumber(s []byte, base int) (*Number, error) {
	return &Number{base, string(s)}, nil
}

func NewComment() (*Comment, error) {
	return &Comment{}, nil
}

func NewLabel(s []byte) (*Label, error) {
	return &Label{string(s)}, nil
}

func NewRegInstr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	rb interface{}) (*RegInstruction, error) {
	return &RegInstruction{
		asString(set, "") == "!",
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		asString(rb, "")}, nil
}

func NewI12Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	num interface{}) (*I12Instruction, error) {
	return &I12Instruction{
		asString(set, "") == "!",
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		num.(*Number)}, nil
}

func NewI16Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	up interface{},
	rd interface{},
	num interface{}) (*I16Instruction, error) {
	return &I16Instruction{
		asString(set, "") == "!",
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		num.(*Number)}, nil
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
