%{
package aut

import (
	"io"
)

var autfile *Aut
%}

%union {
	str   string
	num   int
	trans *Transition
}

%token DES LPAREN RPAREN COMMA DQUOTE
%token <num> DIGITS
%token <str> LABEL
%type <trans> trans

%%

aut : des trans
    ;

des : DES LPAREN DIGITS COMMA DIGITS COMMA DIGITS RPAREN { autfile.SetDes($3, $5, $7) }
    ;

trans : trans LPAREN DIGITS COMMA LABEL COMMA DIGITS RPAREN { autfile.AddTransition(State($3), $5, State($7)) }
      |       LPAREN DIGITS COMMA LABEL COMMA DIGITS RPAREN { autfile.AddTransition(State($2), $4, State($6)) }
      ;

%%

func Parse(r io.Reader) (*Aut, error) {
	autfile = new(Aut)
	l := NewLexer(r)
	autParse(l)
	select {
	case err := <-l.Errors:
		return nil, err
	default:
		return autfile, nil
	}
}
