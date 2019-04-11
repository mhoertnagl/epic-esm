package gen

import(
  "fmt"
  
  //"github.com/mhoertnagl/epic-esm/ast"
)

type AsmContext interface {
  
  Ip() uint32
  
  ResetIp()
  
  IncrementIp(n int)
  
  IncrementLineNo()
  
  AddSymbol(name string)
  
  FindSymbol(name string) (Symbol, bool)
  
  Error(format string, a ...interface{})
  
  HasErrors() bool
  
  Errors() []string
}

type asmContext struct {
  filename string
  st       SymbolTable
  ip       uint32
  lineNo   uint32
  errors   []string
}

func NewAsmContext(filename string) AsmContext {
  return &asmContext{
    filename: filename,
    st: NewSymbolTable(),
    ip: 0,
    lineNo: 1,
    errors: []string{},
  }
}

func (c *asmContext) Ip() uint32 {
  return c.ip
}

func (c *asmContext) ResetIp() {
  c.ip = 0
}

func (c *asmContext) IncrementIp(n int) {
  c.ip += uint32(n)
}

func (c *asmContext) IncrementLineNo() {
  c.lineNo++
}

// func (c *asmContext) AddSymbol(n ast.Node) {
//   switch n := n.(type) {
//   case *ast.Label: 
//     c.st.Add(n.Name, c.ip, c.lineNo)
//     break
//   }
// }

func (c *asmContext) AddSymbol(name string) {
  c.st.Add(name, c.ip, c.lineNo)
}

func (c *asmContext) FindSymbol(name string) (Symbol, bool) {
  v, ok := c.st.Find(name)
  return v, ok
}

func (c *asmContext) Error(format string, a ...interface{}) {
  c.errors = append(c.errors, fmt.Sprintf(format, a...))
}

func (c *asmContext) HasErrors() bool {
  return len(c.errors) > 0
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
