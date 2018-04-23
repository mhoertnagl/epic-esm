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

func WriteInt32BigEndian(w *bufio.Writer, i uint32) {
	w.WriteByte(byte(i >> 24))
	w.WriteByte(byte(i >> 16))
	w.WriteByte(byte(i >> 8))
	w.WriteByte(byte(i))
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
	outWriter := bufio.NewWriter(outFile)

	lstFile, err := os.Create(lstFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer lstFile.Close()
	lstWriter := bufio.NewWriter(lstFile)

	ip := uint32(0)

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
			//fmt.Printf("%24s%s\n", "", scanner.Text())
			break
		case *Label:
			label := node.(*Label)
			fmt.Fprintf(lstWriter, "%24s%s\n", "", label.name)
			//fmt.Printf("%24s%s\n", "", label.name)
			break
		case *RegInstruction, *I12Instruction, *I16Instruction, *BraInstruction:
			code, ok := gen.Generate(node)
			if ok {
				fmt.Fprintf(lstWriter, "0x%08x  0x%08x  %s\n", ip, code, scanner.Text())
				WriteInt32BigEndian(outWriter, code)
				//fmt.Printf("0x%08x  0x%08x  %s\n", ip, code, scanner.Text())
			}
			ip++
			break
		}
		lstWriter.Flush()
		outWriter.Flush()
	}
}

func main() {
	outFileNamePtr := flag.String("o", "x.bin", "an output file")
	lstFileNamePtr := flag.String("l", "x.lst", "a listing file")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		panic("Too few arguments. Provide a single input file.")
	}

	if len(args) != 1 {
		panic("Too many arguments. Provide only a single input file.")
	}

	inFileName := args[0]

	st := scanSymbols(inFileName)
	compile(inFileName, st, *outFileNamePtr, *lstFileNamePtr)

	// fmt.Println("Symbol Table:")
	// fmt.Println(st)
}
