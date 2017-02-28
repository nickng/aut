package aut

import "fmt"

// Tokens for use with lexer and parser.

// Tok is a lexical token.
type Tok int

// Token is a token with metadata.
type Token interface {
	Tok() Tok
	StartPos() TokenPos
	EndPos() TokenPos
}

// SymToken is a symbol token.
type SymToken struct {
	t          Tok
	start, end TokenPos
}

// Tok returns the token id.
func (t *SymToken) Tok() Tok {
	return t.t
}

// StartPos returns starting position of token.
func (t *SymToken) StartPos() TokenPos {
	return t.start
}

// EndPos returns ending position of token.
func (t *SymToken) EndPos() TokenPos {
	return t.end
}

// LabelToken is a token with a freeform sting.
type LabelToken struct {
	str        string
	start, end TokenPos
}

// Tok returns IDENT.
func (*LabelToken) Tok() Tok {
	return LABEL
}

// StartPos returns starting position of token.
func (t *LabelToken) StartPos() TokenPos {
	return t.start
}

// EndPos returns ending position of token.
func (t *LabelToken) EndPos() TokenPos {
	return t.end
}

// DigitsToken is a token with numeric value.
type DigitsToken struct {
	num        int
	start, end TokenPos
}

// Tok returns DIGITS.
func (t *DigitsToken) Tok() Tok {
	return DIGITS
}

// StartPos returns starting position of token.
func (t *DigitsToken) StartPos() TokenPos {
	return t.start
}

// EndPos returns ending position of token.
func (t *DigitsToken) EndPos() TokenPos {
	return t.end
}

const (
	// ILLEGAL is a special token for errors.
	ILLEGAL Tok = iota
)

var eof = rune(0)

// TokenPos is a pair of coordinate to identify start of token.
type TokenPos struct {
	Char  int
	Lines []int
}

func (p TokenPos) String() string {
	return fmt.Sprintf("%d:%d", len(p.Lines)+1, p.Char)
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isUnquoted(ch rune) bool {
	return isLetter(ch) || ('0' <= ch && ch <= '9') || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
