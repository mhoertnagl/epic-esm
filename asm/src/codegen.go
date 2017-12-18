package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PatternTable map[string]string

type ConversionFunction func(string) (uint32, error)

type ConversionsTable map[string]ConversionFunction

type TranslationEntry struct {
	args     string
	template uint32
}

type TranslationTable map[string][]TranslationEntry

type MatchTable map[string]*[]regexp.Regexp

type CodeGeneratorContract struct {
	patterns     PatternTable
	conversions  ConversionsTable
	translations TranslationTable
}

func place(i int64, s uint8, p uint8) (uint32, error) {
	j := i & ((1 << p) - 1)
	if j != i {
		// i cannot be represented with p bits.
		return 0, nil
	}
	return uint32(j << s), nil
}

func convertStr(arr []string, s uint8, p uint8) ConversionFunction {
	return func(r string) (uint32, error) {
		for i, v := range arr {
			if v == r {
				return place(int64(i), s, p)
			}
		}
		// Conversion error: Value can only be one of [arr]
		return 0, nil
	}
}

func convertReg(s uint8) ConversionFunction {
	return func(r string) (uint32, error) {
		if r[0] == 'r' {
			i, err := strconv.ParseInt(r[1:], 10, 32)
			if i < 0 || i > 15 {
				// so ein register gibs nicht.
			}
			return place(i, s, 4)
		}
		// Not a number
		return 0, nil
	}
}

func convertNum(s uint8, p uint8) ConversionFunction {
	return func(n string) (uint32, error) {
		if n[0:2] == "0x" {
			i, err := strconv.ParseInt(n, 16, 32)
			return place(i, s, p)
		} else {
			i, err := strconv.ParseInt(n, 10, 32)
			return place(i, s, p)
		}
	}
}

var conds = []string{"al", "eq", "ne", "lt", "gt", "le", "ge", "nv"}

var contract = CodeGeneratorContract{
	PatternTable{
		"c?": "(al|eq|ne|gt|lt|ge|le|nv)?",
		// "[rb]":   `\[r([0-9]|1[0-5])\]`,
		// "[si12]": `\[-?[0-9]+|0x[0-9a-f]+\]`,
		"[":    `\[`,
		"]":    `\]`,
		"rd":   "r([0-9]|1[0-5])",
		"ra":   "r([0-9]|1[0-5])",
		"rb":   "r([0-9]|1[0-5])",
		"si12": "-?[0-9]+|0x[0-9a-f]+",
		"@lbl": "@[a-zA-Z0-9]",
	},
	ConversionsTable{
		"c?": convertStr(conds, 25, 3),
		// "[rb]":   convertReg(12),
		// "[si12]": convertNum(4, 12),
		"rd":   convertReg(20),
		"ra":   convertReg(16),
		"rb":   convertReg(12),
		"si12": convertNum(4, 12),
		// "@lbl":    "@[a-zA-Z0-9]",
	},
	TranslationTable{
		"add": {
			{"c? rd ra rb", 0x00000000},
			{"c? rd ra si12", 0x00000000}},
		"sll": {
			{"c? rd ra rb", 0x00000000},
			{"c? rd ra si12", 0x00000000}},
		"tst": {
			{"c? rd ra rb", 0x00000000},
			{"c? rd ra si12", 0x00000000}},
		"ldw": {
			{"c? rd ra [ rb ]", 0x00000000},
			{"c? rd ra [ si12 ]", 0x00000000}},
		"stw": {
			{"c? rd ra [ rb ]", 0x00000000},
			{"c? rd ra [ si12 ]", 0x00000000}},
		"brl": {
			{"c? @lbl", 0x00000000}},
	},
}

type CodeGenerator struct {
	contract CodeGeneratorContract
}

func (ge *CodeGenerator) MakeCodeGenerator(contract CodeGeneratorContract) *CodeGenerator {
	return &CodeGenerator{contract}
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
			val, err := conv(ins[loc[0]:loc[1]])
			if err != nil {
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
