package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
		case *RegInstruction, *I12Instruction, *I16Instruction, *BraInstruction:
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

	ip := uint32(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// fmt.Println("==============================")
		// fmt.Println(scanner.Text())
		// fmt.Println("------------------------------")

		if strings.TrimSpace(scanner.Text()) == "" {
			//fmt.Println("Empty line.")
			fmt.Println()
			continue
		}

		root, err := Parse(filename, scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}

		node := root.([]interface{})[1]
		switch node.(type) {
		case *Comment:
			fmt.Printf("%24s%s\n", "", scanner.Text())
			break
		case *Label:
			label := node.(*Label)
			fmt.Printf("%24s%s\n", "", label.name)
			break
		case *RegInstruction, *I12Instruction, *I16Instruction, *BraInstruction:
			code, ok := gen.Generate(node)
			if ok {
				fmt.Printf("0x%08x  0x%08x  %s\n", ip, code, scanner.Text())
			}
			ip++
			break
		}
	}
}

func main() {
	//outFilePtr := flag.String("o", "x.bin", "an output file")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		panic("Too few arguments. Provide a single input file.")
	}

	if len(args) != 1 {
		panic("Too many arguments. Provide only a single input file.")
	}

	inFile := args[0]

	st := scanSymbols(inFile)
	compile(inFile, st)

	fmt.Println("Symbol Table:")
	fmt.Println(st)
}
