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
	p.Nodes = make(map[string]*Node)
	return p
}

type Node struct {
	P *Processor
	// Value previously calculated
	cacheval uint16
	// Function to execute against Lref and Rref to determine Val
	F func(uint16, uint16) uint16
	// just for printability
	Op, Original string
	// References to other nodes, by name
	Lref, Rref string
	// If one of the inputs is a numeric literal, store it here
	// e.g. for "x LSHIFT 2", Rval would be 2, and Rref nil
	Lval, Rval uint16
}

// Recursively traverse nodes to determine the value of the given node
func (n *Node) Value(depth int) (v uint16) {
	if depth > 100 {
		fmt.Println("Too deep for", n)
		return
	}
	if n.cacheval != 0 {
		return n.cacheval
	}
	fmt.Printf("(%d) %s\n", depth, n)
	var l, r uint16
	if n.Lref != "" {
		fmt.Println(n.Op, "finding value for l =", n.Lref)
		l = n.P.NodeByKey(n.Lref).Value(depth + 1)
	} else {
		fmt.Println(n.Op, "using Lval", n.Lval)
		l = n.Lval
	}

	if n.Rref != "" {
		fmt.Println(n.Op, "finding value for r =", n.Rref)
		r = n.P.NodeByKey(n.Rref).Value(depth + 1)
	} else {
		fmt.Println(n.Op, "using Rval", n.Rval)
		r = n.Rval
	}

	n.cacheval = n.F(l, r)
	return n.cacheval
}

func (n *Node) String() (s string) {
	return fmt.Sprintf("\"%s\" ==> %s/%d %s %s/%d", n.Original, n.Lref, n.Lval, n.Op, n.Rref, n.Rval)
}

type Processor struct {
	Nodes map[string]*Node
}

func (p *Processor) ValueForNode(k string) (v uint16) {
	return p.Nodes[k].Value(0)
}

func (p *Processor) AddNodeByString(s string) {
	n := &Node{P: p, Original: s}
	lex := strings.Fields(s)
	key := lex[len(lex)-1]

	// NOT x -> h
	if lex[0] == "NOT" {
		n.Op = lex[0]
		n.F = func(l, r uint16) (x uint16) { return ^r }
		n.Rref, n.Rval = str_or_int(lex[1])
	} else {
		n.Op = lex[1]

		switch lex[1] {

		//123 -> x
		case "->":
			n.F = func(l, r uint16) (x uint16) { return l }
			n.Lref, n.Lval = str_or_int(lex[0])

		// x AND y -> d
		case "AND":
			n.F = func(l, r uint16) (x uint16) { return l & r }
			n.Lref, n.Lval = str_or_int(lex[0])
			n.Rref, n.Rval = str_or_int(lex[2])

		// x OR y -> e
		case "OR":
			n.F = func(l, r uint16) (x uint16) { return l | r }
			n.Lref, n.Lval = str_or_int(lex[0])
			n.Rref, n.Rval = str_or_int(lex[2])

		// x LSHIFT 2 -> f
		case "LSHIFT":
			n.F = func(l, r uint16) (x uint16) { return l << r }
			n.Lref, n.Lval = str_or_int(lex[0])
			n.Rref, n.Rval = str_or_int(lex[2])

		// y RSHIFT 2 -> g
		case "RSHIFT":
			n.F = func(l, r uint16) (x uint16) { return l >> r }
			n.Lref, n.Lval = str_or_int(lex[0])
			n.Rref, n.Rval = str_or_int(lex[2])

		default:
			panic("bad lex")
		}
	}

	p.Nodes[key] = n
}

func (p *Processor) NodeByKey(s string) (n *Node) {
	return p.Nodes[s]
}

func (p *Processor) String() (s string) {
	lines := []string{}
	for key, node := range p.Nodes {
		lines = append(lines, fmt.Sprintf("%s: %d", key, node.Value(0)))
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

func str_or_int(raw string) (s string, i uint16) {
	raw_int, err := strconv.Atoi(raw)
	if err == nil {
		// it's a number
		return "", uint16(raw_int)
	} else {
		// or, it's a variable
		return raw, 0
	}
}
