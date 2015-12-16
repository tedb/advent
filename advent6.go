// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
)

// Advent6_Lights does...
func Advent6_Lights(s string) (a, b int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += LightsExec(scanner.Text())
	}

	err := scanner.Err()
	check_err(err)

	return
}

type LightGrid struct {
	Grid [1000][1000]bool
}

func NewLightGrid() (g *LightGrid) 
	g = &LightGrid{}
	return
}

func LightsExec(s string) int {
	return 0
}
