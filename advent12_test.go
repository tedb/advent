package advent

import (
	"testing"
)

func TestAdvent12JSON(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"[1,2,3]", 6, 6},
		{"{\"a\":2,\"b\":4}", 6, 6},
		{"[[[3]]]", 3, 3},
		{"{\"a\":{\"b\":4},\"c\":-1}", 3, 3},
		{"{\"a\":[-1,1]}", 0, 0},
		{"[-1,{\"a\":1}]", 0, 0},
		{"[]", 0, 0},
		{"{}", 0, 0},
		{"[1,{\"c\":\"red\",\"b\":2},3]", 6, 4},
		{"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 15, 0},
		{"[1,\"red\",5]", 6, 6},
	}

	for i, tt := range tests {
		r1, r2 := Advent12JSON(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent12JSON(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
