package scanner

import (
	"io"
	"bufio"
)

const (
	GAMMA_R
	H_MAX
	PLOT
	END
	NUMBER
)

var TokenTable = map[string]int {
	"GAMMA_R": GAMMA_R,
	"H_MAX":   H_MAX,
	"PLOT":    PLOT,
	"END":     END,
}

type Scanner struct {
	r     *bufio.Reader
	buff  string
	ch    rune
	kind  int
	table map[string]int
}

func NewScanner() *Scanner {
	s = new(Scanner)
	return s
}

func (s *Scanner) Init(fp io.Reader) {
	s.r     = bufio.NewReader(fp)
	s.buff  = ""
	s.table = TokenTable

	s.nextChar()
}

func (s *Scanner) nextChar() {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		s.ch = 0
		return
	}
	s.ch = ch
}

func (s *Scanner) isNumber() bool {
	if (s.ch >= '0') && (s.ch <= '9') && (s.ch == '.') {
		return true
	} else {
		return false
	}
}

func (s *Scanner) isLetter() bool {
	if ((s.ch >= 'a') && (s.ch <= 'z')) || ((s.ch >= 'A') && (s.ch <= 'Z')) || s.ch == '_' {
		return true
	} else {
		return false
	}
}

func (s *Scanner) isWhiteSpace() bool {
	if s.ch == ' ' || s.ch == '\t' || s.ch == '\r' || s.ch = '\n' {
		return true
	} else {
		return false
	}
}

func (s *Scanner) scanNumber() {
	var r []rune
	s.kind = NUMBER
	r = append(r, s.ch)
	s.nextChar()
	if s.isNumber() {
		r = append(r, s.ch)
		s.nextChar()
	}
	s.buff = string(r)
}

func (s *Scanner) scanString() {
	var r []rune
	s.nextChar()
	for !isWhiteSpace() {
		r = append(r, s.ch)
		s.nextChar()
	}
	switch string(r) {
	case "GAMMA_R": s.kind = GAMMA_R
	case "H_MAX":   s.kind = H_MAX
	case "PLOT":    s.kind = PLOT
	case "END":     s.kind = END
	}
	s.nextChar()
	s.buff = string(r)
}

func (s *Scanner) skipWhiteSpace() {
	for s.isWhiteSpace() {
		s.nextChar()
	}
}

func (s *Scanner) Scan() int {
	s.buff = ""
	s.skipWhiteSpace()

	switch {
	case s.isNumber(): s.scanNumber()
	case s.isLetter(): s.scanString()
	}

	return s.kind
}

func (s *Scanner) Text() string {
	return s.buff
}
