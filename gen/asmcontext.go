package gen

import(
  "fmt"
  
  "github.com/mhoertnagl/epic-esm/ast"
)

type AsmContext interface {
  
  Ip() uint32
  
  FindSymbol(name string) (Symbol, bool)
  
  Error(format string, a ...interface{})
  
  Errors() []string
  
  Generate(ins *ast.Instr)
  
  NewCodeGen() *CodeGen
  
  Emit(g *CodeGen)
}

type asmContext struct {
  filename string
  st       SymbolTable
  ip       uint32
  lineNo   uint32
  errors   []string
}

func NewAsmContext(filename string, st SymbolTable) AsmContext {
  return &asmContext{
    filename: filename,
    st: st,
    ip: 0,
    lineNo: 1,
    errors: []string{},
  }
}

func (c *asmContext) Ip() uint32 {
  return c.ip
}

func (c *asmContext) FindSymbol(name string) (Symbol, bool) {
  v, ok := c.st.Find(name)
  return v, ok
}

func (c *asmContext) Error(format string, a ...interface{}) {
  c.errors = append(c.errors, fmt.Sprintf(format, a...))
}

func (c *asmContext) Errors() []string {
	return c.errors
}



func (c *asmContext) NewCodeGen() *CodeGen {
  return &CodeGen{0, c}
}

func (c *asmContext) Emit(g *CodeGen) {
  c.ip++
  c.lineNo++
}
