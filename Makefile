.PHONY: clean

build: main.go inputparser/parser.go
	go build

inputparser/parser.go: inputparser/parser.go.y scanner/scanner.go
	goyacc -o inputparser/parser.go inputparser/parser.go.y

clean:
	if exist parseinput.exe rm parseinput.exe
	rm *.output
	rm inputparser/parser.go
