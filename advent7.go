// Package advent implements attempts at the exercises found at
// http://adventofcode.com/.  Unit tests are in advent_test.go.
// A CLI invocation is at cmd/advent.
package advent

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Advent7_Wires takes a list of instructions, in arbitrary order, and resolves
// them using a tree
func Advent7_Wires(s string) (dump string) {
	p := NewProcessor()
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		p.AddNodeByString(scanner.Text())
	}

	err := scanner.Err()
	check_err(err)

	return p.String()
}

func NewProcessor() *Processor {
	p := &Processor{}
	p.Nodes = make(map[string]uint16)
	return p
}

type Node struct {
	P *Processor
	// Function to execute against Lref and Rref to determine Val
	F func(uint16, uint16) (uint16)
	// References to other nodes, by name
	Lref, Rref string
	// If one of the inputs is a numeric literal, store it here
	// e.g. for "x LSHIFT 2", Rval would be 2, and Rref nil
	Lval, Rval uint16
}

type Processor struct {
	Nodes map[string]*Node
}

// Recursively traverse nodes to determine the value of the given node
func (n *Node) Value() (v int) {
	var l, r int
	if n.Lref != nil {
		l = P.NodeByKey(n.Lref).Value()
	} else {
		l = Lval
	}
	
	if n.Rref != nil {
		r = P.NodeByKey(n.Rref).Value()
	} else {
		r = Rval
	}
	
	return F(l, r)
}

func (p *Processor) ValueForNode(k string) (v int) {
	return p.Nodes[k].Value()
}

func (p *Processor) AddNodeByString(s string) {
	n := &Node{p}
	lex := strings.Fields(s)
	key := lex[len(lex)-1]
	
	// NOT x -> h
	if lex[0] == "NOT" {
		n.F = func(l, r int) (x int) { return ^r }
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
	
	p.Nodes[key] = n
}

func (p *Processor) NodeByKey(s string) (n *Node) {
	return p.Nodes[k]
}

func (p *Processor) String() (s string) {
	lines := []string{}
	for key, node := range p.Nodes {
		lines = append(lines, fmt.Sprintf("%s: %d", key, node.Value()))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func atoi(s string) (i uint16) {
	i2, err := strconv.Atoi(s)
	check_err(err)
	i = uint16(i2)
	return
}

func str_or_int(raw string) (is_int bool, s string, i uint16) {
			raw_int, err := strconv.Atoi(raw)
			if err == nil {
				// it's a number
				return true, uint16(raw_int), nil
			} else {
				// or, it's a variable
				return false, 0, raw
			}
}
