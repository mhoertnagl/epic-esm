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

type Label struct {
  Name string
}

func (n *Label) String() string {
  return n.Name
}

type Empty struct {
  
}

func (n *Empty) String() string {
  return "EMPTY"
}

// TODO: []token.Token or []*token.Token?
type Instr struct {
	Set  bool
  Cmd  string
  Cond string
  Args []token.Token
}

type Instrs []*Instr

func (n *Instr) ArgsString() string {
  params := []string{}

  if n.Set {
    params = append(params, "!")
  } else {
    params = append(params, "_")
  }

  params = append(params, n.Cmd)

  params = append(params, "c")

  for _, tok := range n.Args {
    params = append(params, string(tok.Typ))   
  }

	return strings.Join(params, " ")
}

func (n *Instr) String() string {
  params := []string{}

  if n.Set {
    params = append(params, "!")
  }

  params = append(params, n.Cmd)

  params = append(params, n.Cond)
    
  for _, tok := range n.Args {
    params = append(params, tok.Literal)
  }

	return strings.Join(params, " ")
}
