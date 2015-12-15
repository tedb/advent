package advent

import (
	"testing"
)

func TestMin3(t *testing.T) {
	if min3(1, 2, 3) != 1 {
		t.Errorf("min3(1, 2, 3) != 1")
	}
	if min3(3, 2, 1) != 1 {
		t.Errorf("min3(3, 2, 1) != 1")
	}
	if min3(1, 3, 2) != 1 {
		t.Errorf("min3(1, 3, 2) != 1")
	}
	if min3(2, 2, 1) != 1 {
		t.Errorf("min3(2, 2, 1) != 1")
	}
	if min3(1, 2, 2) != 1 {
		t.Errorf("min3(1, 2, 2) != 1")
	}
}

func TestAdvent2a(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 0},
	}

	for i, tt := range tests {
		r := Advent2a_Box(tt.in)
		if r != tt.out {
			t.Errorf("Test %d: Advent2a_Box(%q) => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}
