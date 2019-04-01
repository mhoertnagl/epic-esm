package ast

import (
  "strings"

	"github.com/mhoertnagl/epic-esm/token"
)

type Node interface {
	String() string
}

type Err struct {
  
}

func (n *Err) String() string {
	return "ERROR"
}

type NumShift struct {
	Cmd string
	Num string
}

type RegInstruction struct {
	Set bool
	Cmd string
	Cnd string
	Rd  string
	Ra  string
	Rb  string
	Sh  *NumShift
}

func (n *RegInstruction) String() string {
  args := []string{}
  if n.Set {
    args = append(args, "!")
  }
  args = append(args, n.Cmd)
  args = append(args, n.Cnd)
  args = append(args, n.Rd)
  args = append(args, n.Ra)
  args = append(args, n.Rb)
  args = append(args, n.Sh.Cmd)
  args = append(args, n.Sh.Num)

	return strings.Join(args, " ")
}

type I12Instruction struct {
	Set bool
	Cmd string
	Cnd string
	Rd  string
	Ra  string
	Num string
}

func (n *I12Instruction) String() string {
  args := []string{}
  if n.Set {
    args = append(args, "!")
  }
  args = append(args, n.Cmd)
  args = append(args, n.Cnd)
  args = append(args, n.Rd)
  args = append(args, n.Ra)
  args = append(args, n.Num)

	return strings.Join(args, " ")
}

type I16Instruction struct {
	Set bool
	Cmd string
	Cnd string
	Up  bool
	Rd  string
	Num string
}

func (n *I16Instruction) String() string {
  args := []string{}
  if n.Set {
    args = append(args, "!")
  }
  args = append(args, n.Cmd)
  args = append(args, n.Cnd)
  args = append(args, n.Rd)
  args = append(args, n.Num)

	return strings.Join(args, " ")
}

type Instr struct {
	Set  bool
  Cmd  string
  Cond string
  Args []token.Token
}

func (c *asmContext) Generate(ins *Instr) []uint32 {
  if gen.IsDataInstruction(ins.Cmd) {
    return GenerateData(ins)
  }
  switch ins.Cmd {
  case "add": return GenerateData(ins)
  }
}

func (c *asmContext) GenerateData(ins *Instr) []uint32 {
  g := NewCoder()
  g.PlaceThis()
  g.PlaceThat()
  // add rd ra rb sop num
  // add rd ra rb
  // add rd ra num
  // add rd ra
  // add rd num 
  //
  c.Emit(g)
}

func GenerateMem(ins *Instr) []uint32 {
  // stw rd ra [ rb ] sop num
  // stw rd ra [ rb ]
  // stw rd ra [ num ]
}

func GenerateBra(ins *Instr) []uint32 {
  // bra rd @lbl
}

// 
// func (n *Instr) String() string {
//   params := []string{}
// 
//   if n.Set {
//     params = append(params, "!")
//   }
// 
//   params = append(params, n.Cmd)
// 
//   if n.Cond != "al" {
//     params = append(params, n.Cond)
//   }
// 
//   for _, tok := range n.Args {
//     params = append(params, tok.Literal)
//   }
// 
// 	return strings.Join(params, " ")
// }
