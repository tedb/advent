// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// Advent4Mining brute forces MD5 hashes to
func Advent4Mining(s string) (winner1, winner2 int) {
	for i := 1; i < 10000000; i++ {
		//println(i)
		byte_s := []byte(s)
		byte_i := []byte(strconv.Itoa(i))
		cand := append(byte_s, byte_i...)
		hex := fmt.Sprintf("%x", md5.Sum(cand))
		//fmt.Printf("%s, %s, %s, %s = %s\n", s, byte_s, byte_i, cand, hex)
		if winner1 == 0 && hex[0:5] == "00000" {
			winner1 = i
		}
		if winner2 == 0 && hex[0:6] == "000000" {
			winner2 = i
		}
		if winner1 > 0 && winner2 > 0 {
			return
		}
	}
	return 0, 0
}
