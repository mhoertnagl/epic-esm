package gen

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

type I16Instruction struct {
	set bool
	cmd string
	cnd string
	up  bool
	rd  string
	num string
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
