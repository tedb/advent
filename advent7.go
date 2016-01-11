package advent

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Advent07Wires takes a list of instructions, in arbitrary order, and resolves
// them using a tree
func Advent07Wires(s string) (dump string) {
	p := NewProcessor()
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		p.AddNodeByString(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	return p.String()
}

// Advent07bWires is same as Advent7Wires but defines "b" before execution
func Advent07bWires(s string) (dump string) {
	p := NewProcessor()
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		p.AddNodeByString(scanner.Text())
	}

	err := scanner.Err()
	checkErr(err)

	p.AddNodeByString("956 -> b")

	return p.String()
}

// NewProcessor creates a new Processor
func NewProcessor() *Processor {
	p := &Processor{}
	p.Nodes = make(map[string]*Node)
	return p
}

// Node is an element in a tree of instructions
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

// Value recursively traverses nodes to determine the value of the given node
func (n *Node) Value(depth int) (v uint16) {
	if depth > 100 {
		fmt.Println("Too deep for", n)
		return
	}
	if n.cacheval != 0 {
		return n.cacheval
	}
	//fmt.Printf("(%d) %s\n", depth, n)
	var l, r uint16
	if n.Lref != "" {
		//fmt.Println(n.Op, "finding value for l =", n.Lref)
		l = n.P.NodeByKey(n.Lref).Value(depth + 1)
	} else {
		//fmt.Println(n.Op, "using Lval", n.Lval)
		l = n.Lval
	}

	if n.Rref != "" {
		//fmt.Println(n.Op, "finding value for r =", n.Rref)
		r = n.P.NodeByKey(n.Rref).Value(depth + 1)
	} else {
		//fmt.Println(n.Op, "using Rval", n.Rval)
		r = n.Rval
	}

	n.cacheval = n.F(l, r)
	return n.cacheval
}

// String dumps the given instruction for debugging
func (n *Node) String() (s string) {
	return fmt.Sprintf("\"%s\" ==> %s/%d %s %s/%d", n.Original, n.Lref, n.Lval, n.Op, n.Rref, n.Rval)
}

// Processor records all the Nodes and allows them to be chained together
type Processor struct {
	Nodes map[string]*Node
}

// ValueForNode calculates the value for the given node name
func (p *Processor) ValueForNode(k string) (v uint16) {
	return p.Nodes[k].Value(0)
}

// AddNodeByString adds a node from an instruction line like "x OR y -> e"
func (p *Processor) AddNodeByString(s string) {
	n := &Node{P: p, Original: s}
	lex := strings.Fields(s)
	key := lex[len(lex)-1]

	// NOT x -> h
	if lex[0] == "NOT" {
		n.Op = lex[0]
		n.F = func(l, r uint16) (x uint16) { return ^r }
		n.Rref, n.Rval = StringOrInt(lex[1])
	} else {
		n.Op = lex[1]

		switch lex[1] {

		//123 -> x
		case "->":
			n.F = func(l, r uint16) (x uint16) { return l }
			n.Lref, n.Lval = StringOrInt(lex[0])

		// x AND y -> d
		case "AND":
			n.F = func(l, r uint16) (x uint16) { return l & r }
			n.Lref, n.Lval = StringOrInt(lex[0])
			n.Rref, n.Rval = StringOrInt(lex[2])

		// x OR y -> e
		case "OR":
			n.F = func(l, r uint16) (x uint16) { return l | r }
			n.Lref, n.Lval = StringOrInt(lex[0])
			n.Rref, n.Rval = StringOrInt(lex[2])

		// x LSHIFT 2 -> f
		case "LSHIFT":
			n.F = func(l, r uint16) (x uint16) { return l << r }
			n.Lref, n.Lval = StringOrInt(lex[0])
			n.Rref, n.Rval = StringOrInt(lex[2])

		// y RSHIFT 2 -> g
		case "RSHIFT":
			n.F = func(l, r uint16) (x uint16) { return l >> r }
			n.Lref, n.Lval = StringOrInt(lex[0])
			n.Rref, n.Rval = StringOrInt(lex[2])

		default:
			panic("bad lex")
		}
	}

	p.Nodes[key] = n
}

// NodeByKey returns the Node given its key
func (p *Processor) NodeByKey(s string) (n *Node) {
	return p.Nodes[s]
}

// String dumps all keys' values for debugging or to show results
func (p *Processor) String() (s string) {
	lines := []string{}
	for key, node := range p.Nodes {
		lines = append(lines, fmt.Sprintf("%s: %d", key, node.Value(0)))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

// StringOrInt takes input like "aa" or "1" and returns whichever one it is
func StringOrInt(raw string) (s string, i uint16) {
	rawInt, err := strconv.Atoi(raw)
	if err == nil {
		// it's a number
		return "", uint16(rawInt)
	}
	// or, it's a variable
	return raw, 0
}
