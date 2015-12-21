package advent

import (
	"testing"
)

func TestAdvent1a(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for i, tt := range tests {
		r := Advent1aParens(tt.in)
		if r != tt.out {
			t.Errorf("Test %d: Advent1a_Parens(%q) => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}

func TestAdvent1b(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{")", 1},
		{"()())", 5},
	}

	for i, tt := range tests {
		r := Advent1bParensBasement(tt.in)
		if r != tt.out {
			t.Errorf("Test %d: Advent1b_ParensBasement(%q) => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}
