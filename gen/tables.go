package gen

const (
	BRA_MIN = -(1 << 24)
	BRA_MAX = 1 << 24
)

var dataInstructions = map[string]uint32{
	"add": 0x00000000,
	"sub": 0x00000001,
	"mul": 0x00000002,
	"div": 0x00000003,

	"and": 0x00000004,
	"oor": 0x00000005,
	"xor": 0x00000006,
	"nor": 0x00000007,

	"adu": 0x00000008,
	"sbu": 0x00000009,
	//"mlu": 0x0000000a, multiplikation ist immer signed
	//"dvu": 0x0000000b, division ist immer signed

	"cmp": 0x0000000c,
	"cpu": 0x0000000d,
	"tst": 0x0000000e,
	"mov": 0x0000000f,
}

var memInstructions = map[string]uint32{
	"stw": 0x40000000,
	"ldw": 0x40000001,
}

var branchInstructions = map[string]uint32{
	"bra": 0xe0000000,
	"brl": 0xe2000000,
}

var shiftOps = map[string]uint32{
	"<<":  0,
	">>":  1,
	">>>": 2,
	"<<>": 3,
	"<>>": 3,
	"sll": 0,
	"srl": 1,
	"sra": 2,
	"rol": 3,
	"ror": 3,
}

var registers = map[string]uint32{
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
	"sp":  13,
	"rp":  14,
	"ip":  15,
}

var conditions = map[string]uint32{
	"nv": 0,
	"eq": 1,
	"lt": 2,
	"le": 3,
	"gt": 4,
	"ge": 5,
	"ne": 6,
	"al": 7,
}
