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
