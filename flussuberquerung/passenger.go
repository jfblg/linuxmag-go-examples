package main

// Passenger represents a kind of object
type Passenger int

const (
	Goat Passenger = iota
	Wolf
	Cabbage
	Ferryman
)

func (p Passenger) String() string {
	return []string{"Goat", "Wolf", "Cabbage", "Ferryman"}[p]
}
