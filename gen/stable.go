package gen

import (
	"bytes"
	"fmt"
)

type Symbol struct {
	lineNo uint32
	addr   uint32
}

type SymbolTable map[string]Symbol

func NewSymbolTable() SymbolTable {
	return make(SymbolTable)
}

func (t SymbolTable) Add(label string, addr uint32, lineNo uint32) {
	s, defined := t[label]
	if defined {
		fmt.Printf("[%d:] ERROR: Label [%s] already defined at line [%d].\n", lineNo, label, s.lineNo)
	}
	t[label] = Symbol{lineNo: lineNo, addr: addr}
}

func (t SymbolTable) Find(label string) (s Symbol, ok bool) {
	s, ok = t[label]
	return s, ok
}

func (t SymbolTable) String() string {
	var buf bytes.Buffer
	for l, s := range t {
		buf.WriteString(fmt.Sprintf("%20s: lineNo=[%d] address=[%d]\n", l, s.lineNo, s.addr))
	}
	return buf.String()
}
