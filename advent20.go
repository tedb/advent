package advent

import (
	"github.com/ntns/goitertools/itertools"
	"github.com/otiai10/primes"
	"sort"
	"strconv"
)

// Advent20InfiniteElves determines the lowest house number of the house
// to get at least as many presents as the number in the puzzle input
func Advent20InfiniteElves(presents_s string) (house, b int) {
	_, err := strconv.Atoi(presents_s)
	checkErr(err)
	/*
		// i is both house number and elf number
		for i := 1; ; i++ {
			p := PresentsForHouse(i)
				if i % elf == 0 {
					houses[i] += elf * 10
				}
				if houses[i] >= presents {
					//println(houses[i], i, presents)
					return i, 0
				}
			}
		}*/
	return

}

// PresentsForHouse returns "sum of divisors" for the given house number, times ten.
// Divisors (or factors) are calculated by determining h's unique prime factors,
// then appending 1 and h to the list, rendering all pairwise combinations
// of the unique factors, then summing those products.  "Sum of divisors"
// is also [OEIS sequence A000203](http://oeis.org/A000203).  A simpler way to solve this
// would be just looping through 1..h and checking divisibility, but using more math makes it fun.
func PresentsForHouse(h int) (p int) {
	primeFactors64 := primes.Factorize(int64(h)).All()
	// convert to []int
	primeFactors := make([]int, 0)
	for _, pF := range primeFactors64 {
		primeFactors = append(primeFactors, int(pF))
	}

	//fmt.Printf("prime factors %d: %v\n", h, primeFactors)
	primeCombos := [][]int{}
	// FIXME: is there an axiom for the max number of factors a number can have?
	for i := 1; i <= len(primeFactors); i++ {
		c := itertools.Combinations(primeFactors, i)
		//fmt.Printf("appending combo %v: %v\n", h, c)

		primeCombos = append(primeCombos, c...)
	}

	//fmt.Printf("primeCombos %d: %v\n", h, primeCombos)

	allFactors := []int{1}

	for _, primeCombo := range primeCombos {
		allFactors = append(allFactors, MultiplyInts(primeCombo))
	}
	allFactors = UniqueInts(allFactors)
	//fmt.Printf("all factors %d: %v\n", h, allFactors)

	// Presents is the sum of all factors
	p = SumInts(allFactors)

	// Puzzle wants "sum of divisors" * 10
	p *= 10
	return
}

func SumInts(i []int) (s int) {
	for _, x := range i {
		s += x
	}
	return
}

func MultiplyInts(i []int) (s int) {
	if len(i) == 0 {
		return 0
	}
	s = 1
	for _, x := range i {
		s *= x
	}
	return
}

// UniqueInts returns the unique elements of []int
func UniqueInts(in []int) (out []int) {
	if len(in) == 0 {
		return
	}
	sort.Ints(in)
	out = []int{in[0]}

	for i := 1; i < len(in); i++ {
		if in[i] != in[i-1] {
			out = append(out, in[i])
		}
	}
	return
}
