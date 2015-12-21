package advent

import (
	"testing"
)

func TestAdvent4Mining(t *testing.T) {
	t.Skip("This test sometimes takes several seconds to run")

	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"abcdef", 609043, 6742839},
		{"pqrstuv", 1048970, 5714438},
	}

	for i, tt := range tests {
		r1, r2 := Advent04Mining(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent4_Mining(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)
		}
	}
}
