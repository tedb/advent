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
		{"())", 3},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for i, tt := range tests {
		r := Advent1a_Parens(tt.in)
		if r != tt.out {
			t.Errorf("Test %d: Advent1a_Parens(%q) => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}
