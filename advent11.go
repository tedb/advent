package advent

import (
	//"strings"
	"regexp"
	//"fmt"
)

// Advent11Password returns the next incrementally allowable password
// given the current password
func Advent11Password(s string) string {
	for {
		s = IncrementString(s)
		if CheckPassword(s) {
			break
		}
	}

	return s
}

// IncrementString takes an input string and increments it just like counting with numbers
func IncrementString(s string) string {
	x := []rune(s)

	for i := len(x) - 1; i >= 0; i-- {
		if x[i] != 'z' {
			x[i]++
			return string(x)
		}
		x[i] = 'a'
	}

	return string(x)
}

// Passwords must include one increasing straight of at least three letters,
// like abc, bcd, cde, and so on, up to xyz. They cannot skip letters;
// abd doesn't count.
// Ruby: ("a".."x").map{|x| x+(x.next)+(x.next.next)}.join('|')
var checkPasswordR1 = regexp.MustCompile("abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz")

// Passwords may not contain the letters i, o, or l, as these letters can be
// mistaken for other characters and are therefore confusing.
var checkPasswordR2 = regexp.MustCompile("i|o|l")

// Passwords must contain at least two different, non-overlapping pairs of
// letters, like aa, bb, or zz.
var checkPasswordR3 = regexp.MustCompile("(aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz)")

// CheckPassword checks a password candidate to see if it matches the necessary rules
func CheckPassword(s string) bool {
	if !checkPasswordR1.MatchString(s) || checkPasswordR2.MatchString(s) {
		return false
	}

	r3Matches := UniqueStrings(checkPasswordR3.FindAllString(s, -1))
	if len(r3Matches) < 2 {
		return false
	}

	return true
}

// UniqueStrings takes a slice of strings and returns the unique values
// (in arbitrary order)
func UniqueStrings(s []string) (r []string) {
	empty := struct{}{}
	seen := make(map[string]struct{})
	for _, val := range s {
		if _, ok := seen[val]; !ok {
			r = append(r, val)
			seen[val] = empty
		}
	}
	return r
}
