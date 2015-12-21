// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"strings"
)

// Advent6Lights scans lines of input and follows instructions
// to toggle lights in a matrix (toggle, on, off) according to specified
// rectangles
func Advent6Lights(s string) (count, b int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += LightsExec(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return
}

// LightGrid represents a grid of lights (1000x1000) and accepts
// instructions for toggling them
type LightGrid struct {
	Grid [1000][1000]bool
}

// NewLightGrid creates a new LightGrid with all lights turned off
func NewLightGrid() (g *LightGrid) {
	g = &LightGrid{}
	return
}

// LightsExec TBD
func LightsExec(s string) int {
	return 0
}
