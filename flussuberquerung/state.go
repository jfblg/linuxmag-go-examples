package main

import (
	"fmt"
)

// State repesents a current state on both sides of a shore
type State struct {
	west Shore
	east Shore
	prev *State
}

func (s State) IsValid() bool {
	return s.west.IsValid() && s.east.IsValid()
}

func (s State) String() string {
	return fmt.Sprintf("%s | %s",
		s.west.String(),
		s.east.String(),
	)
}

func (s *State) Move(p Passenger) {
	from := &s.west
	to := &s.east
	if to.Has(Ferryman) {
		from, to = to, from
	}
	from.Move(p, to)
}

func (s State) IsFinished() bool {
	return len(s.west.passengers) == 0
}

func (s State) Copy() State {
	d := State{}
	d.west = s.west.Copy()
	d.east = s.east.Copy()
	return d
}

func (s State) Successors() []State {
	startShore := s.east
	if s.west.Has(Ferryman) {
		startShore = s.west
	}

	results := []State{}

	for _, passenger := range startShore.passengers {
		candidate := s.Copy()
		candidate.Move(passenger)
		if passenger != Ferryman {
			candidate.Move(Ferryman)
		}
		if candidate.IsValid() {
			results = append(results, candidate)
		}
	}
	return results
}
