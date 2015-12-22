package advent

import (
	"testing"
)

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
