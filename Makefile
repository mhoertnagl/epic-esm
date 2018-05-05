# On Linux
#PEG          = ~/go/bin/pigeon
# On Windows
PEG          = D:/Matse/source/go/bin/pigeon

GRAMMAR_FILE = grammar.peg
GO_FILES = *.go #main.go parser.go parserutils.go ast.go symboltable.go codegen.go converters.go

run: parser.go $(GO_FILES)
	go run $^

build: parser.go
# On Linux
#	go build -o esm
# On Windows
	go build -o esm.exe

parser:
	$(PEG) -o parser.go $(GRAMMAR_FILE)

parser.go:
	$(PEG) -o parser.go $(GRAMMAR_FILE)

clean:
	rm -f esm debug parser.go
