package gen

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
