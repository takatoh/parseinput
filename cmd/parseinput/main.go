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
	input := inputparser.Parse(infile)

//	fmt.Printf("GAMMA_R = %v\n", input.Gamma_r)
//	fmt.Printf("H_MAX   = %v\n", input.H_max)
//	fmt.Printf("PLOT    = %v\n", input.Plot)
	fmt.Printf("%#v\n", input)
}
