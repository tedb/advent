// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"strings"
)

// Advent19Replacements is a template function for other Advent of Code puzzles to start from
func Advent19Replacements(s string) (count, b int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += Foo(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return
}
