package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanSymbols(inFileName string) SymbolTable {
	file, err := os.Open(inFileName)
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

		root, err := Parse(inFileName, scanner.Bytes())
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

func compile(inFileName string, st SymbolTable, outFileName string, lstFileName string) {
	gen := NewCodeGen(inFileName, st)

	inFile, err := os.Open(inFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	binWriter := bufio.NewWriter(outFile)

	lstFile, err := os.Create(lstFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer lstFile.Close()
	lstWriter := bufio.NewWriter(lstFile)

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(scanner.Text()) == "" {
			fmt.Fprintln(lstWriter)
			fmt.Println()
			continue
		}

		root, err := Parse(inFileName, scanner.Bytes())
		if err != nil {
			fmt.Println(err)
			continue
		}

		node := root.([]interface{})[1]
		switch node.(type) {
		case *Comment:
			fmt.Fprintf(lstWriter, "%24s%s\n", "", scanner.Text())
			break
		case *Label:
			label := node.(*Label)
			fmt.Fprintf(lstWriter, "%24s%s\n", "", label.name)
			break
		case Instruction:
			ip := gen.ip
			codes := gen.Generate(node.(Instruction))
			//if ok {
			for i, code := range codes {
				text := " +"
				if i == 0 {
					text = scanner.Text()
				}
				fmt.Fprintf(lstWriter, "0x%08x  0x%08x  %s\n", ip+uint32(i), code, text)
				WriteInt32BigEndian(binWriter, code)
			}
			//}
			break
		}
		lstWriter.Flush()
		binWriter.Flush()
	}
}

func main() {
	outFileNamePtr := flag.String("o", "x.bin", "an output file")
	lstFileNamePtr := flag.String("l", "x.lst", "a listing file")

	flag.Parse()

	//args := flag.Args()

	// if len(args) < 1 {
	// 	panic("Too few arguments. Provide a single input file.")
	// }
	//
	// if len(args) > 1 {
	// 	panic("Too many arguments. Provide only a single input file.")
	// }

	inFileName := "examples/test.esm" //args[0]

	st := scanSymbols(inFileName)
	compile(inFileName, st, *outFileNamePtr, *lstFileNamePtr)

	// fmt.Println("Symbol Table:")
	// fmt.Println(st)
}
