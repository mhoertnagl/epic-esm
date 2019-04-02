package gen

import(
  "fmt"
  
  "github.com/mhoertnagl/epic-esm/ast"
)

type AsmContext interface {
  
  Ip() uint32
  
  IncrementIp()
  
  IncrementLineNo()
  
  AddSymbol(n ast.Node)
  
  FindSymbol(name string) (Symbol, bool)
  
  Error(format string, a ...interface{})
  
  Errors() []string
  // 
  // Generate(ins *ast.Instr)
  // 
  // NewCodeGen() *CodeGen
  // 
  // Emit(g *CodeGen)
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

func (c *asmContext) IncrementIp() {
  c.ip++
}

func (c *asmContext) IncrementLineNo() {
  c.lineNo++
}

func (c *asmContext) AddSymbol(n ast.Node) {
  switch l := n.(type) {
  case *ast.Label: 
    c.st.Add(l.Name, c.ip, c.lineNo)
    break
  }
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
// 
// 
// 
// func (c *asmContext) NewCodeGen() *CodeGen {
//   return &CodeGen{0, c}
// }
// 
// func (c *asmContext) Emit(g *CodeGen) {
//   c.ip++
//   c.lineNo++
// }

// func (c *asmContext) AddSymbol(label *ast.Label) {
// 	c.st.Add(label.Name, c.ip, c.lineNo)
// }
