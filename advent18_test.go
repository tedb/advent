package advent

import (
	"testing"
)

var advent18Steps = []string{
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

func TestAdvent18Animation(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		// input string + 1 step, test number illuminated
		{advent18Steps[0], 11, 18},
		{advent18Steps[0], 8, 18},
		{advent18Steps[0], 4, 18},
		{advent18Steps[0], 4, 14},
	}

	for i, tt := range tests {
		r1, r2 := Advent18Animation(tt.in, i+1)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent18Animation(%q, 1) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)
		}
	}
}

func TestAdvent18NewConway(t *testing.T) {
	c, err := NewConway(advent18Steps[0], '#', '.')
	if err != nil {
		t.Fatalf("Test: NewConway errored: %s", err)
	}

	if c.Lit() != 15 {
		t.Errorf("Test lit at 0 steps: NewConway(%q) => %d, want %d", advent18Steps[0], c.Lit(), 15)
	}
	if c.Width() != 6 || c.Height() != 6 {
		t.Errorf("Test width: NewConway(%q) => %dx%d, want %dx%d", advent18Steps[0], c.Width(), c.Height(), 6, 6)
	}
	if c.NeighborsOn(0, 0) != 1 {
		t.Errorf("Test NeighborsOn(0,0): NewConway(%q) => %d, %d", advent18Steps[0], c.NeighborsOn(0, 0), 1)
	}
	if c.NeighborsOn(5, 2) != 3 {
		t.Errorf("Test NeighborsOn(5,2): NewConway(%q) => %d, %d", advent18Steps[0], c.NeighborsOn(5, 2), 3)
	}
}

func TestAdvent18Step1(t *testing.T) {
	c, _ := NewConway(advent18Steps[0], '#', '.')

	c1 := c.Step(1, false)

	s1 := c1.String()
	if s1 != advent18Steps[1] {
		t.Errorf("Test board after 1 steps: NewConway(%q) + Step(1) => %q, want %q", advent18Steps[0], s1, advent18Steps[1])
	}
}

func TestAdvent18Step4(t *testing.T) {
	c, _ := NewConway(advent18Steps[0], '#', '.')

	c4 := c.Step(4, false)

	s2 := c4.String()
	if s2 != advent18Steps[4] {
		t.Errorf("Test board after 4 steps: NewConway(%q) + Step(4) => %q, want %q", advent18Steps[0], s2, advent18Steps[4])
	}

	if c4.Lit() != 4 {
		t.Errorf("Test lit after 4 steps: NewConway(%q) + Step(4) => %q: %d, want %d", advent18Steps[0], c4, c4.Lit(), 4)
	}
}
