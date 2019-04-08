package token

type TokenType string

type Token struct {
	Typ     TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	ID      TokenType = "ID"
  // CMD     TokenType = "CMD"
  // COND    TokenType = "c"
  REG     TokenType = "r"
	NUM     TokenType = "n"
  LBL     TokenType = "@"
	// ASSIGN  = "="
	// PLUS    = "+"
	// MINUS   = "-"
	// TIMES   = "*"
	// DIV     = "/"
	// INV     = "~"
	// AND     = "&"
	// OR      = "|"
	// XOR     = "^"
  SOP     TokenType = "s"
	// SLL     TokenType = "<<"
	// SRL     TokenType = ">>"
	// SRA     TokenType = ">>>"
	// ROL     TokenType = "<<>"
	// ROR     TokenType = "<>>"
	SET     TokenType = "!"
	// CONJ    = "&&"
	// DISJ    = "||"
	// EQU     = "=="
	// NEQ     = "!="
	// LT      = "<"
	// LE      = "<="
	// GT      = ">"
	// GE      = ">="
	// COMMA   = ","
	// SCOLON  = ";"
	// LPAR    = "("
	// RPAR    = ")"
	// LBRA    = "{"
	// RBRA    = "}"
  LBRK    TokenType = "["
  RBRK    TokenType = "]"
	// FUN     = "FUN"
	// LET     = "LET"
	// TRUE    = "TRUE"
	// FALSE   = "FALSE"
	// IF      = "IF"
	// ELSE    = "ELSE"
	// RETURN  = "RETURN"
)

var keywords = map[string]TokenType{
  // "nv": COND,
  // "eq": COND,
  // "lt": COND,
  // "le": COND,
  // "gt": COND,
  // "ge": COND,
  // "ne": COND,
  // "al": COND,
  
  "r0":  REG,
  "r1":  REG,
  "r2":  REG,
  "r3":  REG,
  "r4":  REG,
  "r5":  REG,
  "r6":  REG,
  "r7":  REG,
  "r8":  REG,
  "r9":  REG,
  "r10": REG,
  "r11": REG,
  "r12": REG,
  "r13": REG,
  "r14": REG,
  "r15": REG,
  "sp":  REG,
  "rp":  REG,
  "ip":  REG,
  
  
	// "fun":    FUN,
	// "let":    LET,
	// "true":   TRUE,
	// "false":  FALSE,
	// "if":     IF,
	// "else":   ELSE,
	// "return": RETURN,
}

func LookupId(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return ID
}
