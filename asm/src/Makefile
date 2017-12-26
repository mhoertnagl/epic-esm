PEG=~/go/bin/peg

run: main.go grammar.peg.go
	go run main.go grammar.peg.go

build: grammar.peg.go
	go build -o esm

grammar.peg.go: grammar.peg
	$(PEG) grammar.peg

clean:
	rm -f esm debug grammar.peg.go
