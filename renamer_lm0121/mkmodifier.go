package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func mkmodifier(cmd string) (func(string) string, error) {
	parts := strings.Split(cmd, "/")
	if len(parts) != 2 {
		return nil, errors.New("Invalid repl command")
	}

	search := parts[0]
	repltmpl := parts[1]
	seq := 1

	var rex *regexp.Regexp

	if len(search) == 0 {
		search = ".*"
	}

	rex = regexp.MustCompile(search)

	modifier := func(org string) string {
		repl := strings.Replace(repltmpl, "{seq}", fmt.Sprintf("%04d", seq), -1)
		seq++
		res := rex.ReplaceAllString(org, repl)
		return string(res)

	}
	return modifier, nil
}
