package aut

//go:generate goyacc -p aut -o parser.y.go aut.y

import "io"

// Lexer for aut.
type Lexer struct {
	scanner *Scanner
	Errors  chan error
}

// NewLexer returns a new yacc-compatible lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{scanner: NewScanner(r), Errors: make(chan error, 1)}
}

func (l *Lexer) Lex(yylval *autSymType) int {
	token := l.scanner.Scan()
	switch token := token.(type) {
	case *DigitsToken:
		yylval.num = token.num
	case *LabelToken:
		yylval.str = token.str
	}
	return int(token.Tok())
}

// Error handles error.
func (l *Lexer) Error(err string) {
	l.Errors <- &ErrParse{Err: err, Pos: l.scanner.pos}
}
