package advent

import (
	"testing"
)

func TestAdvent4Mining(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"abcdef", 609043, 0},
		{"pqrstuv", 1048970, 0},
	}

	for i, tt := range tests {
		r1, r2 := Advent4_Mining(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent4_Mining(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
