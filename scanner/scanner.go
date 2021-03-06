package scanner

import (
	"os"
	"bufio"
)

const (
	STRING  = 57346
	NUMBER  = 57347
	MODEL   = 57348
	GAMMA_R = 57349
	H_MAX   = 57350
	PLOT    = 57351
	END     = 57352
)

var TokenTable = map[string]int {
	"*MODEL":    MODEL,
	"*GAMMA0.5": GAMMA_R,
	"*HMAX":     H_MAX,
	"*PLOT":     PLOT,
	"*END":      END,
}

type Scanner struct {
	r     *bufio.Reader
	buff  string
	ch    rune
	kind  int
	table map[string]int
	end   bool
}

func (s *Scanner) Init(fp *os.File) {
	s.r     = bufio.NewReader(fp)
	s.buff  = ""
	s.table = TokenTable
	s.end   = false

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

func (s *Scanner) isLabel() bool {
	if s.ch == '*' {
		return true
	} else {
		return false
	}
}

func (s *Scanner) isNumber() bool {
	if ((s.ch >= '0') && (s.ch <= '9')) || s.ch == '.' {
		return true
	} else {
		return false
	}
}

func (s *Scanner) isLetter() bool {
	if ((s.ch >= 'A') && (s.ch <= 'Z')) {
		return true
	} else {
		return false
	}
}

func (s *Scanner) isWhiteSpace() bool {
	if s.ch == ' ' || s.ch == '\t' || s.ch == '\r' || s.ch == '\n' {
		return true
	} else {
		return false
	}
}

func (s *Scanner) isComment() bool {
	if s.ch == '/' {
		s.nextChar()
		if s.ch == '/' {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (s *Scanner) isEOF() bool {
	if s.ch == 0 {
		return true
	} else {
		return false
	}
}

func (s *Scanner) scanLabel() {
	var r []rune
	r = append(r, s.ch)
	s.nextChar()
	for !s.isWhiteSpace() {
		r = append(r, s.ch)
		s.nextChar()
	}
	s.kind = s.table[string(r)]
	s.buff = string(r)
	if string(r) == "*END" {
		s.end = true
	}
}

func (s *Scanner) scanNumber() {
	var r []rune
	s.kind = NUMBER
	r = append(r, s.ch)
	s.nextChar()
	for s.isNumber() {
		r = append(r, s.ch)
		s.nextChar()
	}
	s.buff = string(r)
}

func (s *Scanner) scanString() {
	var r []rune
	r = append(r, s.ch)
	s.nextChar()
	for s.isLetter() {
		r = append(r, s.ch)
		s.nextChar()
	}
	s.kind = STRING
	s.buff = string(r)
}

func (s *Scanner) skipWhiteSpace() {
	for s.isWhiteSpace() {
		s.nextChar()
	}
}

func (s *Scanner) skipComment() {
	s.skipWhiteSpace()
	if s.isComment() {
		for !(s.ch == '\n') {
			s.nextChar()
		}
	}
}

func (s *Scanner) Scan() int {
	s.buff = ""
	s.skipComment()
	s.skipWhiteSpace()

	switch {
	case s.end:        s.kind = 0
	case s.isLabel():  s.scanLabel()
	case s.isNumber(): s.scanNumber()
	case s.isLetter(): s.scanString()
	case s.isEOF():    s.kind = 0
	}

	return s.kind
}

func (s *Scanner) Text() string {
	return s.buff
}
