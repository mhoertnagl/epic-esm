package main

import (
	"bytes"
	"fmt"
)

type EsmError struct {
	line   string
	lineNo uint32
	chrPos uint32
	msg    string
}

/**
 * ERROR(s) in <filename>
 *   <line>
 *   [<lineNo>:<chrPos>] <msg>
 */
func (err *EsmError) Error() string {
	return fmt.Sprintf("  %s\n  [%3d:%2d]: %s", err.line, err.lineNo, err.chrPos, err.msg)
}

type ErrorScroll struct {
	fileName string
	line     string
	lineNo   uint32
	chrPos   uint32
	errors   []EsmError
}

func NewErrorScroll(fileName string, maxErrors int) *ErrorScroll {
	return &ErrorScroll{fileName, "", 0, 0, make([]EsmError, maxErrors)}
}

func (s *ErrorScroll) Next(line string) {
	s.line = line
	s.lineNo = s.lineNo + 1
}

func (s *ErrorScroll) NewError(format string, a ...interface{}) {
	if len(s.errors) > cap(s.errors) {
		panic("Too many errors.")
	}
	msg := fmt.Sprintf(format, a)
	s.errors = append(s.errors, EsmError{s.line, s.lineNo, s.chrPos, msg})
}

func (s *ErrorScroll) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString("ERROR(s) in ")
	buffer.WriteString(s.fileName)
	buffer.WriteString("\n")
	for _, err := range s.errors {
		buffer.WriteString(err.Error())
		buffer.WriteString("\n")
	}
	return buffer.String()
}
