package main

var conds = []string{"al", "eq", "ne", "lt", "gt", "le", "ge", "nv"}

var conditions = ConversionMapping{
	mapping: ConversionMap{
		"al": 0,
		"eq": 1,
		"ne": 2,
		"lt": 3,
		"gt": 4,
		"le": 5,
		"ge": 6,
		"nv": 7,
	},
	hasDefValue: true,
	defVal:      0,
	errorMsg:    "Unexpected condition flag [%s]. Expecting one of %s.",
}

var registers = ConversionMapping{
	mapping: ConversionMap{
		"r0":  0,
		"r1":  1,
		"r2":  2,
		"r3":  3,
		"r4":  4,
		"r5":  5,
		"r6":  6,
		"r7":  7,
		"r8":  8,
		"r9":  9,
		"r10": 10,
		"r11": 11,
		"r12": 12,
		"r13": 13,
		"r14": 14,
		"r15": 15,
	},
	hasDefValue: false,
	errorMsg:    "Unexpected register name [%s]. Expecting one of %s.",
}

var contract = CodeGeneratorContract{
	PatternTable{
		//"ws":   `\w*`,
		"cmd":  `[a-z]{3}`,
		"c?":   `(al|eq|ne|gt|lt|ge|le|nv)?`,
		"[":    `\[`,
		"]":    `\]`,
		"rd":   `r([0-9]|1[0-5])`,
		"ra":   `r([0-9]|1[0-5])`,
		"rb":   `r([0-9]|1[0-5])`,
		"si12": `-?[0-9]+|0x[0-9a-f]+`,
		"ui5":  `[0-9]+`,
		"@lbl": `@[a-zA-Z0-9]`,
	},
	ConversionsTable{
		"c?":   convertMap(conditions, 25, 3),
		"rd":   convertMap(registers, 20, 4),
		"ra":   convertMap(registers, 16, 4),
		"rb":   convertMap(registers, 12, 4),
		"si12": convertSignedNum(4, 12),
		"ui5":  convertUnsignedNum(4, 5),
		// "@lbl":    "@[a-zA-Z0-9]",
	},
	TranslationTable{
		"add r r r": {"!? add c? rd ra rb", 0x00000000},
		"add r r i": {"!? add c? rd ra si12", 0x00000000},
		"sub r r r": {"!? sub c? rd ra rb", 0x00000000},
		"sub r r i": {"!? sub c? rd ra si12", 0x00000000},
	},
	// TranslationTable{
	// 	"add": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"sub": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"mul": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"div": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"sll": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra ui5", 0x00000000}},
	// 	"srl": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra ui5", 0x00000000}},
	// 	"sra": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra ui5", 0x00000000}},
	// 	"ror": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra ui5", 0x00000000}},
	// 	"and": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"orr": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"xor": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"bcl": { // Bit Clear (a AND NOT b)
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	"tst": {
	// 		{"c? rd ra rb", 0x00000000},
	// 		{"c? rd ra si12", 0x00000000}},
	// 	// "tst": {
	// 	//   {"c? rd ra rb", 0x00000000},
	// 	//   {"c? rd ra si12", 0x00000000}},
	// 	// "tst": {
	// 	//   {"c? rd ra rb", 0x00000000},
	// 	//   {"c? rd ra si12", 0x00000000}},
	// 	// "tst": {
	// 	//   {"c? rd ra rb", 0x00000000},
	// 	//   {"c? rd ra si12", 0x00000000}},
	// 	"ldw": {
	// 		{"c? rd ra [ rb ]", 0x00000000},
	// 		{"c? rd ra [ si12 ]", 0x00000000}},
	// 	"stw": {
	// 		{"c? rd ra [ rb ]", 0x00000000},
	// 		{"c? rd ra [ si12 ]", 0x00000000}},
	// 	"bra": {
	// 		{"c? @lbl", 0x00000000}},
	// 	"brl": {
	// 		{"c? @lbl", 0x00000000}},
	// },
}
