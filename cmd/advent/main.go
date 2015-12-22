package main

import (
	"flag"
	"fmt"
	"github.com/tedb/advent"
	"io/ioutil"
	"os"
)

var filename = flag.String("f", "", "read arg from `filename`")

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	var arg string

	if len(*filename) == 0 {
		arg = flag.Arg(1)
	} else {
		raw, err := ioutil.ReadFile(*filename)
		if err != nil {
			panic(err)
		}
		arg = string(raw)
	}

	var r1, r2 interface{}
	switch cmd {
	case "1a":
		r1 = advent.Advent01aParens(arg)
	case "1b":
		r1 = advent.Advent01bParensBasement(arg)
	case "2":
		r1, r2 = advent.Advent02Box(arg)
	case "3":
		r1, r2 = advent.Advent03Houses(arg)
	case "4":
		r1, r2 = advent.Advent04Mining(arg)
	case "5a":
		r1, r2 = advent.Advent05Naughty(arg)
	case "5b":
		println(`Use Perl (Go doesn't support regex backreferences)
perl -ne 'BEGIN {$sum = 0}; $sum++ && print if /(..).*\1/ && /(.).\1/; END {print "$sum\n"}' advent5.txt`)
		os.Exit(0)
	case "7":
		r1 = advent.Advent07Wires(arg)
	case "7b":
		r1 = advent.Advent07bWires(arg)
	case "9":
		r1, r2 = advent.Advent09TSP(arg)
	case "10":
		r1, r2 = advent.Advent10LookSay(arg)
	case "12":
		r1, r2 = advent.Advent12JSON(arg)
	case "13":
		r1, r2 = advent.Advent13Seating(arg)
	case "15":
		r1, r2 = advent.Advent15Ingredients(arg)
	case "20":
		r1, r2 = advent.Advent20InfiniteElves(arg)
	default:
		println("No cmd found")
		os.Exit(1)
	}

	if len(*filename) > 0 {
		arg = fmt.Sprintf("%.20s...", arg)
	}
	fmt.Printf("Advent %s('%s') = %v, %v\n", cmd, arg, r1, r2)
}
