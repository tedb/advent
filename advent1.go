// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

// Advent01aParens returns the sum of a sequence of +/- in the form of ( and )
func Advent01aParens(s string) (sum int) {
	for _, x := range s {
		if x == '(' {
			sum++
		} else if x == ')' {
			sum--
		}
	}
	return
}

// Advent01bParensBasement follows a sequence of +/- in the form of ( and ),
// and returns the position in the sequence when the accumulated value is -1
func Advent01bParensBasement(s string) (pos int) {
	var sum int
	for i, x := range s {
		if x == '(' {
			sum++
		} else if x == ')' {
			sum--
		}
		if sum == -1 {
			return i + 1
		}
	}
	return
}
