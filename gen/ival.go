package gen

import(
  //"fmt"
  "strings"
  //"github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/ast"
)

type ArgType string

type ArgTypeMapping map[string]string

const (
  Error ArgType = "ERROR"
  Reg3Shift ArgType = "r r r s n"
  Reg3 ArgType = "r r r"
  Reg2Shift ArgType = "r r s n"
  Reg2I12 ArgType = "r r n"
  Reg2 ArgType = "r r"
  Reg1I16 ArgType = "r n"
)

type IVal struct {
  ctx AsmContext
  mapping ArgTypeMapping
}

func NewIVal(ctx AsmContext) *IVal {
  v := &IVal{ctx, ArgTypeMapping{}}
  
  v.Add("add r r r s n")
  v.Add("add r r r")
  v.Add("add r r s n")
  v.Add("add r r n")
  v.Add("add r r")
  v.Add("add r n")
  
  
  
  return v
}

func (v *IVal) Add(ins string) {
  parts := strings.SplitN(ins, " ", 2)
  v.mapping[ins] = parts[1]
}

func (v *IVal) Validate(ins *ast.Instr) string {
  as := ins.ArgsString()
  a, ok := v.mapping[as]
  if !ok {
    v.ctx.Error("Unsupported arguments [%s] for command [%s].", as, ins.Cmd)
    return Error
  }
  return a  
}
