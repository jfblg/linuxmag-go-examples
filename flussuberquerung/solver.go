package main

import (
	"errors"
	"fmt"
)

func solve(state State) ([]State, error) {
	seen := make(map[string]bool)
	todo := []State{state}

	for len(todo) > 0 {
		// pop off last element
		lastidx := len(todo) - 1
		s := todo[lastidx]
		todo = todo[:lastidx]

		// prevent cycles
		if _, ok := seen[s.String()]; ok {
			continue
		}

		seen[s.String()] = true

		if s.IsFinished() {
			path := []State{}
			for cs := &s; cs.prev != nil; cs = cs.prev {
				c := cs.Copy()
				path = append([]State{c}, path...)
			}
			path = append([]State{state}, path...)
			return path, nil
		}
		for _, succ := range s.Successors() {
			// insert new element at end
			succ.prev = &s
			todo = append([]State{succ}, todo...)
		}
	}
	return []State{}, errors.New("Can't solve")
}

func main() {
	var state State

	state.west.Add(Wolf)
	state.west.Add(Cabbage)
	state.west.Add(Ferryman)
	state.west.Add(Goat)

	solution, err := solve(state)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution: \n\n")
	for _, step := range solution {
		fmt.Printf("[%s]\n", step)
	}
}
