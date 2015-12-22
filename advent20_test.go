package advent

import (
	"testing"
)

func TestAdvent20InfiniteElves(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"150", 8, 0},
	}

	for i, tt := range tests {
		r1, r2 := Advent20InfiniteElves(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent20InfiniteElves(%s) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
