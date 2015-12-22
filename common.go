package advent

import (
//"log"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}
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
