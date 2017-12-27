package main

import (
	"fmt"
	"os"

	"github.com/takatoh/parseinput/inputparser"
)

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Cannot open file:", os.Args[1])
		os.Exit(1)
	}
	defer infile.Close()

	input := inputparser.Parse(infile)

	fmt.Printf("%#v\n", input)
}
