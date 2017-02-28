package aut

import (
	"bytes"
	"fmt"
)

// State is a representation of a state.
type State int

// Transition is a representation of transition between two states.
type Transition struct {
	From  State
	Label string
	To    State
}

// Aut is the root data structure of an aut file.
type Aut struct {
	Init           State
	NumTransitions int
	NumStates      int
	Transitions    []*Transition
}

// SetDes sets the DEScription of an aut representation.
func (a *Aut) SetDes(start, ntrans, nstate int) {
	a.Init = State(start)
	a.NumTransitions = ntrans
	a.NumStates = nstate
}

// AddTransition adds a single transition to the current aut representation.
func (a *Aut) AddTransition(from State, label string, to State) {
	a.Transitions = append(a.Transitions, &Transition{From: from, Label: label, To: to})
}

// String returns a string representation of the aut representation in aut
// format.
func (a *Aut) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("des (%d, %d, %d)\n",
		int(a.Init), a.NumTransitions, a.NumStates))
	for _, t := range a.Transitions {
		buf.WriteString(fmt.Sprintf("(%d, %s, %d)\n",
			int(t.From), t.Label, int(t.To)))
	}

	return buf.String()
}
