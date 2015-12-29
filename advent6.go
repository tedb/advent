package advent

import (
	"math/big"
	"strconv"
	"strings"
	//	"fmt"
	"bytes"
	"regexp"
)

// Advent06Lights scans lines of input and follows instructions
// to toggle lights in a matrix (toggle, on, off) according to specified
// rectangles, e.g.:
//     turn on 887,9 through 959,629
//     turn off 539,243 through 559,965
//     toggle 393,804 through 510,976
// This solution makes creative use of a "bitboard" that's 1 million bits
// in length.  However, the puzzle assumes [][]bool or equivalent,
// which burned me on part 2, which wants each light to be a variable
// brightness instead of on/off.
func Advent06Lights(s string) (count, b int) {
	grid := NewLightGridInt(1000)
	grid.ApplyCommands(s)
	//println(grid.ToBinarySquare())
	return grid.PopCount(), 0

}

// NewLightGridInt creates a new LightGridInt, which uses math/big
// to implement a bit array that we operate on.  Size is one side of
// a square grid, with total number of lights = Size * Size
func NewLightGridInt(size int) (g *LightGridInt) {
	return &LightGridInt{big.NewInt(0), size}
}

// LightGridInt implements a "bitboard" where the NxN matrix
// is represented as a BigInt
type LightGridInt struct {
	I *big.Int
	// grid is this length on one side; Size*Size lights
	Size int
}

// ApplyCommands applies commands from 1 or more line of input, e.g.:
//     turn on 887,9 through 959,629
//     turn off 539,243 through 559,965
//     toggle 393,804 through 510,976
// Parser is implemented with regex this time, although could also be implemented
// with fmt.Sscanf.
func (g *LightGridInt) ApplyCommands(s string) {
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	for i, m := range re.FindAllStringSubmatch(s, -1) {
		x1, err := strconv.Atoi(m[2])
		checkErr(err)
		y1, err := strconv.Atoi(m[3])
		checkErr(err)
		x2, err := strconv.Atoi(m[4])
		checkErr(err)
		y2, err := strconv.Atoi(m[5])
		checkErr(err)

		switch m[1] {
		case "turn on":
			g.RectSetOn(x1, y1, x2, y2)
		case "turn off":
			g.RectSetOff(x1, y1, x2, y2)
		case "toggle":
			g.RectToggle(x1, y1, x2, y2)
		default:
			panic("bad command")
		}

		println("ApplyCommands:", i, m[1])
		//println(i, m[0], "\n", g.ToBinarySquare())
	}
}

// MaskFromRectangle returns a BigInt bit mask that can be applied
// against LightGridInt to turn on, turn off, or toggle the masked bits.
// 0,0 is top-left.  In this example for size=4, rect(1, 2, 3, 3) is:
//     . . . . .
//     . . . . .
//     . o o o .
//     . o o o .
//     . . . . .
// turns into:
//    . . . . . . . . . . . o o o . . o o o . . . . . .
// ...which is 14784 (least significant bit is on the right).
func (g *LightGridInt) MaskFromRectangle(x1, y1, x2, y2 int) (mask *big.Int) {
	rW := x2 - x1 + 1
	rH := y2 - y1 + 1

	// zeroes between the end of one row of 1's and the start of the next
	spaceBetween := g.Size - rW

	// how many zeroes between the bottom-right
	// corner of the rectangle and the end of the square;
	// in the example above, 6
	spaceAtEnd := (g.Size * (g.Size - 1 - y2)) + (g.Size - 1 - x2) - spaceBetween
	//println("spaceBetween, spaceAtEnd", spaceBetween, spaceAtEnd)
	//_ = "breakpoint"

	mask = big.NewInt(0)
	for i := 0; i < rH; i++ {
		// could replace this with exponent, but that uses float which
		// could be imprecise with large values
		for digit := 0; digit < rW; digit++ {
			mask.Lsh(mask, uint(1))
			mask.Add(mask, big.NewInt(1))
			//println("mask", mask.Int64(), BigIntToBinary(mask))
		}
		mask.Lsh(mask, uint(spaceBetween))
		//println("mask with space", mask.Int64(), BigIntToBinary(mask))

	}
	mask.Lsh(mask, uint(spaceAtEnd))
	//println("mask at end", mask.Int64(), BigIntToBinary(mask))

	return mask
}

// RectSetOn sets a given rectangle of lights to the On position, using binary or
func (g *LightGridInt) RectSetOn(x1, y1, x2, y2 int) {
	mask := g.MaskFromRectangle(x1, y1, x2, y2)
	g.I = g.I.Or(g.I, mask)
}

// RectSetOff sets a given rectangle of lights to the Off position, using binary and-not
func (g *LightGridInt) RectSetOff(x1, y1, x2, y2 int) {
	mask := g.MaskFromRectangle(x1, y1, x2, y2)
	//negMask := mask.Not(mask)
	//println("negmask", BigIntToBinary(negMask))
	g.I = g.I.AndNot(g.I, mask)
}

// RectToggle toggles the on/off of a given rectangle of lights, using binary xor
func (g *LightGridInt) RectToggle(x1, y1, x2, y2 int) {
	mask := g.MaskFromRectangle(x1, y1, x2, y2)
	g.I = g.I.Xor(g.I, mask)
}

// ToBinary returns a string of 0/1 digits, of length
// size*size
func (g *LightGridInt) ToBinary() (s string) {
	totalSize := g.Size * g.Size
	padLen := totalSize - g.I.BitLen()
	s = strings.Repeat("0", padLen)
	//println("padLen, bitlen", padLen, g.I.BitLen())
	s += BigIntToBinary(g.I)
	return s
}

// PopCount (population count aka Hamming Weight) returns the number of 1 bits in the number.
// Guaranteed to be <= size*size.
func (g *LightGridInt) PopCount() (s int) {
	for _, i := range g.I.Bits() {
		s += PopCountInt(uintptr(i))
	}
	return
}

// ToBinarySquare is the same as ToBinary but
// returns the results in the form of a square.  0,0 is top-left.
func (g *LightGridInt) ToBinarySquare() (s string) {
	bin := g.ToBinary()
	//_ = "breakpoint"
	//println("len", len(bin))
	for i := 0; i < g.Size; i++ {
		//println("slice", i*g.Size, ":", i*g.Size+g.Size)
		s += bin[i*g.Size:i*g.Size+g.Size] + "\n"
	}
	return s
}

// PopCountInt returns the number of binary 1's in the integer x
func PopCountInt(x uintptr) (count int) {
	for count = 0; x > 0; count++ {
		x &= x - 1
	}
	return count
}

// BigIntToBinary converts a big.Int to a string of 0/1 binary digits
func BigIntToBinary(in *big.Int) (s string) {
	var b bytes.Buffer
	// this is a horrible, slow way to do this, but I don't know if it matters
	for i := in.BitLen() - 1; i >= 0; i-- {
		b.WriteString(strconv.Itoa(int(in.Bit(i))))
	}
	/*for _, n := range in.Bits() {
		s += fmt.Sprintf("%0.64b", n)
	}*/
	return b.String()
}

/*
// LightGrid represents a grid of lights (1000x1000) and accepts
// instructions for toggling them
type LightGrid struct {
	Grid [1000][1000]bool
}

// NewLightGrid creates a new LightGrid with all lights turned off
func NewLightGrid() (g *LightGrid) {
	g = &LightGrid{}
	return
}

// LightsExec TBD
func LightsExec(s string) int {
	return 0
}
*/
