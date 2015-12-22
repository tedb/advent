package advent

import (
	"testing"
)

func TestAdvent09TSP(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{
			`London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`, 605, 982},
	}

	for i, tt := range tests {
		r1, r2 := Advent09TSP(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent09TSP(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
