.PHONY: clean

build: main.go inputparser/inputparser.go
	go build

inputparser/inputparser.go: inputparser/parser.go.y scanner/scanner.go
	goyacc -o inputparser/inputparser.go inputparser/parser.go.y

clean:
	rm *.exe
	rm *.output
	rm inputparser/inputparser.go
