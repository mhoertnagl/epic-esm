PEG          = ~/go/bin/pigeon
GRAMMAR_FILE = grammar.peg

run: main.go parser.go
	go run $^

build: parser.go
	go build -o esm

parser:
	$(PEG) -o parser.go $(GRAMMAR_FILE)

parser.go: 
	$(PEG) -o parser.go $(GRAMMAR_FILE)

clean:
	rm -f esm debug parser.go
