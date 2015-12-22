package advent

import (
	"testing"
)

func TestAdvent13Seating(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{
			`Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`, 330, 286},
	}

	for i, tt := range tests {
		r1, r2 := Advent13Seating(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent13Seating(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}

func TestPermutations(t *testing.T) {
	r := Permutations([]int{2, 4, 6, 8}, 0)
	t.Logf("%v", r)
	if len(r) != 24 {
		t.Errorf("Permutations with 4 elements should return 24 permutations, got %d", len(r))
	}

	if r[0][0] != 2 {
		t.Errorf("Permutations: first element of first response should be 2, got %d", r[0][0])
	}

	if r[23][3] != 6 {
		t.Errorf("Permutations: last element of last response should be 6, got %d", r[23][3])
	}

}

func TestSeq(t *testing.T) {
	r := Seq(2, 7)
	if len(r) != 6 {
		t.Errorf("Seq(2,7) should have len 6, got %d", len(r))
	}

	if r[0] != 2 {
		t.Errorf("Seq: first element should be 2, got %d", r[0])
	}

	if r[5] != 7 {
		t.Errorf("Seq: last element should be 7, got %d", r[5])
	}

}
