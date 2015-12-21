package advent

import (
	"testing"
)

func TestAdvent3Houses(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for i, tt := range tests {
		r := NewRoute().Nav(tt.in).HowManyUnique()
		if r != tt.out {
			t.Errorf("Test %d: NewRoute().Nav(%q).How_many_unique() => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}

func TestAdvent3HousesRoboSanta(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for i, tt := range tests {
		r := NewRoute().DualNav(tt.in).HowManyUnique()
		if r != tt.out {
			t.Errorf("Test %d: NewRoute().DualNav(%q).How_many_unique() => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}
