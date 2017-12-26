package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("examples/test.esm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 1
	for scanner.Scan() {
		fmt.Println("NEW LINE:::")

		parser := &EsmParser{Buffer: scanner.Text()}
		parser.Init()

		if err := parser.Parse(); err != nil {
			log.Fatal(err)
		}

		parser.PrintSyntaxTree()
		lineNo++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
