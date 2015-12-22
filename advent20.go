package advent

import (
	"strconv"
)

// Advent20InfiniteElves determines the lowest house number of the house
// to get at least as many presents as the number in the puzzle input
func Advent20InfiniteElves(presents_s string) (house, b int) {
	presents, err := strconv.Atoi(presents_s)
	checkErr(err)
	houses := make([]int, 10000000)

	for _, i := range Seq(1, len(houses)) {
		for _, elf := range(Seq(1, len(houses))) {
			if i % elf == 0 {
				houses[i] += elf * 10
			}
			if houses[i] >= presents {
				//println(houses[i], i, presents)
				return i, 0
			}
		}
	}
	//println(houses[0:20])
	return

	// for elf := 1; ; elf++ {
	// 			//println("elf", elf)

	// 	for house := elf; house < len(houses); house += elf {
	// 		if house == elf {
	// 		println("elf", elf, "house", house)
	// 		}
	// 		houses[house] += elf * 10

	// 	}
	// }
}
