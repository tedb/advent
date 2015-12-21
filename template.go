package advent

import (
	"bufio"
	"strings"
)

// Advent99Template is a template function for other Advent of Code puzzles to start from
func Advent99Template(s string) (count, b int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += Foo(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return
}

// Foo is an example and does nothing
func Foo(s string) int {
	return 0
}
