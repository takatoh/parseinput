.PHONY: clean

build: parseinput.exe

parseinput.exe: main.go inputparser/inputparser.go
	go build

inputparser/inputparser.go: inputparser/parser.go.y scanner/scanner.go
	go tool yacc -o inputparser/inputparser.go inputparser/parser.go.y

clean:
	rm *.exe
	rm *.output
	rm inputparser/inputparser.go
