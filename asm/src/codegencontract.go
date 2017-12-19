package main

import (
	"regexp"
	"strconv"
)

type PatternTable map[string]string

type ConversionFunction func(string, *ErrorScroll) (uint32, bool)

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

func place(i int64, s uint8, p uint8, scroll *ErrorScroll) (uint32, bool) {
	j := i & ((1 << p) - 1)
	if j != i {
		scroll.NewError("Cannot represent [%d] with [%d] bits.", i, p)
		return 0, false
	}
	return uint32(j << s), true
}

func convertStr(arr []string, s uint8, p uint8) ConversionFunction {
	return func(r string, scroll *ErrorScroll) (uint32, bool) {
		for i, v := range arr {
			if v == r {
				return place(int64(i), s, p, scroll)
			}
		}
		scroll.NewError("Unexpected argument [%s]. Expecting one of %s.", r, arr)
		return 0, false
	}
}

func convertReg(s uint8, p uint8) ConversionFunction {
	return func(r string, scroll *ErrorScroll) (uint32, bool) {
		if r[0] == 'r' {
			i, err := strconv.ParseInt(r[1:], 10, 32)
			if i < 0 || i > 15 {
				scroll.NewError("Unexpected register name [%s]. Expecting one of [r0,r1,...,r15].", r)
			}
			return place(i, s, p, scroll)
		}
		scroll.NewError("Unexpected register name [%s]. Expecting one of [r0,r1,...,r15].", r)
		return 0, false
	}
}

func convertNum(s uint8, p uint8) ConversionFunction {
	return func(n string, scroll *ErrorScroll) (uint32, bool) {
		if n[0:2] == "0x" {
			i, err := strconv.ParseInt(n, 16, 32)
			return place(i, s, p, scroll)
		} else {
			i, err := strconv.ParseInt(n, 10, 32)
			return place(i, s, p, scroll)
		}
	}
}
