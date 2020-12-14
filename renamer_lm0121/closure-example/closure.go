package main

import "fmt"

// run "go run closure.go"

func main() {
	mycounter := mkmycounter()

	mycounter()
	mycounter()
	mycounter()
	mycounter()
}

func mkmycounter() func() {
	count := 1

	return func() {
		fmt.Printf("%d\n", count)
		count++
	}
}
