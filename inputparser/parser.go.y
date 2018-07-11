%{
package inputparser

import (
	"os"
	"strconv"

	"github.com/takatoh/parseinput/scanner"
)

type InputData struct {
	Model   string
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
	string  string
	num     float64
	numlist []float64
}

%type<input> input
%type<string> model
%type<num> gamma_r
%type<num> h_max
%type<numlist> plot
%type<string> string
%type<numlist> numlist
%type<num> num
%token<token> STRING
%token<token> NUMBER
%token<token> MODEL
%token<token> GAMMA_R
%token<token> H_MAX
%token<token> PLOT
%token<token> END

%%

input
	: model gamma_r h_max plot END
	{
		$$ = InputData{ Model: $1, Gamma_r: $2, H_max: $3, Plot: $4 }
		yylex.(*Lexer).result = $$
	}

model
	: MODEL string
	{
		$$ = $2
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

string
	: STRING
	{
		$$ = $1.literal
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
