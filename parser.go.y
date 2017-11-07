%{
package main

import (
	"fmt"
	"text/scanner"
	"os"
	"bufio"
	"strconv"
)

type InputData struct {
	Gamma_r float64
	H_max   float64
	Plot    []float64
}

type Token struct {
	token   int
	literal string
}

%}

%union {
	input   InputData
	token   Token
	num     float64
	numlist []float64
}

%type<input> input
%type<num> gamma_r
%type<num> h_max
%type<numlist> plot
%type<numlist> numlist
%type<num> num
%token<token> NUMBER
%token<token> GAMMA_R
%token<token> H_MAX
%token<token> PLOT
%token<token> END

%%

input
	: gamma_r h_max plot END
	{
		$$ = InputData{ Gamma_r: $1, H_max: $2, Plot: $3 }
		yylex.(*Lexer).result = $$
	}

gamma_r
	: GAMMA_R num
	{
		$$ = $2
	}

h_max
	: H_MAX num
	{
		$$ = $2
	}

plot
	: PLOT numlist
	{
		$$ = $2
	}

numlist
	: num
	{
		$$ = []float64{ $1 }
	}
	| numlist num
	{
		$$ = append($1, $2)
	}

num
	: NUMBER
	{
		f, _ := strconv.ParseFloat($1.literal, 64)
		$$ = f
	}

%%

type Lexer struct {
	scanner.Scanner
	result InputData
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Float {
		token = NUMBER
	} else if token == scanner.Ident {
		switch l.TokenText() {
		case "GAMMA_R": token = GAMMA_R
		case "H_MAX":   token = H_MAX
		case "PLOT":    token = PLOT
		case "END":     token = END
		}
	}
	lval.token = Token{ token: token, literal: l.TokenText() }
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func main() {
	l := new(Lexer)
	infile, _ := os.Open(os.Args[1])
	input := bufio.NewReader(infile)
	l.Init(input)
	yyParse(l)
	fmt.Printf("GAMMA_H = %v\n", l.result.Gamma_r)
	fmt.Printf("H_MAX   = %v\n", l.result.H_max)
	fmt.Printf("PLOT    = %v\n", l.result.Plot)
}