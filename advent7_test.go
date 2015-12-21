package advent

import (
	"testing"
)

func TestAdvent7Wires(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"100 -> d\n200 -> a", "a: 200\nd: 100"},
		{"100 -> a\na LSHIFT 2 -> b", "a: 100\nb: 400"},
		{"a RSHIFT 2 -> b\n100 -> a", "a: 100\nb: 25"},
		{"100 -> a\n200 -> b\na AND b -> c", "a: 100\nb: 200\nc: 64"},
		{"100 -> a\n200 -> b\na OR b -> c", "a: 100\nb: 200\nc: 236"},
		{"NOT a -> b\n100 -> a", "a: 100\nb: 65435"},
		{"100 -> a\na -> b", "a: 100\nb: 100"},
		{"100 -> a\nNOT a -> b\nb RSHIFT 2 -> c", "a: 100\nb: 65435\nc: 16358"},
		{
			`123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			`d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456`},
	}

	// helpful site for binary-decimal conversions: https://www.branah.com/ascii-converter

	for i, tt := range tests {
		r1 := Advent7Wires(tt.in)
		if r1 != tt.out {
			t.Errorf("Test %d: Advent7_Wires(%q) =>\n%s\n- want -\n%s\n", i, tt.in, r1, tt.out)

		}
	}
}

func TestStrOrInt(t *testing.T) {
	tests := []struct {
		in   string
		out1 string
		out2 uint16
	}{
		{"0", "", 0},
		{"1", "", 1},
		{"10", "", 10},
		{"fad", "fad", 0},
		{"3fdsfsa", "3fdsfsa", 0},
	}

	// helpful site for binary-decimal conversions: https://www.branah.com/ascii-converter

	for i, tt := range tests {
		r1, r2 := StringOrInt(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: str_or_int(%q) => %s, %d want %s, %d\n", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
