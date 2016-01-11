package advent

import (
	"fmt"
	"strings"
)

// Advent18Animation implements a version of Conway's Game of Life.
func Advent18Animation(s string, steps int) (count, b int) {
	c, err := NewConway(s, '#', '.')
	if err != nil {
		panic(err)
	}
	c.Step(steps)
	return c.Lit(), 0
}

// NewConway takes a string representing a 2d array of lights,
// where the Lit and Unlit runes represent the characters that show whether the cell is lit
// Valid input chars are the Lit and Unlit runes, and newlines; others err
// Edge cases where blank lines or empty matrixes are input will likely fail
func NewConway(input string, Lit rune, Unlit rune) (ConwayBoard, error) {
	size := strings.IndexRune(input, '\n')
	c := NewConwayBoard(size)
	row, col := 0, 0

	for _, r := range input {
		switch r {
		case '\n':
			row++
			col = 0
		case Lit:
			c[row][col] = true
		case Unlit:
			c[row][col] = false
		default:
			return nil, fmt.Errorf("Char %s invalid for NewConway (not \\n, %s, or %s)", r, Lit, Unlit)
		}
	}

	return c, nil
}

// ConwayBoard is a 2d array (matrix) of bool.  True is lit, False is unlit
type ConwayBoard []ConwayRow
type ConwayRow []bool

// NewConwayBoard creates a new Conway ([][]bool) of given size (square)
func NewConwayBoard(size int) ConwayBoard {
	c := make(ConwayBoard, size)
	for i := 0; i < size; i++ {
		c[i] = make(ConwayRow, size)
	}
	return c
}

// Step plays through the Game of Life rules zero or more times
func (c ConwayBoard) Step(steps int) ConwayBoard {
	newBoard := NewConwayBoard(c.Width())
	return newBoard
}

// Lit returns number of cells that are true
func (c ConwayBoard) Lit() int {
	return 0
}

// Width returns width of board (same as Height since it's a square)
func (c ConwayBoard) Width() int {
	return len(c[0])
}

// Height returns height of board (same as Width since it's a square)
func (c ConwayBoard) Height() int {
	return len(c)
}
