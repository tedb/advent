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
		{"100 -> a\na RSHIFT 2 -> b", "a: 100\nb: 25"},
		{"100 -> a\n200 -> b\na AND b -> c", "a: 100\nb: 200\nc: 64"},
		{"100 -> a\n200 -> b\na OR b -> c", "a: 100\nb: 200\nc: 236"},
		{"100 -> a\nNOT a -> b", "a: 100\nb: 65435"},
		{"100 -> a\na -> b", "a: 100\nb: 100"},
	}

	for i, tt := range tests {
		r1 := Advent7_Wires(tt.in)
		if r1 != tt.out {
			t.Errorf("Test %d: Advent7_Wires(%q) =>\n%s\n- want -\n%s\n", i, tt.in, r1, tt.out)

		}
	}
}
