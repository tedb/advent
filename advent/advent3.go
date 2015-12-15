// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
)

// Advent3_Houses takes a list of ^>v< and returns the number of positions visited
func Advent3_Houses(s string) (sum1, sum2 int) {
	return NewRoute().Nav(s).How_many_unique(), NewRoute().DualNav(s).How_many_unique()
}

type RoutePos struct {
	x, y int
}
type Route struct {
	visited map[RoutePos]struct{}
	pos_x1, pos_y1, pos_x2, pos_y2 int
}

// NewRoute sets up our x/y movement plane and returns a Route
func NewRoute() (r *Route) {
	r = &Route{}
	r.visited = make(map[RoutePos]struct{})
	// don't move any, just get the side effect of recording the origin visit
	r.MoveSanta(0, 0)
	return r
}

// Move Santa around according to ^>v< instructions 
func (r *Route) Nav(s string) (*Route) {
	for _, m := range s {
		switch m {
			case '^':
			r.MoveSanta(0, 1)
			case '>':
			r.MoveSanta(1, 0)
			case 'v':
			r.MoveSanta(0, -1)
			case '<':
			r.MoveSanta(-1, 0)
		}
	}
	return r
}

// Move Santa and Robo-Santa around according to alternating ^>v< instructions
func (r *Route) DualNav(s string) (*Route) {
	for i, m := range s {
		f := r.MoveSanta
		if i % 2 != 0 {
			f = r.MoveRoboSanta
		}
		switch m {
			case '^':
			f(0, 1)
			case '>':
			f(1, 0)
			case 'v':
			f(0, -1)
			case '<':
			f(-1, 0)
		}
	}
	return r
}

// move changes position in a delta of x/y direction and records the visit
func (r *Route) MoveSanta(d_x, d_y int) {
	r.pos_x1 += d_x
	r.pos_y1 += d_y
	r.visited[RoutePos{r.pos_x1, r.pos_y1}] = struct{}{}
}

// move changes position in a delta of x/y direction and records the visit
func (r *Route) MoveRoboSanta(d_x, d_y int) {
	r.pos_x2 += d_x
	r.pos_y2 += d_y
	r.visited[RoutePos{r.pos_x2, r.pos_y2}] = struct{}{}
}

func (r *Route) How_many_unique() int {
	return len(r.visited)
}

