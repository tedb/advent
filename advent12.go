// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"encoding/json"
	//"fmt"
)

// Advent12JSON parses a JSON string, and returns the sum of all the numbers found
// recursively in the structure, as well as the sum of the numbers where objects
// containing the key "red" are notched out.  Part A was solved with this Bash/Perl
// one-liner as a hack:
//
//     egrep -o '[0-9-]+' data/advent12.txt | perl -ne '$sum+=$_; END{print $sum}'`)
func Advent12JSON(s string) (sumNoExclusions, sumWithoutRed int) {
	var j interface{}
	err := json.Unmarshal([]byte(s), &j)
	checkErr(err)

	//println(s)
	return int(JSONRecurseSum(j, "")), int(JSONRecurseSum(j, "red"))
}

// JSONRecurseSum traverses a JSON structure recursively, summing
// all integers found, while excluding objects (not arrays) containing the key "red"
func JSONRecurseSum(jUncast interface{}, skip string) (sum float64) {
	switch j := jUncast.(type) {
	// j is a map
	case map[string]interface{}:
		//println("map len j:", len(j))
		for _, vUncast := range j {
			//fmt.Printf("%d map vUncast type: %T\n", k, vUncast)
			switch v := vUncast.(type) {
			case string:
				if v == skip && skip != "" {
					//println(k, "map skip for ", v, "==", skip)
					return 0
				}
			case float64:
				sum += v
				//println(k, "map int added", v, "=", sum)
			default:
				r := JSONRecurseSum(v, skip)
				sum += r
				//println(k, "map recurse added", r, "=", sum)

			}

		}
	// j is an array -- don't exclude "red" keys for arrays, only objects
	case []interface{}:
		//println("array len j:", len(j))

		for _, vUncast := range j {
			//fmt.Printf("%d array vUncast type: %T\n", k, vUncast)

			switch v := vUncast.(type) {
			case string:
			case float64:
				sum += v
				//println(k, "array int added", v, "=", sum)
			default:
				r := JSONRecurseSum(v, skip)
				sum += r
				//println(k, "array recurse added", r, "=", sum)
			}

		}
	}
	return
}
