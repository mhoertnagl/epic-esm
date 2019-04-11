package asm

import(
  "bufio"
  "fmt"
  "os"
  
  "github.com/mhoertnagl/epic-esm/lexer"
  "github.com/mhoertnagl/epic-esm/parser"
  "github.com/mhoertnagl/epic-esm/ast"
  "github.com/mhoertnagl/epic-esm/gen"
)

type AsmConfig struct {
  SrcFilePath string
  BinFilePath string
  LstFilePath string
}

func Assemble(cfg *AsmConfig) {
  
  ctx := gen.NewAsmContext(cfg.SrcFilePath)
  
  srcFile, err1 := os.Open(cfg.SrcFilePath)
  binFile, err2 := os.Open(cfg.BinFilePath)
  lstFile, err3 := os.Open(cfg.LstFilePath)
  
  binWriter := bufio.NewWriter(binFile)
  lstWriter := bufio.NewWriter(lstFile)
  
  if err1 != nil {
    ctx.Error("Could not load source file. [%s].", err1)
  }
  
  if err2 != nil {
    ctx.Error("Could not open output file. [%s].", err2)
  }
  
  if err3 != nil {
    ctx.Error("Could not open listing file. [%s].", err3)
  }
  
  scan(ctx, srcFile)
  // Rewind file pointer and reset IP.
  srcFile.Seek(0, 0)
  ctx.ResetIp()
  compile(ctx, srcFile, binWriter, lstWriter)
  
  if ctx.HasErrors() {
    for _, err := range ctx.Errors() {
      fmt.Println(err)
    }
  }

  binWriter.Flush()
  lstWriter.Flush()
  
  srcFile.Close()
  binFile.Close()
  lstFile.Close()
}

func scan(ctx gen.AsmContext, srcFile *os.File) {
  
  igen := gen.NewInstrGen(ctx)
  scanner := bufio.NewScanner(srcFile)

  for scanner.Scan() {
    
    if err := scanner.Err(); err != nil {
      ctx.Error("Could not scan file. [%s].", err)
    }
    
    line := scanner.Text()
    lexer := lexer.NewLexer(line)
    parser := parser.NewParser(lexer)  
    node := parser.Parse()
    
    switch n := node.(type) {
    case *ast.Label: 
      ctx.AddSymbol(n.Name)
      break
    case *ast.Instr:
      instrs := igen.Generate(n)
      ctx.IncrementIp(len(instrs))
      break
    }
    
    ctx.IncrementLineNo()
  }
}

func compile(ctx gen.AsmContext, srcFile *os.File, binWriter *bufio.Writer, lstWriter *bufio.Writer) {
  
  igen := gen.NewInstrGen(ctx)
  cgen := gen.NewCodeGen(ctx)
  scanner := bufio.NewScanner(srcFile)

  for scanner.Scan() {
    
    if err := scanner.Err(); err != nil {
      ctx.Error("Could not scan file. [%s].", err)
    }
    
    line := scanner.Text()
    lexer := lexer.NewLexer(line)
    parser := parser.NewParser(lexer)  
    node := parser.Parse()
    
    switch n := node.(type) {
    case *ast.Instr:
      instrs := igen.Generate(n)
      for i, ins := range instrs {
        code := cgen.Generate(ins)
        ip := ctx.Ip() + uint32(i)
        fmt.Fprintf(lstWriter, "0x%08x  0x%08x  %s\n", ip, code, line)
        writeInt32BigEndian(binWriter, code)
        ctx.IncrementIp(1)
      }
      break
    default:
      fmt.Fprintf(lstWriter, "%24s%s\n", "", line)
      break
    }
    
    ctx.IncrementLineNo()
  }
}

func writeInt32BigEndian(w *bufio.Writer, i uint32) {
	w.WriteByte(byte(i >> 24))
	w.WriteByte(byte(i >> 16))
	w.WriteByte(byte(i >> 8))
	w.WriteByte(byte(i))
}
