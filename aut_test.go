package aut_test

import (
	"strings"
	"testing"

	"github.com/nickng/aut"
)

// exampleAut is the example valid AUT file from the AUT manual page.
const exampleAut = `
   des (0, 5, 4)
   (0, PUT_6, 1)
   (0, "GET !true !7 !CONS (A, CONS (B, NIL))", 2)
   (1, i, 3)
   (2, "SEND !"hello" !"world"", 3)
   (3, "DUMP ... " ... \" ...", 0)
`

// emptyAut is a minimal AUT file that has one self-looping state.
const minimalAut = `
des (0, 1, 1)
(0, Loop, 0)
`

// This tests some of the more tracking parsing of label.
func TestExampleAut(t *testing.T) {
	r := strings.NewReader(exampleAut)
	parsed, err := aut.Parse(r)
	if err != nil {
		t.Error("Parse failed:", err)
	}
	if parsed.Init != 0 {
		t.Error("Init state: expects 0, got", parsed.Init)
	}
	if parsed.NumTransitions != 5 {
		t.Error("Number of Transitions: expects 5, got", parsed.NumTransitions)
	}
	if parsed.NumStates != 4 {
		t.Error("Number of States: expects 4, got", parsed.NumTransitions)
	}
	if len(parsed.Transitions) != 5 {
		t.Error("Transition records: expects 5, got", len(parsed.Transitions))
	}
	// With quote.
	if tr := parsed.Transitions[3]; !(tr.From == 2 && tr.To == 3 && tr.Label == `"SEND !"hello" !"world""`) {
		t.Error(`Transition3: expects (2, "SEND !"hello" !"world"", 3), got`, parsed.Transitions[3])
	}
	// With tricky quote.
	if tr := parsed.Transitions[4]; !(tr.From == 3 && tr.To == 0 && tr.Label == `"DUMP ... " ... \" ..."`) {
		t.Error(`Transition3: expects (3, "DUMP ... " ... \" ...", 0), got`, parsed.Transitions[4])
	}
}

// This tests the correct construction of a aut data structure.
func TestMinimalAut(t *testing.T) {
	r := strings.NewReader(minimalAut)
	parsed, err := aut.Parse(r)
	if err != nil {
		t.Error("Parse failed:", err)
	}
	if parsed.Init != 0 {
		t.Error("Init state: expects 0, got", parsed.Init)
	}
	if parsed.NumTransitions != 1 {
		t.Error("Number of Transitions: expects 1, got", parsed.NumTransitions)
	}
	if parsed.NumStates != 1 {
		t.Error("Number of States: expects 1, got", parsed.NumTransitions)
	}
	if len(parsed.Transitions) != 1 {
		t.Error("Transition records: expects 1, got", len(parsed.Transitions))
	}
	// No quote.
	if tr := parsed.Transitions[0]; !(tr.From == 0 && tr.To == 0 && tr.Label == `Loop`) {
		t.Error("Transition0: expects (0, Loop, 0), got", parsed.Transitions[0])
	}
}
