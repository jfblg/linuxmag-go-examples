package main

import (
	"fmt"
)

func main() {
	items := []string{"zelena", "modra", "ruzova", "zlta", "biela"}

	for i := range items {
		for j := i + 1; j < len(items); j++ {
			if items[i] > items[j] {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	fmt.Printf("%+v\n", items)
}
