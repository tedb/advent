package advent

import (
	"bufio"
	"strings"
)

// Advent11Password (NOT DONE) is a template function for other Advent of Code puzzles to start from
func Advent11Password(s string) (count, b int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += Foo(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return
}
