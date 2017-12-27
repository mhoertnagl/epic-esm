package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanSymbols(path string, file *os.File) SymbolTable {
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

func compile(path string, file *os.File, t SymbolTable) SymbolTable {
	scanner := bufio.NewScanner(file)
	lineNo := uint32(1)
	ip := uint32(0)

	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		root, err := Parse(path, scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}

		node := root.([]interface{})[1]
		switch node.(type) {
		case *RegInstruction:
			ip++
			break
		case *ImmInstruction:
			ip++
			break
		case *BraInstruction:
			ip++
			break
		default:
			break
		}

		lineNo++
	}
	return t
}

func main() {
	path := "examples/test.esm"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := scanSymbols(path, file)

	fmt.Println("Symbol Table:")
	fmt.Println(t)
}
