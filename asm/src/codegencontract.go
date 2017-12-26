package main

import (
	"regexp"
	"strconv"
)

type PatternTable map[string]string

type ConversionMap map[string]int64

type ConversionMapping struct {
	mapping     map[string]int64
	hasDefValue bool
	defVal      int64
	errorMsg    string
}

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

func convertMap(conMap ConversionMapping, s uint8, p uint8) ConversionFunction {
	return func(r string, scroll *ErrorScroll) (uint32, bool) {
		val, ok := conMap.mapping[r]
		if ok {
			return place(val, s, p)
		}
		if conMap.hasDefValue {
			return place(conMap.defVal, s, p)
		}
		keys := make([]string, len(conMap.mapping))
		i := 0
		for k := range conMap.mapping {
			keys[i] = k
			i++
		}
		scroll.NewError(conMap.errorMsg, r, keys)
		return 0, false
	}
}

func convertSignedNum(s uint8, p uint8) ConversionFunction {
	return convertNum(s, p, -(1 << p), 1<<p)
}

func convertUnsignedNum(s uint8, p uint8) ConversionFunction {
	return convertNum(s, p, 0, 1<<p)
}

func convertNum(s uint8, p uint8, min int64, max int64) ConversionFunction {
	return func(n string, scroll *ErrorScroll) (uint32, bool) {
		i, err := parseNum(n)

		if err != nil {
			scroll.NewError("Unexpected number [%s]. Invalid syntax.", n)
			return 0, false
		}
		if i < min {
			scroll.NewError("Unexpected number [%s]. Number must be greater than [%d].", n, min)
			return 0, false
		}
		if i >= max {
			scroll.NewError("Unexpected number [%s]. Number must be less than [%d]", n, max)
			return 0, false
		}
		return place(i, s, p)
	}
}

func parseNum(n string) (int64, error) {
	if n[0:2] == "0x" {
		return strconv.ParseInt(n[2:], 16, 32)
	}
	return strconv.ParseInt(n, 10, 32)
}

func place(i int64, s uint8, p uint8) (uint32, bool) {
	return uint32((i & ((1 << p) - 1)) << s), true
}
