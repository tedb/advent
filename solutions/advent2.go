// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
	"unicode"
	//"fmt"
)

// Advent2a_Box takes a list of lines like 1x2x3 and returns the sum of
// Box.Sqft
func Advent2_Box(s string) (sum_sqft, sum_ribbon int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		b := NewBox(scanner.Text())
		sum_sqft += b.Sqft()
		sum_ribbon += b.Ribbon()
	}

	err := scanner.Err()
	check_err(err)

	return
}

type Box struct {
	x, y, z int
}

// NewBox takes strings like 1x2x3 (XxYxZ)
// and returns a Box initialized with those values
func NewBox(s string) (b *Box) {
	b = &Box{}
	splitter := func(c rune) bool {
		return !unicode.IsNumber(c)
	}
	parts := strings.FieldsFunc(s, splitter)

	var err error
	b.x, err = strconv.Atoi(parts[0])
	check_err(err)
	b.y, err = strconv.Atoi(parts[1])
	check_err(err)
	b.z, err = strconv.Atoi(parts[2])
	check_err(err)

	return b
}

// Sqft returns the square footage to cover a box of the given dimensions,
// plus the surface area of the smallest side
func (b *Box) Sqft() int {
	face1, face2, face3 := b.x*b.y, b.y*b.z, b.z*b.x
	return (2 * face1) + (2 * face2) + (2 * face3) + min3(face1, face2, face3)
}

func (b *Box) Ribbon() int {
	s := b.sorted_sides()
	return 2*s[0] + 2*s[1] + b.x*b.y*b.z
}

func (b *Box) sorted_sides() (s []int) {
	s = []int{b.x, b.y, b.z}
	sort.Ints(s)
	return
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}
