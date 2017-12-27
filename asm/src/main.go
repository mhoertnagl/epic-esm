package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type Symbol struct {
	lineNo uint32
	addr   uint32
}

type SymbolTable map[string]Symbol

func NewSymbolTable() SymbolTable {
	return make(map[string]Symbol)
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
		buf.WriteString(fmt.Sprintf("%20s: lineNo=[%d] address=[%d]", l, s.lineNo, s.addr))
	}
	return buf.String()
}

func main() {
	path := "examples/test.esm"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := uint32(1)
	ip := uint32(0)
	t := NewSymbolTable()
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(scanner.Text())

		root, err := Parse(path, scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch root.([]interface{})[1].(type) {
		case *Comment:
			break
		case *Label:
			t.Add(root.([]interface{})[1].(*Label).name, ip, lineNo)
			break
		case *Instruction:
			ip++
			break
		}
		// parser := &EsmParser{Buffer: scanner.Text()}
		// parser.Init()
		//
		// if err := parser.Parse(); err != nil {
		// 	log.Fatal(err)
		// }
		//
		// ast := parser.AST()
		// fmt.Println(ast.pegRule)
		//parser.PrintSyntaxTree()

		lineNo++
	}

	fmt.Println("Symbol Table:")
	fmt.Println(t)
}
