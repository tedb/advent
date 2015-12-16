package advent

import (
	"testing"
)

func TestAdvent7Wires(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{{
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
		{"100 -> d\n200 -> a", "a: 200\nd: 100"},
		{"100 -> a\na LSHIFT 2 -> b", "a: 100\nb: 400"},
		{"a RSHIFT 2 -> b\n100 -> a", "a: 100\nb: 25"},
		{"100 -> a\n200 -> b\na AND b -> c", "a: 100\nb: 200\nc: 64"},
		{"100 -> a\n200 -> b\na OR b -> c", "a: 100\nb: 200\nc: 236"},
		{"NOT a -> b\n100 -> a", "a: 100\nb: 65435"},
		{"100 -> a\na -> b", "a: 100\nb: 100"},
		{"100 -> a\nNOT a -> b\nb RSHIFT 2 -> c", "a: 100\nb: 65435\nc: 16358"},
		{"100 -> a\nNOT a -> b\nb RSHIFT 2 -> b", "a: 100\nb: 16358"},
	}

	// helpful site for binary-decimal conversions: https://www.branah.com/ascii-converter

	for i, tt := range tests {
		r1 := Advent7_Wires(tt.in)
		if r1 != tt.out {
			t.Errorf("Test %d: Advent7_Wires(%q) =>\n%s\n- want -\n%s\n", i, tt.in, r1, tt.out)

		}
	}
}

func TestStrOrInt(t *testing.T) {
	tests := []struct {
		in  string
		out1 bool
		out2 string
		out3 uint16
	}{
		{"0", true, nil, 0},
		{"-1", true, nil, -1},
		{"10", true, nil, 10},
		{"fad", false, "fad", 0},
		{"3fdsfsa", false, "3fdsfsa", 0},
	}

	// helpful site for binary-decimal conversions: https://www.branah.com/ascii-converter

	for i, tt := range tests {
		r1, r2, r3 := str_or_int(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 || r3 != tt.out3 {
			t.Errorf("Test %d: str_or_int(%q) => %d, %s, %d want %d, %s, %d\n", i, tt.in, r1, r2, r3, tt.out1, tt.out2, tt.out3)

		}
	}	
}
