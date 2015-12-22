package advent

import (
	"bytes"
	"strconv"
)

// Advent10LookSay calls LookSay 40 and 50 times and returns the length of those iterations
func Advent10LookSay(s string) (int, int) {
	var s40 string
	for i := 0; i < 50; i++ {
		if i == 40 {
			s40 = s
		}
		s = LookSay(s)
		//println(" LS ", i, s)
	}
	return len(s40), len(s)
}

// LookSay takes a string of digits and returns the Look-and-say sequence for it
func LookSay(s string) (r string) {
	c := 0
	var last rune
	var b bytes.Buffer

	for i, l := range s {
		if l != last && i > 0 {
			//println(i, last)
			b.WriteString(strconv.Itoa(c))
			b.WriteRune(last)
			c = 0
		}
		c++
		last = l
	}
	b.WriteString(strconv.Itoa(c))
	b.WriteRune(last)

	return b.String()
}
