package advent

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
	"unicode"
	//"fmt"
)

// Advent02Box takes a list of lines like 1x2x3 and returns the sums from
// Box.Sqft and Box.Ribbon
func Advent02Box(s string) (sumSqft, sumRibbon int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		b := NewBox(scanner.Text())
		sumSqft += b.Sqft()
		sumRibbon += b.Ribbon()
	}

	err := scanner.Err()
	checkErr(err)

	return
}

// Box provides methods Sqft and Ribbon to calculate
// the amount of paper and ribbon needed for them
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
	checkErr(err)
	b.y, err = strconv.Atoi(parts[1])
	checkErr(err)
	b.z, err = strconv.Atoi(parts[2])
	checkErr(err)

	return b
}

// Sqft returns the square footage to cover a box of the given dimensions,
// plus the surface area of the smallest side
func (b *Box) Sqft() int {
	face1, face2, face3 := b.x*b.y, b.y*b.z, b.z*b.x
	return (2 * face1) + (2 * face2) + (2 * face3) + Min3(face1, face2, face3)
}

// Ribbon returns the linear length of ribbon needed to cover the given box
func (b *Box) Ribbon() int {
	s := b.sortedSides()
	return 2*s[0] + 2*s[1] + b.x*b.y*b.z
}

// sortedSides returns the box's dimensions in order
func (b *Box) sortedSides() (s []int) {
	s = []int{b.x, b.y, b.z}
	sort.Ints(s)
	return
}

// Min3 returns the minimum value given a list of 3 ints
func Min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
