package main

import (
	"fmt"
	"time"
)

func dow() {
	var countByDow [7]int

	err := histWalk(func(stamp int64, line string) error {
		dt := time.Unix(stamp, 0)
		countByDow[int(dt.Weekday())]++
		return nil
	})
	if err != nil {
		panic(err)
	}

	for dow := 0; dow < len(countByDow); dow++ {
		dowStr := time.Weekday(dow).String()
		fmt.Printf("%s: %v\n", dowStr, countByDow[dow])
	}
}
