package advent

import (
	"testing"
)

func TestAdvent15Ingredients(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`, 62842880, 57600000},
	}

	for i, tt := range tests {
		r1, r2 := Advent15Ingredients(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent15Ingredients(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
