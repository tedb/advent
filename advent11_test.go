package advent

import (
	"testing"
)

func TestAdvent11Password(t *testing.T) {
	tests := []struct {
		in   string
		out1 string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for i, tt := range tests {
		r1 := Advent11Password(tt.in)
		if r1 != tt.out1 {
			t.Errorf("Test %d: Advent11Password(%q) => '%s', want '%s'", i, tt.in, r1, tt.out1)

		}
	}
}

func TestAdvent11IncrementString(t *testing.T) {
	tests := []struct {
		in   string
		out1 string
	}{
		{"abcdefgh", "abcdefgi"},
		{"ghijklmn", "ghijklmo"},
		{"ghijklmz", "ghijklna"},
		{"bbbbkzzz", "bbbblaaa"},
		{"yzzzzzzz", "zaaaaaaa"},
	}

	for i, tt := range tests {
		r1 := IncrementString(tt.in)
		if r1 != tt.out1 {
			t.Errorf("Test %d: IncrementString(%q) => '%s', want '%s'", i, tt.in, r1, tt.out1)

		}
	}
}

func TestAdvent11CheckPassword(t *testing.T) {
	tests := []struct {
		in   string
		out1 bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdefgh", false},
		{"abcdffaa", true},
		{"ghijklmn", false},
		{"ghjaabcc", true},
	}

	for i, tt := range tests {
		r1 := CheckPassword(tt.in)
		if r1 != tt.out1 {
			t.Errorf("Test %d: CheckPassword(%q) => %v, want %v", i, tt.in, r1, tt.out1)

		}
	}
}
