.PHONY: clean

build: main.go inputparser/y.go
	go build

inputparser/y.go: inputparser/parser.go.y scanner/scanner.go
	goyacc -o inputparser/y.go inputparser/parser.go.y

clean:
	if exist parseinput.exe rm parseinput.exe
	rm *.output
	rm inputparser/y.go
