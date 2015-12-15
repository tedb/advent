package advent

import (
	"testing"
)

func TestAdvent5Naughty(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"ugknbfddgicrmopn", 1, 0},
		{"aaa", 1, 0},
		{"jchzalrnumimnmhp", 0, 0},
		{"haegwjzuvuyypxyu", 0, 0},
		{"dvszwmarrgswjxmb", 0, 0},
	}

	for i, tt := range tests {
		r1, r2 := IsNaughty(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: IsNaughty(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
