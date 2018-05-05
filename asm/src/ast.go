package main

type Instruction interface {
	Generate(g *CodeGen) []uint32
}

type Comment struct{}

type Label struct {
	name string
}

func NewComment() (*Comment, error) {
	return &Comment{}, nil
}

func NewLabel(s []byte) (*Label, error) {
	return &Label{string(s)}, nil
}

// type RegInstruction struct {
// 	set bool
// 	cmd string
// 	cnd string
// 	rd  string
// 	ra  string
// 	rb  string
// 	sh  *NumShift
// }

type NumShift struct {
	cmd string
	num string
}

// type I12Instruction struct {
// 	set bool
// 	cmd string
// 	cnd string
// 	rd  string
// 	ra  string
// 	num string
// }

// type I16Instruction struct {
// 	set bool
// 	cmd string
// 	cnd string
// 	up  bool
// 	rd  string
// 	num string
// }

// type MemRegInstruction struct {
// 	set bool
// 	cmd string
// 	cnd string
// 	rd  string
// 	ra  string
// 	rb  string
// }

// type MemI12Instruction struct {
// 	set bool
// 	cmd string
// 	cnd string
// 	rd  string
// 	ra  string
// 	num string
// }

// type BraInstruction struct {
// 	cmd string
// 	cnd string
// 	lbl *Label
// }

type LdcInstruction struct {
	set bool
	cnd string
	rd  string
	num string
}

type LdaInstruction struct {
	set bool
	cnd string
	rd  string
	lbl *Label
}

// func NewRegInstr(
// 	set interface{},
// 	cmd interface{},
// 	cnd interface{},
// 	rd interface{},
// 	ra interface{},
// 	rb interface{},
// 	sh interface{}) (*RegInstruction, error) {
// 	return &RegInstruction{
// 		set != nil,
// 		asString(cmd, ""),
// 		asString(cnd, "al"),
// 		asString(rd, ""),
// 		asString(ra, ""),
// 		asString(rb, ""),
// 		sh.(*NumShift)}, nil
// }

func NewNumShift(
	cmd interface{},
	num interface{}) (*NumShift, error) {
	return &NumShift{
		asString(cmd, ""),
		asString(num, "")}, nil
}

func NewShiftInstr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	rb interface{},
	num interface{}) (*RegInstruction, error) {
	sh, _ := NewNumShift(cmd, num)
	return &RegInstruction{
		set != nil,
		"mov",
		asString(cnd, "al"),
		asString(rd, ""),
		"r0",
		asString(rb, ""),
		sh}, nil
}

// func NewI12Instr(
// 	set interface{},
// 	cmd interface{},
// 	cnd interface{},
// 	rd interface{},
// 	ra interface{},
// 	num interface{}) (*I12Instruction, error) {
// 	return &I12Instruction{
// 		set != nil,
// 		asString(cmd, ""),
// 		asString(cnd, "al"),
// 		asString(rd, ""),
// 		asString(ra, ""),
// 		asString(num, "")}, nil
// }

// func NewI16Instr(
// 	set interface{},
// 	cmd interface{},
// 	cnd interface{},
// 	up interface{},
// 	rd interface{},
// 	num interface{}) (*I16Instruction, error) {
// 	return &I16Instruction{
// 		set != nil,
// 		asString(cmd, ""),
// 		asString(cnd, "al"),
// 		up != nil,
// 		asString(rd, ""),
// 		asString(num, "")}, nil
// }

// func NewMemRegInstr(
// 	set interface{},
// 	cmd interface{},
// 	cnd interface{},
// 	rd interface{},
// 	ra interface{},
// 	rb interface{}) (*MemRegInstruction, error) {
// 	return &MemRegInstruction{
// 		set != nil,
// 		asString(cmd, ""),
// 		asString(cnd, "al"),
// 		asString(rd, ""),
// 		asString(ra, ""),
// 		asString(rb, "")}, nil
// }

// func NewMemI12Instr(
// 	set interface{},
// 	cmd interface{},
// 	cnd interface{},
// 	rd interface{},
// 	ra interface{},
// 	num interface{}) (*MemI12Instruction, error) {
// 	return &MemI12Instruction{
// 		set != nil,
// 		asString(cmd, ""),
// 		asString(cnd, "al"),
// 		asString(rd, ""),
// 		asString(ra, ""),
// 		asString(num, "")}, nil
// }

// func NewBraInstr(
// 	cmd interface{},
// 	cnd interface{},
// 	lbl interface{}) (*BraInstruction, error) {
// 	return &BraInstruction{
// 		asString(cmd, ""),
// 		asString(cnd, "al"),
// 		lbl.(*Label)}, nil
// }

func NewNopInstr(
	cnd interface{}) (*RegInstruction, error) {
	return NewRegInstr(nil, "mov", cnd, "r0", "r0", "r0", nil)
}

func NewClrInstr(
	set interface{},
	cnd interface{},
	rd interface{}) (*RegInstruction, error) {
	return NewRegInstr(set, "xor", cnd, rd, rd, rd, nil)
}

func NewLdcInstr(
	set interface{},
	cnd interface{},
	rd interface{},
	num interface{}) (*LdcInstruction, error) {
	return &LdcInstruction{
		set != nil,
		asString(cnd, "al"),
		asString(rd, ""),
		asString(num, "")}, nil
}

func NewLdaInstr(
	set interface{},
	cnd interface{},
	rd interface{},
	lbl interface{}) (*LdaInstruction, error) {
	return &LdaInstruction{
		set != nil,
		asString(cnd, "al"),
		asString(rd, ""),
		lbl.(*Label)}, nil
}
