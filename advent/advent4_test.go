package advent

import (
	"testing"
)


func TestAdvent4Mining(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 0},
	}

	for i, tt := range tests {
		r, _ := Advent4_Mining(tt.in)
		if r != tt.out {
			t.Errorf("Test %d: Advent4_Mining(%q) => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}
