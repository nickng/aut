package aut

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// Scanner is a lexical scanner.
type Scanner struct {
	r   *bufio.Reader
	pos TokenPos
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), pos: TokenPos{Char: 0, Lines: []int{}}}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if reached the end or error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if ch == '\n' {
		s.pos.Lines = append(s.pos.Lines, s.pos.Char)
		s.pos.Char = 0
	} else {
		s.pos.Char++
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	if s.pos.Char == 0 {
		s.pos.Char = s.pos.Lines[len(s.pos.Lines)-1]
		s.pos.Lines = s.pos.Lines[:len(s.pos.Lines)-1]
	} else {
		s.pos.Char--
	}
}

// Scan returns the next token and parsed value.
func (s *Scanner) Scan() Token {
	var startPos, endPos TokenPos
	ch := s.read()

	if isWhitespace(ch) {
		s.skipWhitespace()
		ch = s.read()
	}
	if isLetter(ch) {
		s.unread()
		return s.scanUnquoted()
	}
	if isDigit(ch) {
		s.unread()
		return s.scanDigit()
	}

	// Track token positions.
	startPos = s.pos
	defer func() { endPos = s.pos }()

	switch ch {
	case eof:
		return &SymToken{t: 0, start: startPos, end: endPos}
	case ',':
		return &SymToken{t: COMMA, start: startPos, end: endPos}
	case '(':
		return &SymToken{t: LPAREN, start: startPos, end: endPos}
	case ')':
		return &SymToken{t: RPAREN, start: startPos, end: endPos}
	case '*': // special case for unquoted.
		return &LabelToken{str: "*", start: startPos, end: endPos}
	case '"': // Quoted string.
		s.unread()
		return s.scanQuoted()
	}
	return &SymToken{t: ILLEGAL, start: startPos, end: endPos}
}

func (s *Scanner) scanUnquoted() Token {
	var startPos, endPos TokenPos
	var buf bytes.Buffer

	startPos = s.pos
	defer func() { endPos = s.pos }()

	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isUnquoted(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	if buf.String() == "des" {
		return &SymToken{t: DES, start: startPos, end: endPos}
	}
	return &LabelToken{str: buf.String(), start: startPos, end: endPos}
}

func (s *Scanner) scanDigit() Token {
	var startPos, endPos TokenPos
	var buf bytes.Buffer

	startPos = s.pos
	defer func() { endPos = s.pos }()

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	if i, err := strconv.Atoi(buf.String()); err == nil {
		return &DigitsToken{num: i, start: startPos, end: endPos}
	}
	return &SymToken{t: ILLEGAL, start: startPos, end: endPos}
}

func (s *Scanner) scanQuoted() Token {
	var startPos, endPos TokenPos
	var buf bytes.Buffer

	startPos = s.pos
	defer func() { endPos = s.pos }()

	buf.WriteRune(s.read())

QUOTESEARCH:
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '"' {
			var searchBuf bytes.Buffer
			var nextRune rune
			searchBuf.WriteRune(ch)
			for nextRune = s.read(); isWhitespace(nextRune); nextRune = s.read() {
				searchBuf.WriteRune(nextRune)
			}
			if nextRune == ',' {
				// BUG(nickng): Heuristic to detect end of quoted string, could give wrong end-of-quote if string includes '" ,'
				s.unread()
				buf.WriteRune('"') // Put a final quote back in.
				break QUOTESEARCH
			} else {
				s.unread()
				buf.WriteString(searchBuf.String())
			}
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return &LabelToken{str: buf.String(), start: startPos, end: endPos}
}

func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
