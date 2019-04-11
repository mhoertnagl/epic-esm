package main

import (
	"flag"
  
  "github.com/mhoertnagl/epic-esm/asm"
)

func main() {
	binFilePathPtr := flag.String("o", "x.bin", "an output file")
	lstFilePathPtr := flag.String("l", "x.lst", "a listing file")

	flag.Parse()

	//args := flag.Args()

	// if len(args) < 1 {
	// 	panic("Too few arguments. Provide a single input file.")
	// }
	//
	// if len(args) > 1 {
	// 	panic("Too many arguments. Provide only a single input file.")
	// }

	srcFilePath := "examples/test.esm" //args[0]

  cfg := &asm.AsmConfig{
    SrcFilePath: srcFilePath,
    BinFilePath: *binFilePathPtr,
    LstFilePath: *lstFilePathPtr,
  }
  
  asm.Assemble(cfg)
}
