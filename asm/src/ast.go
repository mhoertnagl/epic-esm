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
	num string
}

type I16Instruction struct {
	set bool
	cmd string
	cnd string
	up  bool
	rd  string
	num string
}

type MemRegInstruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	rb  string
}

type MemI12Instruction struct {
	set bool
	cmd string
	cnd string
	rd  string
	ra  string
	num string
}

type BraInstruction struct {
	cmd string
	cnd string
	lbl *Label
}

type Label struct {
	name string
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
		set != nil,
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
		set != nil,
		asString(cmd, ""),
		asString(cnd, "al"),
		asString(rd, ""),
		asString(ra, ""),
		asString(num, "")}, nil
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

func NewBraInstr(
	cmd interface{},
	cnd interface{},
	lbl interface{}) (*BraInstruction, error) {
	return &BraInstruction{
		asString(cmd, ""),
		asString(cnd, "al"),
		lbl.(*Label)}, nil
}

func NewNopInstr(
	cnd interface{}) (*RegInstruction, error) {
	return NewRegInstr(nil, "mov", cnd, "r0", "r0", "r0")
}

func NewClrInstr(
	set interface{},
	cnd interface{},
	rd interface{}) (*RegInstruction, error) {
	return NewRegInstr(set, "xor", cnd, rd, rd, rd)
}
