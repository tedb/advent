package advent

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// Advent04Mining brute forces MD5 hashes
func Advent04Mining(s string) (winner1, winner2 int) {
	inputPrefix := []byte(s)

	for i := 1; i < 10000000; i++ {
		//println(i)
		byteTrial := []byte(strconv.Itoa(i))
		cand := append(inputPrefix, byteTrial...)
		hex := fmt.Sprintf("%x", md5.Sum(cand))
		//fmt.Printf("%s, %s, %s, %s = %s\n", s, inputPrefix, byteTrial, cand, hex)
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
