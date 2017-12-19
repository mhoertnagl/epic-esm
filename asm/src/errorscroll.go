package main

import (
	"bytes"
	"fmt"
)

type EsmError struct {
	lineNo uint32
	chrPos uint32
	msg    string
}

func (err *EsmError) Error() string {
	return fmt.Sprintf("ERROR [%3d:%2d]: %s", err.lineNo, err.chrPos, err.msg)
}

type ErrorScroll struct {
	fileName string
	lineNo   uint32
	chrPos   uint32
	errors   []EsmError
}

func NewErrorScroll(fileName string, maxErrors int) *ErrorScroll {
	return &ErrorScroll{fileName, 0, 0, make([]EsmError, maxErrors)}
}

func (s *ErrorScroll) NewError(format string, a ...interface{}) {
	if len(s.errors) > cap(s.errors) {
		panic("Too many errors.")
	}
	msg := fmt.Sprintf(format, a)
	s.errors = append(s.errors, EsmError{s.lineNo, s.chrPos, msg})
}

func (s *ErrorScroll) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.fileName)
	buffer.WriteString("\n")
	for _, err := range s.errors {
		buffer.WriteString(err.Error())
		buffer.WriteString("\n")
	}
	return buffer.String()
}
