package main

import (
	"fmt"
	"regexp"
	"strings"
)

type CodeGenerator struct {
	contract CodeGeneratorContract
	scroll   *ErrorScroll
}

func NewCodeGenerator(contract CodeGeneratorContract, scroll *ErrorScroll) *CodeGenerator {
	return &CodeGenerator{contract, scroll}
}

func (gen *CodeGenerator) Compile(ins string) uint32 {
	cmd := ins[0:3]
	tra, ok := gen.contract.translations[cmd]
	if !ok {
		// Unknown command.
	}
	for _, e := range tra {
		i := 3
		asm := e.template
		for _, arg := range strings.Split(e.args, " ") {
			pat, ok := gen.contract.patterns[arg]
			if !ok {
				panic(fmt.Sprintf("FATAL ERROR: Unspecified translation argument [%s]", arg))
			}
			rex := regexp.MustCompile(pat)
			loc := rex.FindStringIndex(ins[i:])
			if loc == nil {
				// Did not match. Print an error if this was the last applyable rule.
				break
			}
			conv := gen.contract.conversions[arg]
			val, ok := conv(ins[loc[0]:loc[1]], gen.scroll)
			if !ok {
				// Conversion error.
				break
			}
			asm |= val
			i = loc[1]
		}
		// If some errors happened. Remember them and try another one else return asm.
		return asm
	}
	return 0
}
