package main

import (
	"flag"
	"github.com/tedb/advent/synacor"
	"os"
	//"bytes"
	//"bufio"
)

var filename = flag.String("f", "/home/ubuntu/workspace/synacor/challenge.bin", "execute program from `filename`")
var extract = flag.Bool("s", false, "extract strings from source (opcode 19), like Linux strings utility")

func main() {
	flag.Parse()
	
	program, err := synacor.Load(*filename)
	if err != nil {
		panic(err)
	}
	
	if ! *extract {
		status, err := synacor.Exec(program, os.Stdin, os.Stdout)
		if err != nil {
			println("Error:", err.Error())
		}
		println("Ended with status:", status)
	} else {
		l := synacor.ExtractStrings(program)
		for _, s := range l {
			println(s)
		}
	}
}

// SPOILER ALERT
// Game solved with Ruby:
//   ruby -e 'puts [2,3,5,7,9].permutation.select {|x| x[0] + x[1] * x[2]**2 + x[3]**3 - x[4] == 399}'