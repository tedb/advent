package advent

import (
	"testing"
)

func TestAdvent8Matchsticks(t *testing.T) {
	input := `""
              "abc"
              "aaa\"aaa"
              "\x27"
              "trajs\x5brom\xf1yoijaumkem\"\"tahlzs"`
	diff1, diff2 := Advent08Matchsticks(input)

	originalLen := 2 + 5 + 10 + 6 + 38
	want1 := originalLen - 39
	reencodedLen := 6 + 9 + 16 + 11 + 48
	want2 := reencodedLen - originalLen
	if diff1 != want1 || diff2 != want2 {
		t.Errorf("Advent08Matchsticks(`%s`) => %d, %d, want %d, %d", input, diff1, diff2, want1, want2)
	}
}

func TestAdvent8CountDecodedChars(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{`""`, 0, 0},
		{`"abc"`, 3, 0},
		{`"aaa\"aaa"`, 7, 0},
		{`"\x27"`, 1, 0},
		{`"trajs\x5brom\xf1yoijaumkem\"\"tahlzs"`, 28, 0},
	}

	for i, tt := range tests {
		r1, r2 := CountDecodedChars(tt.in), 0
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: CountDecodedChars(`%s`) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}

func TestAdvent8CountReencodedChars(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{`""`, 6, 0},          // "\"\""
		{`"abc"`, 9, 0},       // "\"abc\""
		{`"aaa\"aaa"`, 16, 0}, // "\"aaa\\\"aaa\""
		{`"\x27"`, 11, 0},     // "\"\\x27\""
		// "\"trajs\\x5brom\\xf1yoijaumkem\\\"\\\"tahlzs\""
		{`"trajs\x5brom\xf1yoijaumkem\"\"tahlzs"`, 48, 0},
	}

	for i, tt := range tests {
		r1, r2 := CountReencodedChars(tt.in), 0
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: CountReencodedChars(`%s`) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
