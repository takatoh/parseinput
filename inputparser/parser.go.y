%{
package inputparser

import (
	"os"
	"strconv"

	"github.com/takatoh/parseinput/scanner"
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
	tok := l.Scan()
	lval.token = Token{ token: tok, literal : l.Text() }
	return tok
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func Parse(infile *os.File) InputData {
	l := new(Lexer)
	l.Init(infile)
	yyParse(l)
	return l.result
}
