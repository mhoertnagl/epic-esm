package main

var conds = []string{"al", "eq", "ne", "lt", "gt", "le", "ge", "nv"}

var contract = CodeGeneratorContract{
	PatternTable{
		"cmd":  `\w*[a-z]{3}`,
		"c?":   "(al|eq|ne|gt|lt|ge|le|nv)?",
		"[":    `\[`,
		"]":    `\]`,
		"rd":   "r([0-9]|1[0-5])",
		"ra":   "r([0-9]|1[0-5])",
		"rb":   "r([0-9]|1[0-5])",
		"si12": "-?[0-9]+|0x[0-9a-f]+",
		"@lbl": "@[a-zA-Z0-9]",
	},
	ConversionsTable{
		"c?":   convertStr(conds, 25, 3),
		"rd":   convertReg(20, 4),
		"ra":   convertReg(16, 4),
		"rb":   convertReg(12, 4),
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
