package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"source.developers.google.com/p/ted-cloud/advent"
)

var filename = flag.String("f", "", "read arg from `filename`")

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	var arg string

	println(len(*filename))

	if len(*filename) == 0 {
		arg = flag.Arg(1)
	} else {
		raw, err := ioutil.ReadFile(*filename)
		if err != nil {
			panic(err)
		}
		arg = string(raw)
	}

	var r interface{}
	switch cmd {
	case "1a":
		r = advent.Advent1a_Parens(arg)
	case "1b":
		r = advent.Advent1b_ParensBasement(arg)
	default:
		println("No cmd found")
		os.Exit(1)
	}

	if len(*filename) > 0 {
		arg = fmt.Sprintf("%.20s...", arg)
	}
	fmt.Printf("Advent %s('%s') = %v\n", cmd, arg, r)
}
