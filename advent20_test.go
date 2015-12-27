package advent

import (
	"reflect"
	"testing"
)

func TestAdvent20InfiniteElves(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"150", 8, 0},
		{"80", 6, 0},
	}

	for i, tt := range tests {
		r1, r2 := Advent20InfiniteElves(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent20InfiniteElves(%s) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}

func TestPresentsForHouse(t *testing.T) {
	tests := []struct {
		in         int
		out1, out2 int
	}{
		{1, 10, 0},
		{2, 30, 0},
		{3, 40, 0},
		{4, 70, 0},
		{5, 60, 0},
		{6, 120, 0},
		{7, 80, 0},
		{8, 150, 0},
		{9, 130, 0},
	}

	for i, tt := range tests {
		r1, r2 := PresentsForHouse(tt.in), 0
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: PresentsForHouse(%d) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)
		}
	}
}

func TestMultiplySumUniqueInts(t *testing.T) {
	if x := MultiplyInts([]int{}); x != 0 {
		t.Errorf("Test MultiplyInts() => %d, want 0", x)
	}

	if x := MultiplyInts([]int{1, 2, 3}); x != 6 {
		t.Errorf("Test MultiplyInts(1, 2, 3) => %d, want 6", x)
	}

	if x := MultiplyInts([]int{2, 2}); x != 4 {
		t.Errorf("Test MultiplyInts(2, 2) => %d, want 4", x)
	}

	if x := SumInts([]int{}); x != 0 {
		t.Errorf("Test SumInts() => %d, want 0", x)
	}

	if x := SumInts([]int{2, 3, 4}); x != 9 {
		t.Errorf("Test SumInts(2, 3, 4) => %d, want 9", x)
	}

	if x := UniqueInts([]int{2, 4, 3}); !reflect.DeepEqual(x, []int{2, 3, 4}) {
		t.Errorf("Test UniqueInts(2,4,3) => %v, want [2,3,4]", x)
	}

	if x := UniqueInts([]int{1, 0, 1, 0}); !reflect.DeepEqual(x, []int{0, 1}) {
		t.Errorf("Test UniqueInts(1, 1, 0, 0) => %v, want [0, 1]", x)
	}

}
