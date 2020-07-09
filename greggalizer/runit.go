package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func runit(argv []string) string {
	out, err := exec.Command(argv[0], argv[1:]...).Output()

	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	r := regexp.MustCompile("\\t")
	return r.ReplaceAllString(string(out), " ")
}
