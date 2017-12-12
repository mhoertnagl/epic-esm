package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

var text = `
The quick brown fox jumps over the lazy dog #1.
Быстрая коричневая лиса перепрыгнула через ленивую собаку.
`

type tokentype int

const (
  EOF tokentype = iota
  COMMAND 
  COMMAND_MODIFIER
  REGISTER
  NUMBER
  LABEL_DECLARATION
  LABEL_REFERENCE
  SINGLE_LINE_COMMENT
  MULTI_LINE_COMMENT
)

type Token struct {
  typ    tokentype
  lexeme string
  number uint32
}

// type EsmScanner interface {
// 
// }

type EsmScanner struct {
  reader *bufio.Reader
  lineNo int
  chrPos int
}

func NewEsmScanner(reader *bufio.Reader) *EsmScanner {
  return &EsmScanner{reader, 0, 0}
}

func (scanner *EsmScanner) NextToken() *Token {
  token := &Token{EOF, "", 0}
  for {
    c, sz, err := scanner.reader.ReadRune()
    if err == nil {
      // Hier state weitermachen
      fmt.Printf("%q [%d]\n", string(c), sz)
    } else if err == io.EOF {
      break
    } else {
      log.Fatal(err)
    }
  }
  return token
}

/**
 * Test
 */
func main() {
	reader := bufio.NewReader(strings.NewReader(text))
  scanner := NewEsmScanner(reader)
  fmt.Println(scanner)
	// for {
	// 	if c, sz, err := r.ReadRune(); err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			log.Fatal(err)
	// 		}
	// 	} else {
	// 		fmt.Printf("%q [%d]\n", string(c), sz)
	// 	}
	// }
}
