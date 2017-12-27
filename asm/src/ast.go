package main

import "strconv"

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
	value int32
}

type Label struct {
	name string
}

type Identifer struct {
	value string
}

func NewNumber(s []byte, base int) (*Number, error) {
	n, err := strconv.ParseInt(string(s), base, 32)
	return &Number{int32(n)}, err
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

func NewImmInstruction(cmd *Command, rd *Register, ra *Register, num *Number) (*ImmInstruction, error) {
	return &ImmInstruction{cmd, rd, ra, num}, nil
}

func NewBraInstruction(cmd *Command, lbl *Label) (*BraInstruction, error) {
	return &BraInstruction{cmd, lbl}, nil
}
