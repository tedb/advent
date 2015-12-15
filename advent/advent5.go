// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"regexp"
	"strings"
)

// Advent4_Mining brute forces MD5 hashes to
func Advent5_Naughty(s string) (count, z int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += IsNaughty(scanner.Text())
	}

	err := scanner.Err()
	check_err(err)

	return
}

func IsNaughty(s string) int {
	rule1 := regexp.MustCompile("[aeiou].*[aeiou].*[aeiou]")
	rule2 := regexp.MustCompile("aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz")
	badrule3 := regexp.MustCompile("ab|cd|pq|xy")
	if rule1.MatchString(s) && rule2.MatchString(s) && !badrule3.MatchString(s) {
		return 1
	}
	return 0
}
