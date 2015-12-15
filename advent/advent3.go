// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
)

// Advent3_Houses takes a list of ^>v< and returns the number of positions visited
func Advent3_Houses(s string) (sum1, sum2 int) {
	r := NewRoute().Nav(s)
	return r.How_many_unique(), 0
}

type RoutePos struct {
	x, y int
}
type Route struct {
	visited map[RoutePos]int
	pos_x, pos_y int
}

// NewRoute sets up our x/y movement plane and returns a Route
func NewRoute() (r *Route) {
	r = &Route{}
	r.visited = make(map[RoutePos]int)
	// don't move any, just to get the side effect of recording the origin visit
	r.Move(0, 0)
	return r
}

func (r *Route) Nav(s string) (*Route) {
	for _, m := range s {
		switch m {
			case '^':
			r.Move(0, 1)
			case '>':
			r.Move(1, 0)
			case 'v':
			r.Move(0, -1)
			case '<':
			r.Move(-1, 0)
		}
	}
	return r
}

// move changes position in a delta of x/y direction and records the visit
func (r *Route) Move(d_x, d_y int) {
	r.pos_x += d_x
	r.pos_y += d_y
	r.visited[RoutePos{r.pos_x, r.pos_y}]++
}

func (r *Route) How_many_unique() int {
	return len(r.visited)
}

