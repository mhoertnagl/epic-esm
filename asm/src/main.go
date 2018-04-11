package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanSymbols(filename string) SymbolTable {
	file, err := os.Open(filename)
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

		root, err := Parse(filename, scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}

		node := root.([]interface{})[1]
		switch node.(type) {
		case *Comment:
			break
		case *Label:
			label := node.(*Label)
			t.Add(label.name, ip, lineNo)
			break
		case *RegInstruction, *ImmInstruction, *BraInstruction:
			ip++
			break
		}

		lineNo++
	}
	return t
}

func compile(filename string, st SymbolTable) {
	gen := NewCodeGen(filename, st)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		root, err := Parse(filename, scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}

		node := root.([]interface{})[1]
		switch node.(type) {
		case *Comment:
			//comment := node.(*Comment)
			//fmt.Printf("         @%s\n", comment)
			fmt.Printf("            %s\n", scanner.Text())
			break
		case *Label:
			label := node.(*Label)
			fmt.Printf("            %s\n", label.name)
			break
		case *RegInstruction, *ImmInstruction, *BraInstruction:
			code, ok := gen.Generate(node)
			if ok {
				fmt.Printf("0x%08x  %s\n", code, scanner.Text())
			}
			break
		}
	}
}

func main() {
	filename := "examples/test.esm"

	st := scanSymbols(filename)
	compile(filename, st)

	fmt.Println("Symbol Table:")
	fmt.Println(st)
}
