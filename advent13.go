package advent

import (
	"bufio"
	"fmt"
	"strings"
)

// Advent13Seating accepts lines stating each HappyPerson's happiness delta
// resulting from sitting next to a given HappyPerson, then permutes a
// SeatingArrangement to maximize the total happiness
func Advent13Seating(s string) (dHappiness1, dHappiness2 int) {
	m := make(HappyPersonMap)

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		m.AddFromString(line)
	}
	err := scanner.Err()
	checkErr(err)

	// for k, v := range m {
	// 	fmt.Printf("%s: %v, ", k, v)
	// }
	// fmt.Print("\n")

	dHappiness1 = m.BestSeatingArrangement().TotalHappiness()
	m.AddMeToo()
	dHappiness2 = m.BestSeatingArrangement().TotalHappiness()
	return
}

// HappyPersonMap contains all the people sitting at the table, no particular order
type HappyPersonMap map[string]*HappyPerson

// AddFromString adds a HappyPerson (or appends more Factors for an existing HappyPerson)
// to HappyPersonMap by parsing a string of the form:
//     Alice would gain 54 happiness units by sitting next to Bob.
func (m HappyPersonMap) AddFromString(s string) {
	var name, gainLose, adj string
	var happiness int

	_, err := fmt.Sscanf(s, "%s would %s %d happiness units by sitting next to %s", &name, &gainLose, &happiness, &adj)
	checkErr(err)

	adj = strings.TrimRight(adj, ".")
	if gainLose == "lose" {
		happiness *= -1
	}
	//println(name, gainLose, happiness, adj)

	if m[name] == nil {
		m[name] = &HappyPerson{Name: name, Factors: map[string]int{adj: happiness}}
	} else {
		m[name].Factors[adj] = happiness
	}
}

// ToSeatingArrangement returns a slice of the HappyPerson values in HappyPersonMap
func (m HappyPersonMap) ToSeatingArrangement() (s SeatingArrangement) {
	s = make(SeatingArrangement, 0)
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// BestSeatingArrangement permutes SeatingArrangement orders based on the contents
// of HappyPersonMap, and returns the SeatingArrangement with the maximum
// TotalHappiness
func (m HappyPersonMap) BestSeatingArrangement() (best SeatingArrangement) {
	base := m.ToSeatingArrangement()
	max := 0

	for _, p := range Permutations(Seq(0, len(base)-1), 0) {
		a := base.Reorder(p)
		h := a.TotalHappiness()
		//fmt.Printf("perm %v: %v, %d\n", p, a, h)

		if h > max {
			max = h
			best = a
		}
	}

	return best
}

// AddMeToo addresses part 2, where the host is added to the seating arrangements
// with an effective delta-Happiness of 0 in both directions.  This should mean
// we can just set an empty Factors map for this person.
func (m HappyPersonMap) AddMeToo() {
	m["Me"] = &HappyPerson{Name: "Me", Factors: map[string]int{}}
}

// HappyPerson is one person within HappyPersonMap, sitting within SeatingArrangement
type HappyPerson struct {
	Name string
	// Factors records the happiness delta resulting from being seated next to the given HappyPerson
	Factors map[string]int
}

// SeatingArrangement specifies the ring order of HappyPerson seatings
type SeatingArrangement []*HappyPerson

// Reorder arranges SeatingArrangement in order of specified indexes
func (a SeatingArrangement) Reorder(o []int) (r SeatingArrangement) {
	r = make(SeatingArrangement, 0)
	for _, i := range o {
		r = append(r, a[i])
	}
	return r
}

// TotalHappiness sums the HappyPerson factors around a circular table, in both directions
// for each person in a given SeatingArrangement
func (a SeatingArrangement) TotalHappiness() (sum int) {
	for i, me := range a {
		var l, r *HappyPerson
		// first HappyPerson's left person is at end of slice
		if i == 0 {
			l = a[len(a)-1]
		} else {
			l = a[i-1]
		}

		// last HappyPerson's right person is at start of slice
		if i == len(a)-1 {
			r = a[0]
		} else {
			r = a[i+1]
		}

		sum += me.Factors[l.Name] + me.Factors[r.Name]
	}

	return sum
}

// Permutations returns a slice of all the permutations of the input slice.
// Inspired by Python code at 
// [Stack Overflow](http://stackoverflow.com/questions/2710713/algorithm-to-generate-all-possible-permutations-of-a-list)
// and heavily rewritten.
func Permutations(in []int, low int) (r [][]int) {
	xs := make([]int, len(in))
	copy(xs, in)
	r = make([][]int, 0)
	if low+1 >= len(xs) {
		r = append(r, xs)
		return
	}

	r = append(r, Permutations(xs, low+1)...)
	for i := low + 1; i < len(xs); i++ {
		xs[low], xs[i] = xs[i], xs[low]
		r = append(r, Permutations(xs, low+1)...)
		xs[low], xs[i] = xs[i], xs[low]
	}
	return r
}

// Seq returns digits m..n, for range
func Seq(m, n int) (r []int) {
	length := n - m + 1
	r = make([]int, length)
	for i := 0; i < length; i++ {
		r[i] = m + i
	}
	return r
}
