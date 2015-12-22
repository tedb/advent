package advent

import (
	"bufio"
	"fmt"
	"strings"
)

// Advent15Ingredients (NOT DONE) is a template function for other Advent of Code puzzles to start from
func Advent15Ingredients(s string) (score1, score2 int) {
	i := make(Ingredients, 0)
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		i = i.AddFromString(line)
	}

	err := scanner.Err()
	checkErr(err)

	//fmt.Printf("%+v", i)

	c := i.NewCookie()
	score1 = c.BestCookie(0)
	score2 = c.BestCookie(500)
	return
}

type Ingredients []Ingredient

func (i Ingredients) AddFromString(s string) Ingredients {
	newI := Ingredient{}
	_, err := fmt.Sscanf(s, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&newI.Name, &newI.Capacity, &newI.Durability, &newI.Flavor,
		&newI.Texture, &newI.Calories)
	checkErr(err)

	i = append(i, newI)
	return i
}

type Ingredient struct {
	Name                                            string
	Capacity, Durability, Flavor, Texture, Calories int
	Quantity                                        int
}

func (i Ingredients) NewCookie() (c Cookie) {
	c = make(Cookie, len(i))
	copy(c, i)

	return c
}

type Cookie []Ingredient

// BestCookie combines its ingredients such that the ingredient quantity sum == 100
// and cookie "score" is maximized.  Brute force all ingredient quantities 0..100
// until we find the best one
func (c Cookie) BestCookie(calorieTarget int) (max int) {
	//var max int
	//var best Cookie
	var scoreCounter int
	// this is a super nasty hack... replace with a recursive function?
	for j := 0; j <= 100; j++ {
		for k := 0; k <= 100; k++ {
			//println("j k", j, k)
			for l := 0; l <= 100; l++ {
				for m := 0; m <= 100; m++ {
					jQ, kQ, lQ, mQ := j, k, l, m
					if len(c) < 1 {
						jQ = 0
					}
					if len(c) < 2 {
						kQ = 0
					}
					if len(c) < 3 {
						lQ = 0
					}
					if len(c) < 4 {
						mQ = 0
					}

					totalQuantity := jQ + kQ + lQ + mQ
					if totalQuantity > 100 {
						continue
					}

					if len(c) >= 1 {
						c[0].Quantity = j
					}
					if len(c) >= 2 {
						c[1].Quantity = k
					}
					if len(c) >= 3 {
						c[2].Quantity = l
					}
					if len(c) >= 4 {
						c[3].Quantity = m
					}

					scoreCounter++
					s := c.Score()
					if s > max {
						if calorieTarget != 0 && c.Calories() != 500 {
							continue
						}
						max = s
						//best = c
						//println("best", max, best)
					}

				}
			}
		}
	}

	//fmt.Println("It took us", scoreCounter, "tries")
	return //best
}

func (c Cookie) Score() (s int) {
	var sC, sD, sF, sT int

	for _, i := range c {
		//fmt.Printf("%v", i)
		sC += i.Quantity * i.Capacity
		sD += i.Quantity * i.Durability
		sF += i.Quantity * i.Flavor
		sT += i.Quantity * i.Texture

	}
	if sC < 0 {
		sC = 0
	}
	if sD < 0 {
		sD = 0
	}
	if sF < 0 {
		sF = 0
	}
	if sT < 0 {
		sT = 0
	}
	return sC * sD * sF * sT
}

func (c Cookie) Calories() (s int) {
	for _, i := range c {
		s += i.Quantity * i.Calories
	}
	return s
}
