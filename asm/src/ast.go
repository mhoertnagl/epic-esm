package main

type Comment struct{}

type Command struct {
	set bool
	cmd string
	cnd string
}

type RegInstruction struct {
	cmd *Command
	rd  *Register
	ra  *Register
	rb  *Register
}

type ImmInstruction struct {
	cmd *Command
	rd  *Register
	ra  *Register
	num *Number
}

type BraInstruction struct {
	cmd *Command
	lbl *Label
}

type Register struct {
	name string
}

type Number struct {
	base  int
	value string
}

type Label struct {
	name string
}

type Identifer struct {
	value string
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

func NewRegister(s []byte) (*Register, error) {
	return &Register{string(s)}, nil
}

func NewCommand(set interface{}, cmd interface{}, cnd interface{}) (*Command, error) {
	return &Command{asString(set, "") == "!", asString(cmd, ""), asString(cnd, "al")}, nil
}

func NewRegInstruction(cmd *Command, rd *Register, ra *Register, rb *Register) (*RegInstruction, error) {
	return &RegInstruction{cmd, rd, ra, rb}, nil
}

func NewRegInstr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	rb interface{}) (*RegInstruction, error) {
	command, _ := NewCommand(set, cmd, cnd)
	return &RegInstruction{
		command,
		rd.(*Register),
		ra.(*Register),
		rb.(*Register)}, nil
}

func NewImm12Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	ra interface{},
	num interface{}) (*ImmInstruction, error) {
	command, _ := NewCommand(set, cmd, cnd)
	return &ImmInstruction{
		command,
		rd.(*Register),
		ra.(*Register),
		num.(*Number)}, nil
}

func NewImm16Instr(
	set interface{},
	cmd interface{},
	cnd interface{},
	rd interface{},
	num interface{}) (*ImmInstruction, error) {
	command, _ := NewCommand(set, cmd, cnd)
	return &ImmInstruction{
		command,
		rd.(*Register),
		rd.(*Register),
		num.(*Number)}, nil
}

func NewImmInstruction(cmd *Command, rd *Register, ra *Register, num *Number) (*ImmInstruction, error) {
	return &ImmInstruction{cmd, rd, ra, num}, nil
}

func NewBraInstruction(cmd *Command, lbl *Label) (*BraInstruction, error) {
	return &BraInstruction{cmd, lbl}, nil
}
