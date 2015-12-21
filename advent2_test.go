package advent

import (
	"testing"
)

func TestAdvent2aMin3(t *testing.T) {
	if Min3(1, 2, 3) != 1 {
		t.Errorf("min3(1, 2, 3) != 1")
	}
	if Min3(3, 2, 1) != 1 {
		t.Errorf("min3(3, 2, 1) != 1")
	}
	if Min3(1, 3, 2) != 1 {
		t.Errorf("min3(1, 3, 2) != 1")
	}
	if Min3(2, 2, 1) != 1 {
		t.Errorf("min3(2, 2, 1) != 1")
	}
	if Min3(1, 2, 2) != 1 {
		t.Errorf("min3(1, 2, 2) != 1")
	}
	if Min3(1, 2, 1) != 1 {
		t.Errorf("min3(1, 2, 1) != 1")
	}
	if Min3(2, 1, 2) != 1 {
		t.Errorf("min3(2, 1, 2) != 1")
	}
	if Min3(1, 1, 2) != 1 {
		t.Errorf("min3(1, 1, 2) != 1")
	}
}

func TestAdvent2aSqftSingle(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"2x3x4", 58},
		{"2x4x3", 58},
		{"1x1x10", 43},
		{"11x11x14", 979},
		{"11x14x11", 979},
		{"14x11x11", 979},
		{"9x9x8", 522},
		{"29x13x26", 3276},
		{"25x2x25", 1500},
	}

	for i, tt := range tests {
		r := NewBox(tt.in).Sqft()
		if r != tt.out {
			t.Errorf("Test %d: NewBox(%q).Sqft() => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}

func TestAdvent2bRibbonSingle(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for i, tt := range tests {
		r := NewBox(tt.in).Ribbon()
		if r != tt.out {
			t.Errorf("Test %d: NewBox(%q).Ribbon() => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}

func TestAdvent2(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"2x3x4\n1x1x10\n2x3x4\n1x1x10\n", 202, 96},
	}

	for i, tt := range tests {
		r1, r2 := Advent02Box(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent2_Box(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
