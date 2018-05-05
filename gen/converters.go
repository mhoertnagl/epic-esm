package gen

import "strconv"

const (
	BRA_MIN = -(1 << 24)
	BRA_MAX = 1 << 24
)

func (g *CodeGen) convertSignedNum(n string, s uint8, p uint8) uint32 {
	return g.convertNum(n, s, p, -(1 << p), 1<<p)
}

func (g *CodeGen) convertUnsignedNum(n string, s uint8, p uint8) uint32 {
	return g.convertNum(n, s, p, 0, 1<<p)
}

func (g *CodeGen) convertNum(n string, s uint8, p uint8, min int64, max int64) uint32 {
	i, err := g.parseNum(n)

	if err != nil {
		g.Error("Number [%s] too long.", n)
	}
	if i < min {
		g.Error("Unexpected number [%s]. Number must be greater than [%d].", n, min)
	}
	if i >= max {
		g.Error("Unexpected number [%s]. Number must be less than [%d]", n, max)
	}
	return g.place(i, s, p)
}

func (g *CodeGen) convertAddr(addr uint32) uint32 {
	bra := int64(addr - g.ip)
	if bra < BRA_MIN || bra >= BRA_MAX {
		g.Error("Branch distance [%d] too large.", bra)
	}
	return g.place(bra, 0, 25)
}

func (g *CodeGen) parseNum(n string) (int64, error) {
	if len(n) > 2 && n[0:2] == "0b" {
		return strconv.ParseInt(n[2:], 2, 32)
	}
	if len(n) > 2 && n[0:2] == "0x" {
		return strconv.ParseInt(n[2:], 16, 32)
	}
	return strconv.ParseInt(n, 10, 32)
}

func (g *CodeGen) place(i int64, s uint8, p uint8) uint32 {
	return uint32((i & ((1 << p) - 1)) << s)
}
