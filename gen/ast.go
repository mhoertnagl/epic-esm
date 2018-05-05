package gen

type Comment struct{}

func NewComment() (*Comment, error) {
	return &Comment{}, nil
}

type Label struct {
	Name string
}

func NewLabel(s []byte) (*Label, error) {
	return &Label{string(s)}, nil
}

type NumShift struct {
	cmd string
	num string
}

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
