package main

import (
	"fmt"
	"sort"
)

// Shore represents a side of a river. It will be either West or East side.
type Shore struct {
	passengers []Passenger
}

func (s Shore) String() string {
	sort.Sort(Passengers(s.passengers))
	return fmt.Sprintf("%s", s.passengers)
}

func (s Shore) Copy() Shore {
	c := Shore{}
	c.passengers = make([]Passenger, len(s.passengers))
	copy(c.passengers, s.passengers)
	return c
}

func (s *Shore) Add(p Passenger) {
	s.passengers = append(s.passengers, p)
}

func (s *Shore) Remove(p Passenger) {
	result := []Passenger{}
	for _, passenger := range s.passengers {
		if passenger != p {
			result = append(result, passenger)
		}
	}
	s.passengers = result
}

func (s *Shore) Move(p Passenger, t *Shore) {
	s.Remove(p)
	t.Add(p)
}

func (s *Shore) Has(p Passenger) bool {
	for _, passenger := range s.passengers {
		if passenger == p {
			return true
		}
	}
	return false
}

func (s Shore) IsValid() bool {
	if s.Has(Ferryman) {
		return true
	}
	if s.Has(Wolf) && s.Has(Goat) {
		return false
	}
	if s.Has(Goat) && s.Has(Cabbage) {
		return false
	}
	return true
}
