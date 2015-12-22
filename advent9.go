package advent

import (
	"bufio"
	"fmt"
	"strings"
)

// Advent09TSP accepts lines stating distances from one city to another,
// then returns the length of the shortest route among them (visiting each once)
func Advent09TSP(s string) (shortest, longest int) {
	c := make(Cities, 0)

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		c.AddFromString(line)
	}
	err := scanner.Err()
	checkErr(err)

	shortest = c.ToCityRoute().BestRoute().TotalDistance()
	longest = c.ToCityRoute().WorstRoute().TotalDistance()

	return
}

// City specifies the distances to a given city
type City struct {
	Name string
	Km   map[string]int
}

// Cities contains distances from one city to another
type Cities map[string]City

// AddFromString adds to Cities by parsing a string of the form:
//     London to Dublin = 464
func (c Cities) AddFromString(s string) {
	var src, dest string
	var km int

	_, err := fmt.Sscanf(s, "%s to %s = %d", &src, &dest, &km)
	checkErr(err)

	_, ok := c[src]
	if !ok {
		c[src] = City{Name: src, Km: map[string]int{dest: km}}
	} else {
		c[src].Km[dest] = km
	}

	_, ok = c[dest]
	if !ok {
		c[dest] = City{Name: dest, Km: map[string]int{src: km}}
	} else {
		c[dest].Km[src] = km
	}

}

// ToCityRoute returns a slice of the City values in Cities
func (c Cities) ToCityRoute() (r CityRoute) {
	r = make(CityRoute, 0)
	for _, v := range c {
		r = append(r, v)
	}
	return r
}

// BestRoute permutes CityMap orders and returns the Route with the shortest distance.
func (r CityRoute) BestRoute() (best CityRoute) {
	min := 100000000

	for _, p := range Permutations(Seq(0, len(r)-1), 0) {
		re := r.Reorder(p)
		d := re.TotalDistance()
		//fmt.Printf("perm %v: %v, %d\n", p, re, d)

		if d == 0 {
			continue
		}
		if d < min {
			min = d
			best = re
		}
	}
	return best
}

// WorstRoute permutes CityMap orders and returns the Route with the longest distance.
func (r CityRoute) WorstRoute() (best CityRoute) {
	max := 0

	for _, p := range Permutations(Seq(0, len(r)-1), 0) {
		re := r.Reorder(p)
		d := re.TotalDistance()
		//fmt.Printf("perm %v: %v, %d\n", p, re, d)

		if d == 0 {
			continue
		}
		if d > max {
			max = d
			best = re
		}
	}
	return best
}

type CityRoute []City

// Reorder arranges Route in order of specified indexes
func (r CityRoute) Reorder(o []int) (ro CityRoute) {
	ro = make(CityRoute, 0)
	for _, i := range o {
		ro = append(ro, r[i])
	}
	return ro
}

// TotalDistance sums the distances from each city to the next
func (r CityRoute) TotalDistance() (sum int) {
	for i := 0; i < len(r)-1; i++ {
		nextName := r[i+1].Name
		dist := r[i].Km[nextName]
		// println("distance:", r[i].Name, nextName, dist)

		// the next city isn't one we can reach
		if dist == 0 {
			return 0
		}
		sum += dist
	}

	return sum
}
