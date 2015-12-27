package advent

import (
	"github.com/ntns/goitertools/itertools"
	"github.com/otiai10/primes"
	"sort"
	"strconv"
	"sync"
)

// Advent20InfiniteElves determines the lowest house number of the house
// to get at least as many presents as the number in the puzzle input.
// Tried using Sum of Divisors algorithm, and was successful, but it
// was super slow.  Trying brute force algorithm.
func Advent20InfiniteElves(presentsStr string) (house1, house2 int) {
	presents, err := strconv.Atoi(presentsStr)
	checkErr(err)

	houses1 := make([]int, presents/10)
	houses2 := make([]int, 50+1)
	houses2Done := 0

	var wg sync.WaitGroup

	c1 := make(chan int, 0)
	c2 := make(chan int, 0)

	for mainElf := 1; ; mainElf++ {
		wg.Add(1)
		go func(elf int) {
			defer wg.Done()
			if elf%100 == 0 {
				//println("done with elf", elf)
			}
			for houseIdx := elf; houseIdx < len(houses1); houseIdx += elf {
				houses1[houseIdx] += elf * 10
				//fmt.Println("delivering for elf", elf, "to house", houseIdx, "=", houses[houseIdx])
			}

			for i, v := range houses1 {
				if v >= presents && elf >= i && houses2Done != 0 {
					//fmt.Printf("winning houses(%d): i=%d, v=%d, elf=%d %v\n", presents, i, v, elf, houses[1:])
					c1 <- i
					break
				}
			}

			for houseIdx := elf; houseIdx < len(houses2); houseIdx += elf {
				houses2[houseIdx] += elf * 11
			}

			for i, v := range houses2 {
				if v >= presents && elf >= i && elf >= len(houses2) {
					//fmt.Printf("winning houses(%d): i=%d, v=%d, elf=%d %v\n", presents, i, v, elf, houses[1:])
					c2 <- i
					break
				}
			}
		}(mainElf)
	}
	wg.Wait()
	return <-c1, <-c2
}

// Advent20InfiniteElves determines the lowest house number of the house
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
	return
}

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

// PresentsForHouse returns "sum of divisors" for the given house number, times ten.
// Divisors (or factors) are calculated by determining h's unique prime factors,
// then appending 1 and h to the list, rendering all pairwise combinations
// of the unique factors, then summing those products.  "Sum of divisors"
// is also [OEIS sequence A000203](http://oeis.org/A000203).  A simpler way to solve this
// would be just looping through 1..h and checking divisibility, but using more math makes it fun.
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
