// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"strings"
)

// Advent0Foo is a template function for other Advent of Code puzzles to start from
func Advent0Foo(s string) (count, b int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += Foo(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return
}

// Foo is a template function to process one line from AdventXFoo
func Foo(s string) int {
	return 0
}
