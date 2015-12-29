package advent

import (
	"bufio"
	"strings"
)

// Advent08Matchsticks decodes escaped strings and counts the decoded chars,
// minus input string whitespace outside the quoted strings
func Advent08Matchsticks(s string) (diff, b int) {
	var countDecoded, countOriginal, countReencoded int
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		s := scanner.Text()
		countDecoded += CountDecodedChars(s)
		countReencoded += CountReencodedChars(s)
		countOriginal += len(s)
	}

	err := scanner.Err()
	checkErr(err)

	return countOriginal - countDecoded, countReencoded - countOriginal
}

// CountDecodedChars decodes a string with backslash escapes and start/end quotes,
// and returns the length of decoded string.  Decoded count is guaranteed to be
// less than len(s).  Note this will exhibit undefined behavior for malformed strings.
func CountDecodedChars(s string) (c int) {
	r := []rune(s)
	// skip first and last char (quotes)
	for i := 1; i < len(r)-1; i++ {
		// hex: \xAA
		if r[i] == '\\' && r[i+1] == 'x' {
			// skip an extra 3 chars
			i += 3
			// another escape
		} else if r[i] == '\\' {
			// skip an extra char
			i += 1
		}

		c++
	}
	return c
}

// CountReencodedChars counts the chars required to escape
// the input string, including starting/ending quotes
func CountReencodedChars(s string) (c int) {
	c = 2
	for _, r := range s {
		if r == '"' || r == '\\' {
			c += 2
		} else {
			c++
		}

	}
	return c
}
