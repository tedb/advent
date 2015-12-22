package advent

import (
	"testing"
)

func TestAdvent10LookSay(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 string
	}{
		{"1", "11", ""},
		{"11", "21", ""},
		{"21", "1211", ""},
		{"1211", "111221", ""},
		{"111221", "312211", ""},
	}

	for i, tt := range tests {
		r1, r2 := LookSay(tt.in), ""
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: LookSay(%q) => %s, %s, want %s, %s", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}

func BenchmarkLookSay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LookSay("311311221112131221123113112211322112211213322113")
	}
}
