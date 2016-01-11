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
	case "6":
		r1, r2 = advent.Advent06Lights(arg)
	case "7":
		r1 = advent.Advent07Wires(arg)
	case "8":
		r1, r2 = advent.Advent08Matchsticks(arg)
	case "7b":
		r1 = advent.Advent07bWires(arg)
	case "9":
		r1, r2 = advent.Advent09TSP(arg)
	case "10":
		r1, r2 = advent.Advent10LookSay(arg)
	case "11":
		r1, r2 = advent.Advent11Password(arg)
	case "12":
		r1, r2 = advent.Advent12JSON(arg)
	case "13":
		r1, r2 = advent.Advent13Seating(arg)
	case "15":
		r1, r2 = advent.Advent15Ingredients(arg)
	case "17a":
		println("solution in ruby:")
		println(`list = "1 2 3 4".split.map{|x| x.to_i}`)
		println("ways = (1..list.length).reduce([]) { |acc, i| acc << list.combination(i) }.map{|e| e.select {|x| x.reduce(:+) == 150}}.flatten(1)")
		println("puts ways.count")
	case "17b":
		println("solution in ruby (17a, plus):")
		println("groups = ways.group_by{|x|x.length}")
		println("puts groups[groups.keys.min].length")
	case "18":
		r1, r2 = advent.Advent18Animation(arg, 100)
	case "20":
		r1 = advent.Advent20InfiniteElves(arg)
	case "20b":
		r1 = advent.Advent20bInfiniteElves(arg)
	case "25":
		println("solution in ruby:")
		println("offset = (r * (r+1))/2 - (r-1) + r*(c-1) + (c-1)*(c)/2")
		println("code = (1..offset-1).reduce(20151125) {|acc, n| acc = acc * 252533 % 33554393 }")
		os.Exit(0)
	default:
		println("No cmd found")
		os.Exit(1)
	}

	if len(*filename) > 0 {
		arg = fmt.Sprintf("%.20s...", arg)
	}
	fmt.Printf("Advent %s('%s') = %v, %v\n", cmd, arg, r1, r2)
}
