package parser

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

type Instr struct {
	Set  bool
  Cmd  string
  Cond string
  Args []token.Token
}

func (n *Instr) String() string {
  params := []string{}
  
  if n.Set {
    params = append(params, "!")
  }
  
  params = append(params, n.Cmd)
  
  if n.Cond != "al" {
    params = append(params, n.Cond)
  }
  
  for _, tok := range n.Args {
    params = append(params, tok.Literal)
  }
  
	return strings.Join(params, " ")
}
