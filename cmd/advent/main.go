package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/tedb/advent"
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
		r1 = advent.Advent1a_Parens(arg)
	case "1b":
		r1 = advent.Advent1b_ParensBasement(arg)
	case "2":
		r1, r2 = advent.Advent2_Box(arg)
	case "3":
		r1, r2 = advent.Advent3_Houses(arg)
	case "4":
		r1, r2 = advent.Advent4_Mining(arg)
	case "5a":
		r1, r2 = advent.Advent5_Naughty(arg)
	case "5b":
		println(`Use Perl (Go doesn't support regex backreferences)
perl -ne 'BEGIN {$sum = 0}; $sum++ && print if /(..).*\1/ && /(.).\1/; END {print "$sum\n"}' advent5.txt`)
		os.Exit(0)
	case "7":
		r1 = advent.Advent7_Wires(arg)
	case "7b":
		r1 = advent.Advent7b_Wires(arg)
	case "12":
	println(`use shell and Perl:
egrep -o '[0-9-]+' data/advent12.txt | perl -ne '$sum+=$_; END{print $sum}'`)
	default:
		println("No cmd found")
		os.Exit(1)
	}

	if len(*filename) > 0 {
		arg = fmt.Sprintf("%.20s...", arg)
	}
	fmt.Printf("Advent %s('%s') = %v, %v\n", cmd, arg, r1, r2)
}
