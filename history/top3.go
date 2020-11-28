package main

import (
	"fmt"
	"sort"
)

func top3() {
	cmds := map[string]int{}

	err := histWalk(func(stamp int64, line string) error {
		cmds[line]++
		return nil
	})
	if err != nil {
		panic(err)
	}

	type kv struct {
		Key   string
		Value int
	}

	kvs := []kv{}
	for k, v := range cmds {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})

	for i := 0; i < 3; i++ {
		fmt.Printf("%s, (%dx)\n", kvs[i].Key, kvs[i].Value)
	}

}
