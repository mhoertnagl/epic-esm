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

type Instr struct {
	Set  bool
  Cmd  string
  Cond string
  Args []token.Token
}

func (n *Instr) ArgsString() string {
  params := []string{}

  if n.Set {
    params = append(params, "!")
  } else {
    params = append(params, "_")
  }

  params = append(params, n.Cmd)

  // TODO: TokenType kann den Kürzeln entsprechen.
  for _, tok := range n.Args {
    switch tok.Typ {
    case token.REG:
      params = append(params, "r")
      break
    case token.NUM:
      params = append(params, "n")
      break
    case token.LBL:
      params = append(params, "@")
      break
    case token.SLL:
      params = append(params, "s")
      break
    case token.SRL:
      params = append(params, "s")
      break
    case token.SRA:
      params = append(params, "s")
      break
    case token.ROL:
      params = append(params, "s")
      break
    case token.ROR:
      params = append(params, "s")
      break
    case token.LBRK:
      params = append(params, "[")
      break
    case token.RBRK:
      params = append(params, "]")
      break
    }    
  }

	return strings.Join(params, " ")
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
