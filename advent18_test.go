package advent

import (
	"testing"
)

func TestAdvent18Animation(t *testing.T) {
	steps := []string{
		// 0
		`.#.#.#
...##.
#....#
..#...
#.#..#
####..
`,
		// 1
		`..##..
..##.#
...##.
......
#.....
#.##..
`,
		// 2
		`..###.
......
..###.
......
.#....
.#....
`,
		// 3
		`...#..
......
...#..
..##..
......
......
`,
		// 4
		`......
......
..##..
..##..
......
......
`}
	tests := []struct {
		in         string
		out1, out2 int
	}{
		// input string + 1 step, test number illuminated
		{steps[0], 11, 0},
		{steps[1], 8, 0},
		{steps[2], 4, 0},
		{steps[3], 4, 0},
	}

	for i, tt := range tests {
		r1, r2 := Advent18Animation(tt.in, 1)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent18Animation(%q, 1) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)
		}
	}

	c, err := NewConway(steps[0], '#', '.')
	if err != nil {
		t.Fatalf("Test: NewConway errored: %s", err)
	}

	if c.Lit() != 15 {
		t.Errorf("Test: NewConway(%q) => %d, want %d", steps[0], c.Lit(), 15)
	}
	if c.Width() != 6 || c.Height() != 6 {
		t.Errorf("Test: NewConway(%q) => %dx%d, want %dx%d", steps[0], c.Width(), c.Height(), 6, 6)
	}

	c = c.Step(4)

	if c.Lit() != 4 {
		t.Errorf("Test: NewConway(%q) + Step(4) => %d, want %d", steps[0], c.Lit(), 4)
	}
}
