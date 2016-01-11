package advent

import (
	"fmt"
	"strings"
)

// Advent18Animation implements a version of Conway's Game of Life.
func Advent18Animation(s string, steps int) (a, b int) {
	c, err := NewConway(s, '#', '.')
	if err != nil {
		panic(err)
	}
	c1 := c.Step(steps, false)
	c2 := c.Step(steps, true)

	return c1.Lit(), c2.Lit()
}

// NewConway takes a string representing a 2d array of lights,
// where the Lit and Unlit runes represent the characters that show whether the cell is lit
// Valid input chars are the Lit and Unlit runes, and newlines; others err
// Edge cases where blank lines or empty matrixes are input will likely fail
func NewConway(input string, Lit rune, Unlit rune) (ConwayBoard, error) {
	size := strings.IndexRune(input, '\n')
	//fmt.Println("NewConway of size", size)
	if size <= 1 {
		return nil, fmt.Errorf("Invalid size %d for input length %d", size, len(input))
	}
	c := NewConwayBoard(size)
	row, col := 0, 0

	for _, r := range input {
		switch r {
		case '\n':
			//fmt.Printf("at row %d:\n%v\n", row, c)
			row++
			col = 0
		case Lit:
			c[row][col] = true
			col++
		case Unlit:
			// false by default
			col++
		default:
			return nil, fmt.Errorf("Char %v invalid for NewConway (not \\n, %v, or %v)", r, Lit, Unlit)
		}
	}

	return c, nil
}

// ConwayBoard is a 2d array (matrix) of bool.  True is lit, False is unlit
// As an alternative, this could be presented as a vector of size^2 elements,
// but that makes our algorithm a little more contrived.
type ConwayBoard []ConwayRow

// ConwayRow is a row within ConwayBoard
type ConwayRow []bool

// NewConwayBoard creates a new Conway ([][]bool) of given size (square)
func NewConwayBoard(size int) ConwayBoard {
	c := make(ConwayBoard, size)
	for i := 0; i < size; i++ {
		c[i] = make(ConwayRow, size)
	}
	return c
}

// Step plays through the Game of Life rules zero or more times.
// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
// TODO: this creates a lot of garbage; potential for refactoring to recycle the
// previous ConwayBoard.
func (c ConwayBoard) Step(steps int, cornersStuckOn bool) ConwayBoard {
	// Part 2 of puzzle has this twist
	if cornersStuckOn {
		// top left
		c[0][0] = true
		// top right
		c[c.Width()-1][0] = true
		// bottom left
		c[0][c.Height()-1] = true
		// bottom right
		c[c.Width()-1][c.Height()-1] = true
	}

	for i := 0; i < steps; i++ {
		newBoard := NewConwayBoard(c.Width())

		for row, rowData := range c {
			for col, cell := range rowData {
				neighborsOn := c.NeighborsOn(row, col)
				if (cell && neighborsOn == 2) || neighborsOn == 3 {
					newBoard[row][col] = true
				}
			}
		}

		c = newBoard

		if cornersStuckOn {
			// top left
			c[0][0] = true
			// top right
			c[c.Width()-1][0] = true
			// bottom left
			c[0][c.Height()-1] = true
			// bottom right
			c[c.Width()-1][c.Height()-1] = true
		}

		fmt.Printf("After %d steps:\n%s\n", i+1, c.String())
	}

	return c
}

// NeighborsOn returns the number the 8 neighboring cells that are true,
// given a row and col number.  Cells off the board (left, right, top, bottom)
// are considered false.
func (c ConwayBoard) NeighborsOn(row, col int) (count int) {
	maxRow, maxCol := c.Height()-1, c.Width()-1

	// left
	if col > 0 && c[row][col-1] {
		count++
	}
	// right
	if col < maxCol && c[row][col+1] {
		count++
	}

	// all the tops
	if row > 0 {
		// top left
		if col > 0 && c[row-1][col-1] {
			count++
		}
		// top
		if c[row-1][col] {
			count++
		}
		// top right
		if col < maxCol && c[row-1][col+1] {
			count++
		}
	}

	// all the bottoms
	if row < maxRow {
		// bottom right
		if col < maxCol && c[row+1][col+1] {
			count++
		}
		// bottom
		if c[row+1][col] {
			count++
		}
		// bottom left
		if col > 0 && c[row+1][col-1] {
			count++
		}
	}

	return count
}

// Lit returns number of cells that are true
func (c ConwayBoard) Lit() (count int) {
	//fmt.Printf("Lit:\n%v\n", c)
	for _, rowData := range c {
		for _, cell := range rowData {
			if cell {
				count++
			}
		}
	}
	return
}

// Width returns width of board (same as Height since it's a square)
func (c ConwayBoard) Width() int {
	return len(c[0])
}

// Height returns height of board (same as Width since it's a square)
func (c ConwayBoard) Height() int {
	return len(c)
}

// String returns a string representation of the board
func (c ConwayBoard) String() (s string) {
	for _, rowData := range c {
		for _, cell := range rowData {
			if cell {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return
}
