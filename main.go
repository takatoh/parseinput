package main

import (
	"fmt"
	"os"

	"./inputparser"
)

func main() {
	infile, _ := os.Open(os.Args[1])
	input := inputparser.Parse(infile)

//	fmt.Printf("GAMMA_R = %v\n", input.Gamma_r)
//	fmt.Printf("H_MAX   = %v\n", input.H_max)
//	fmt.Printf("PLOT    = %v\n", input.Plot)
	fmt.Printf("%#v\n", input)
}
