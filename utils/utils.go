package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileLines(f string) []string {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	s := string(b)

	lines := strings.Split(strings.TrimSpace(s), "\n")

	return lines
}

func ReadFileLinesInt(f string) []int {
	l := ReadFileLines(f)

	var ints []int

	for _, s := range l {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		ints = append(ints, i)
	}

	return ints
}
