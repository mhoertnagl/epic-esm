package gengen

import (
  "fmt"
  "strings"
)

type Pattern interface {
  Pattern() string
}

type BitsPattern struct {
  Val uint32
  Len uint8
}

func (p *BitsPattern) Pattern() string {
  return fmt.Sprintf("bits %dd%d", p.Len, p.Val)
}

func NewBitsPattern(val uint32, len uint8) *BitsPattern {
  return &BitsPattern{ Val: val, Len: len }
}

type VarPattern struct {
  Name string
}

func NewVarPattern(name string) *VarPattern {
  return &VarPattern{ Name: name }
}

func (p *VarPattern) Pattern() string {
  return p.Name
}

type LiteralPattern struct {
  Val string
}

func NewLiteralPattern(val string) *LiteralPattern {
  return &LiteralPattern{ Val: val }
}

func (p *LiteralPattern) Pattern() string {
  return fmt.Sprintf("'%s'", p.Val)
}

type OptionalPattern struct {
  Kid Pattern
}

func NewOptionalPattern(kid Pattern) *OptionalPattern {
  return &OptionalPattern{ Kid: kid }
}

func (p *OptionalPattern) Pattern() string {
  return p.Kid.Pattern()
}

type UnionPattern struct {
  Kids []Pattern
}

func (p *UnionPattern) Pattern() string {
  kids := []string{}
  for _, kid := range p.Kids {
    kids = append(kids, kid.Pattern())
  }
  return strings.Join(kids, " | ")
}

func NewUnionPattern(kids ...Pattern) *UnionPattern {
  return &UnionPattern{ Kids: kids }
}

type SubPattern struct {
  Kids []Pattern
}

func (p *SubPattern) Pattern() string {
  kids := []string{}
  for _, kid := range p.Kids {
    kids = append(kids, kid.Pattern())
  }
  return fmt.Sprintf("(%s)", strings.Join(kids, " | "))
}

func NewSubPattern(kids ...Pattern) *SubPattern {
  return &SubPattern{ Kids: kids }
}

var patternDefaults = map[string]Pattern{
  "!": NewBitsPattern(0, 1),
  "cond": NewVarPattern("al"),
  "(<< 16)": NewBitsPattern(0, 1),
}

var patterns = map[string]Pattern{
  "!": NewBitsPattern(1, 1),
  
  "nv": NewBitsPattern(0, 3),
  "eq": NewBitsPattern(1, 3),
  "lt": NewBitsPattern(2, 3),
  "le": NewBitsPattern(3, 3),
  "gt": NewBitsPattern(4, 3),
  "ge": NewBitsPattern(5, 3),
  "ne": NewBitsPattern(6, 3),
  "al": NewBitsPattern(7, 3),
  
  "cond": anyCondPattern,
  
	"r0":  NewBitsPattern(0, 5),
	"r1":  NewBitsPattern(1, 5),
	"r2":  NewBitsPattern(2, 5),
	"r3":  NewBitsPattern(3, 5),
	"r4":  NewBitsPattern(4, 5),
	"r5":  NewBitsPattern(5, 5),
	"r6":  NewBitsPattern(6, 5),
	"r7":  NewBitsPattern(7, 5),
	"r8":  NewBitsPattern(8, 5),
	"r9":  NewBitsPattern(9, 5),
	"r10": NewBitsPattern(10, 5),
	"r11": NewBitsPattern(11, 5),
	"r12": NewBitsPattern(12, 5),
	"r13": NewBitsPattern(13, 5),
	"r14": NewBitsPattern(14, 5),
	"r15": NewBitsPattern(15, 5),
	"sp":  NewBitsPattern(13, 5),
	"rp":  NewBitsPattern(14, 5),
	"ip":  NewBitsPattern(15, 5),
  
  "rd":  anyRegPattern,
  "ra":  anyRegPattern,
  "rb":  anyRegPattern,
  
  "<<": NewBitsPattern(0, 2),
  ">>": NewBitsPattern(1, 2),
  ">>>": NewBitsPattern(2, 2),
  "<<>": NewBitsPattern(3, 2),
  "<>>": NewBitsPattern(3, 2), // NumConverter???
  
  "(<< 16)": NewBitsPattern(1, 1),
}

var anyCondPattern = NewUnionPattern(
    NewVarPattern("nv"),
    NewVarPattern("eq"),
    NewVarPattern("lt"),
    NewVarPattern("le"),
    NewVarPattern("gt"),
    NewVarPattern("ge"),
    NewVarPattern("ne"),
    NewVarPattern("al"),
) 

var anyRegPattern = NewUnionPattern(
    NewVarPattern("r0"),
    NewVarPattern("r1"),
    NewVarPattern("r2"),
    NewVarPattern("r3"),
    NewVarPattern("r4"),
    NewVarPattern("r5"),
    NewVarPattern("r6"),
    NewVarPattern("r7"),
    NewVarPattern("r8"),
    NewVarPattern("r9"),
    NewVarPattern("r10"),
    NewVarPattern("r11"),
    NewVarPattern("r12"),
    NewVarPattern("r13"),
    NewVarPattern("r14"),
    NewVarPattern("r15"),
    NewVarPattern("sp"),
    NewVarPattern("rp"),
    NewVarPattern("ip"),
) 

type AsmContext interface {
  
  Ip() uint32
  
  FindSymbol(name string) (*Symbol, bool)
  
  //GenerateIns(seq *InstrSeq) []uint32
  
  GeneratePat(seq *PatternSeq, env Env) uint32 
  
  Error(format string, a ...interface{})
}

type asmContext struct {
  
}

func NewAsmContext() AsmContext {
  return &asmContext{}
}

type SymConstraint func (c AsmContext, val string) bool

func unionConstraint (union *UnionPattern) SymConstraint {
  return func (c AsmContext, val string) bool {
    for _, pattern := range union.Kids {
      switch p := pattern.(type) {
      case *VarPattern:
        if val == p.Name {
          return true
        }
      }
    }
    c.Error("Expecting one of [%s] but was [%s]", union.Kids, val)
    return false
  }
}

var patConstraints = map[string]SymConstraint {
  "rd": unionConstraint(anyRegPattern),
  "ra": unionConstraint(anyRegPattern),
  "rb": unionConstraint(anyRegPattern),
}

type BitConstraint func (c AsmContext, val int32) bool

func rangeConstraint (min int32, max int32) BitConstraint {
  return func (c AsmContext, val int32) bool {
    if val < min {
      c.Error("")
      return false
    }
    if val > max {
      c.Error("")
      return false      
    }
    return true
  }
}

var bitConstraints = map[string]BitConstraint {
  "u5": rangeConstraint(0, 31),
  "s12": rangeConstraint(-4095, 4096),
  "u12": rangeConstraint(0, 8192),
}

type SymConverter func (c AsmContext, val string) uint32

func pcRelativeConverter () SymConverter {
  return func (c AsmContext, lbl string) uint32 {
    sym, ok := c.FindSymbol(lbl)
    if !ok {
      c.Error("")
      return 0
    }
    return uint32(sym.addr - c.Ip());
  }
}

var converters = map[string]SymConverter {
  "@lbl": pcRelativeConverter(),
}

type Seq interface {
  Seq() string
}

type InstrSeq struct {
  Literals []*LiteralPattern
}

func (s *InstrSeq) Seq() string {
  kids := []string{}
  for _, kid := range s.Literals {
    kids = append(kids, kid.Pattern())
  }
  return fmt.Sprintf("ins %s", strings.Join(kids, " "))
}

type PatternSeq struct {
  Patterns []Pattern
}

func (s *PatternSeq) Seq() string {
  kids := []string{}
  for _, kid := range s.Patterns {
    kids = append(kids, kid.Pattern())
  }
  return fmt.Sprintf("pat %s", strings.Join(kids, " "))
}

type Transformation struct {
  Key InstrSeq
  Val []Seq
}

var transformations = map[string][]Transformation{
  // "add": &[]Transformation{
  // 
  // }
}

// "(<< 16)": NewSubPattern(
//   NewLiteralPattern("<<"),
//   NewLiteralPattern("16"),
// ),

func (c *asmContext) place(i uint32, s uint8, p uint8) uint32 {
	return uint32((i & ((1 << p) - 1)) << s)
}

// Finden einer Instruktion in den Transformations.
// Constraints überprüfen.
// Conversions anwenden.
// Erstellen einer env Map aus der gematchten Instruktion.

type Env map[string]*BitsPattern

func (c *asmContext) GeneratePat(seq *PatternSeq, env Env) uint32 {
  var ins uint32
  var idx uint8 = 32
  for _, pattern := range seq.Patterns {
    switch p := pattern.(type) {
    case *BitsPattern:
      idx -= p.Len
      ins |= c.place(p.Val, idx, p.Len)
    case *VarPattern:
      bp, ok := env[p.Name]
      if !ok {
        // Report unzugewiesener parameter.
      }
      idx -= bp.Len
      ins |= c.place(bp.Val, idx, bp.Len)      
    default:
      // Report error
    }
  }
  return ins
}

func (c *asmContext) Ip() uint32 {
  return 0
}

func (c *asmContext) FindSymbol(name string) (*Symbol, bool) {
  return nil, false
}

func (c *asmContext) Error(format string, a ...interface{}) {
  
}

// for _, pattern := range seq.Patterns {
//   switch p := pattern.(type) {
//   case *BitsPattern:
//     idx, ins = c.GenerateBitsPat(idx, ins, p)
//   case *VarPattern:
//     bp := c.Resolve(p)
//     idx, ins = c.GenerateBitsPat(idx, ins, bp)
//   }
// }

// func (c *asmContext) GenerateBitsPat(idx uint8,ins uint32, p *BitsPattern) (uint8, uint32) {
//   return idx - p.Len, ins | c.place(p.Val, p.Len, idx - p.Len)
// }

func (c *asmContext) Resolve(p *VarPattern) *BitsPattern {
  return nil
}
