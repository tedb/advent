package advent

import (
	"fmt"
	"github.com/ntns/goitertools/itertools"
	"github.com/otiai10/primes"
	"runtime"
	"sort"
	"strconv"
	"sync"
)

// Advent20InfiniteElves determines the lowest house number of the house
// to get at least as many presents as the number in the puzzle input.
// Tried using Sum of Divisors algorithm, and was successful, but it
// was super slow.  Trying brute force algorithm.  FUTURE: Retry
// with a different implementation of a prime factor based algorithm.  Calculate
// computational complexity of each approach.
func Advent20InfiniteElves(presentsStr string) (house1 int) {
	presents, err := strconv.Atoi(presentsStr)
	checkErr(err)

	houses1 := make([]int, presents/10)

	for elf := 1; ; elf++ {
		if elf%100 == 0 {
			println("done with elf", elf)
		}
		for houseIdx := elf; houseIdx < len(houses1); houseIdx += elf {
			houses1[houseIdx] += elf * 10
			//fmt.Println("delivering for elf", elf, "to house", houseIdx, "=", houses[houseIdx])
		}

		for i, v := range houses1 {
			if v >= presents && elf >= i {
				//fmt.Printf("winning houses(%d): i=%d, v=%d, elf=%d %v\n", presents, i, v, elf, houses[1:])
				return i
			}
		}
	}
	//return
}

// Advent20bInfiniteElves is basically the same as Advent20InfiniteElves,
// but we use a parallel algorithm and the elves each go to only 50 houses
// instead of infinite.
func Advent20bInfiniteElves(presentsStr string) (house2 int) {
	presents, err := strconv.Atoi(presentsStr)
	checkErr(err)

	houses2 := make([]int, 1000000)

	var wg sync.WaitGroup
	elves := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(elfC chan int) {
			defer wg.Done()
			for elf := range elfC {
				if elf%1000 == 0 {
					println("done with elf", elf)
				}

				for houseIdx := elf; houseIdx <= elf*50 && houseIdx < len(houses2); houseIdx += elf {
					houses2[houseIdx] += elf * 11
				}
			}
		}(elves)
	}

	for elf := 1; elf <= len(houses2); elf++ {
		elves <- elf
	}
	close(elves)
	wg.Wait()
	for i, v := range houses2 {
		if v >= presents {
			fmt.Printf("winning houses(%d): i=%d, v=%d\n", presents, i, v)
			return i
		}
	}

	return
}

// Advent20InfiniteElvesFactors determines the lowest house number of the house
// to get at least as many presents as the number in the puzzle input.
// This version uses Sum of Divisors algorithm.
func Advent20InfiniteElvesFactors(presentsStr string) (house, b int) {
	presents, err := strconv.Atoi(presentsStr)
	checkErr(err)

	var runningSum int
	var runningCount int
	var stopJumping bool

	// i is both house number and elf number.
	// This is pretty slow (about 1000 houses per second),
	// so to narrrow the search space, we start counting at 20% short of
	// the input value.
	// Could we instead find a prime number > presents,
	// then start counting from the previous prime number?
	for i := 1; ; i++ {
		pH := PresentsForHouse(i)

		runningSum += pH
		runningCount++
		runningAvg := runningSum / runningCount

		if i%1000 == 0 {
			println("so far:", i, pH, runningAvg)
			runningSum = 0
			runningCount = 0
		}

		// we've got a winner!
		if pH >= presents {
			return i, 0
		}

		if !stopJumping && runningAvg > int(float64(presents)*0.6) {
			stopJumping = true
			println("stopping jumping at presents=", pH)
		}
		// fast forward the loop if we're not close
		if i > 100 && !stopJumping {
			//if i > 100 && pH < int(float64(presents) * 0.3) {
			//i = int(float64(i) * 1.05)
			i += 1000
			runningSum = 0
			runningCount = 0

			println("jumping i to", i, "so far:", pH, runningAvg)
		}
	}
	//return
}

// Advent20InfiniteElvesSlow ended up being a very slow algorithm
// due to the underlying library's repeated calculation of primes.
// Would be great if the otiai10 library memoized the prime list.
func Advent20InfiniteElvesSlow(presentsStr string) (house, b int) {
	presents, err := strconv.Atoi(presentsStr)
	checkErr(err)

	for i := 1; i < 100000000; i++ {
		pH := PresentsForHouseSlowWithFactorization(i)
		/*if i % 100 == 0 {
			println("so far, slow:", i, pH)
		}*/
		if pH >= presents {
			return i, 0
		}
	}
	return
}

// PresentsForHouseSlowWithFactorization returns
// "sum of divisors" for the given house number, times ten.
// Divisors (or factors) are calculated by determining h's unique prime factors,
// then appending 1 and h to the list, rendering all pairwise combinations
// of the unique factors, then summing those products.  "Sum of divisors"
// is also [OEIS sequence A000203](http://oeis.org/A000203).  A simpler way to solve this
// would be just looping through 1..h and checking divisibility, but using more math makes it fun.
// Try alternate factoring algorithm from https://en.wikipedia.org/wiki/Trial_division .
func PresentsForHouseSlowWithFactorization(h int) (p int) {
	primeFactors64 := primes.Factorize(int64(h)).All()
	// convert to []int
	var primeFactors []int
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

// PresentsForHouse uses a very brute force algorithm to calculate divisibility,
// but is much simpler than the prime factorization algorithm above.
func PresentsForHouse(h int) (p int) {
	for i := 1; i <= h; i++ {
		if h%i == 0 {
			p += i
		}
	}
	return p * 10
}

// SumInts sums a slice of ints (big sigma)
func SumInts(i []int) (s int) {
	for _, x := range i {
		s += x
	}
	return
}

// MultiplyInts returns the product of ints (big pi)
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
