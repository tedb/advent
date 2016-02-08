package advent

import (
	"testing"
)

func TestAdvent19Replacements(t *testing.T) {
	tests := []struct {
		in1, in2   string
		out1, out2 int
	}{
		{"HOH", "H => HO\nH => OH\nO => HH", 4, 0},
		{"HOHOHO", "H => HO\nH => OH\nO => HH", 7, 0},
		{"H2O", "H => OO", 1, 0},
	}

	for i, tt := range tests {
		r1, r2 := Advent19DistinctMolecules(tt.in1, tt.in2), 0
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent19DistinctMolecules(%q, %q) => %d, %d, want %d, %d", i, tt.in1, tt.in2, r1, r2, tt.out1, tt.out2)

		}
	}
}

func TestAdvent19MoleculeSwapsNone(t *testing.T) {
	r := UniqueStrings(MoleculeSwaps("ABCD", "YZ", "M", 0))
	if len(r) != 0 {
		t.Errorf("should be no swaps, got %d", len(r))
	}
}

func TestAdvent19MoleculeSwapsOne(t *testing.T) {
	r := UniqueStrings(MoleculeSwaps("ABCD", "BC", "M", 0))
	if len(r) != 1 {
		t.Errorf("should be 1 swaps, got %d", len(r))
	}
	if r[0] != "AMD" {
		t.Errorf("should be ABCD(BC => M) = AMD, got %s", r[0])
	}
}

func TestAdvent19MoleculeSwapsThree(t *testing.T) {
	r := UniqueStrings(MoleculeSwaps("ABZZABZZZAB", "AB", "M", 0))
	if len(r) != 3 {
		t.Errorf("should be 3 swaps, got %d", len(r))
	}
	if r[0] != "MZZABZZZAB" {
		t.Errorf("should be ABZZABZZZAB(AB => M) = MZZABZZZAB #1, got %s", r[0])
	}
	if r[1] != "ABZZMZZZAB" {
		t.Errorf("should be ABZZABZZZAB(AB => M) = ABZZMZZZAB #2, got %s", r[1])
	}
	if r[2] != "ABZZABZZZM" {
		t.Errorf("should be ABZZABZZZAB(AB => M) = ABZZABZZZM #3, got %s", r[2])
	}
}
