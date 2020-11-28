package main

import (
	"bufio"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

func histWalk(cb func(int64, string) error) error {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	home := usr.HomeDir
	histfile := filepath.Join(home, ".bash_history")

	f, err := os.Open(histfile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var timestamp int64
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#' {
			timestamp, err = strconv.ParseInt(line[1:], 10, 64)
			if err != nil {
				panic(err)
			}
		} else {
			err := cb(timestamp, line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
