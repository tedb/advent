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
		r := NewRoute().Nav(tt.in).How_many_unique()
		if r != tt.out {
			t.Errorf("Test %d: NewRoute(%q).Houses() => %d, want %d", i, tt.in, r, tt.out)

		}
	}
}

