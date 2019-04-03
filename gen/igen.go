package gen

import(
  //"fmt"
  //"github.com/mhoertnagl/epic-esm/token"
  "github.com/mhoertnagl/epic-esm/ast"
)

type InsGen struct {
  code     []uint32
  ctx      AsmContext
}

func NewInsGen(ctx AsmContext) *InsGen {
  return &InsGen{[]uint32{}, ctx}
}

func (c *InsGen) Generate(ins *ast.Instr) {
  
  if IsDataInstruction(ins.Cmd) {
    c.generateData(ins)
  }
  switch ins.Cmd {
  case "add": 
    c.generateData(ins)
    break
  case "stw":
    c.generateMem(ins)
  }
}

func (c *InsGen) generateData(ins *ast.Instr) {
  g := NewCodeGen(c.ctx)
  g.PlaceDataCmd(ins.Cmd)
  g.PlaceCnd(ins.Cond)
  g.PlaceSetBit(ins.Set)
  g.PlaceRd(ins.Args[0].Literal)
  // add rd ra rb sop num
  // add rd ra rb
  // add rd ra num
  // add rd ra
  // add rd num 
  //
  g.Emit()
  //c.Emit(g)
}

func (c *InsGen) generateMem(ins *ast.Instr) {
  // stw rd ra [ rb ] sop num
  // stw rd ra [ rb ]
  // stw rd ra [ num ]
}

func (c *InsGen) generateBra(ins *ast.Instr) {
  // bra rd @lbl
}

// type Instruction interface {
// 	Generate(g *CodeGen) []uint32
// }
// 
// func (ins *RegInstruction) Generate(g *CodeGen) []uint32 {
// 	code := g.placeDataCmd(ins.cmd)
// 	code |= g.placeCnd(ins.cnd)
// 	code |= g.placeSetBit(ins.set)
// 	code |= g.placeRd(ins.rd)
// 	code |= g.placeRa(ins.ra)
// 	code |= g.placeRb(ins.rb)
// 	code |= g.placeNumShift(ins.sh)
// 	return []uint32{code}
// }
// 
// func (ins *I12Instruction) Generate(g *CodeGen) []uint32 {
// 	code := g.placeDataCmd(ins.cmd)
// 	code |= g.placeCnd(ins.cnd)
// 	code |= g.placeSetBit(ins.set)
// 	code |= g.placeI12Bit()
// 	code |= g.placeRd(ins.rd)
// 	code |= g.placeRa(ins.ra)
// 	// hängt von der operation ab ob signed oder unsigned
// 	code |= g.convertSignedNum(ins.num, 4, 12)
// 	return []uint32{code}
// }
// 
// func (ins *I16Instruction) Generate(g *CodeGen) []uint32 {
// 	code := g.placeDataCmd(ins.cmd)
// 	code |= g.placeCnd(ins.cnd)
// 	code |= g.placeI16Bit()
// 	code |= g.placeSetBit(ins.set)
// 	code |= g.placeRd(ins.rd)
// 	// hängt von der operation ab ob signed oder unsigned
// 	code |= g.convertSignedNum(ins.num, 4, 16)
// 	return []uint32{code}
// }
// 
// func (ins *MemRegInstruction) Generate(g *CodeGen) []uint32 {
// 	code := g.placeMemCmd(ins.cmd)
// 	code |= g.placeCnd(ins.cnd)
// 	code |= g.placeSetBit(ins.set)
// 	code |= g.placeRd(ins.rd)
// 	code |= g.placeRa(ins.ra)
// 	code |= g.placeRb(ins.rb)
// 	return []uint32{code}
// }
// 
// func (ins *MemI12Instruction) Generate(g *CodeGen) []uint32 {
// 	code := g.placeMemCmd(ins.cmd)
// 	code |= g.placeCnd(ins.cnd)
// 	code |= g.placeSetBit(ins.set)
// 	code |= g.placeRd(ins.rd)
// 	code |= g.placeRa(ins.ra)
// 	code |= g.convertSignedNum(ins.num, 4, 12)
// 	return []uint32{code}
// }
// 
// func (ins *BraInstruction) Generate(g *CodeGen) []uint32 {
// 	code := g.placeBranchCmd(ins.cmd)
// 	code |= g.placeCnd(ins.cnd)
// 	code |= g.placeBranchAddress(ins.lbl)
// 	return []uint32{code}
// }
