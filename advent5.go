package advent

import (
	"bufio"
	"regexp"
	"strings"
)

// Advent05Naughty scans over input lines and counts the number of "naughty"
// strings according to defined rules
func Advent05Naughty(s string) (count, z int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		count += IsNaughty(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return
}

// IsNaughty determines whether a string is "naughty" using regex rules
func IsNaughty(s string) int {
	rule1 := regexp.MustCompile("[aeiou].*[aeiou].*[aeiou]")
	rule2 := regexp.MustCompile("aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz")
	badrule3 := regexp.MustCompile("ab|cd|pq|xy")
	if rule1.MatchString(s) && rule2.MatchString(s) && !badrule3.MatchString(s) {
		return 1
	}
	return 0
}
