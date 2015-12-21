package advent

import (
	"testing"
)

func TestAdvent21BattleBoss(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"theinput", 0, 0},
	}

	for i, tt := range tests {
		r1, r2 := Foo(tt.in), 0
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Foo(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}
