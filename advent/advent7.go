// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

// Advent4_Mining brute forces MD5 hashes to
func Advent7_Wires(s string) (dump string) {
	p := NewProcessor()
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		p.Run(scanner.Text())
	}

	err := scanner.Err()
	check_err(err)

	return p.String()
}

func NewProcessor() *Processor {
	p := &Processor{}
	p.Reg = make(map[string]uint16)
	return p
}

type Processor struct {
	Reg map[string]uint16
}

// Take a command in string form and run it
// assign, AND, OR, LSHIFT, RSHIFT, NOT
func (p *Processor) Run(s string) {
	lex := strings.Fields(s)
	var val uint16

	// NOT x -> h
	if lex[0] == "NOT" {
		val = ^p.Reg[lex[1]]
	} else {
		switch lex[1] {
		//123 -> x
		case "->":
		source, err := strconv.Atoi(lex[0])
		// it's a number
		if err == nil {
			val = uint16(source)
		// it's a variable
		} else {
			val = p.Reg[lex[0]]
		}

		// x AND y -> d
		case "AND":
			val = p.Reg[lex[0]] & p.Reg[lex[2]]

		// x OR y -> e
		case "OR":
			val = p.Reg[lex[0]] | p.Reg[lex[2]]

		// x LSHIFT 2 -> f
		case "LSHIFT":
			val = p.Reg[lex[0]] << atoi(lex[2])

		// y RSHIFT 2 -> g
		case "RSHIFT":
			val = p.Reg[lex[0]] >> atoi(lex[2])

		default:
			panic("bad lex")
		}
	}

	key := lex[len(lex)-1]
	p.Reg[key] = val
}

func (p *Processor) String() (s string) {
	lines := []string{}
	for key, value := range p.Reg {
		lines = append(lines, fmt.Sprintf("%s: %d", key, value))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func atoi(s string) (i uint16) {
	i2, _ := strconv.Atoi(s)
	i = uint16(i2)
	return
}
