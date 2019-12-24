package main

// Passengers is a slice of a Passenger type
type Passengers []Passenger

func (p Passengers) Len() int {
	return len(p)
}

func (p Passengers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Passengers) Less(i, j int) bool {
	return p[i] < p[j]
}
