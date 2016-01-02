package main

import (
	"flag"
	"github.com/tedb/advent/synacor"
	"os"
	"bytes"
	"bufio"
)

var filename = flag.String("f", "/home/ubuntu/workspace/synacor/challenge.bin", "execute program from `filename`")

func main() {
	flag.Parse()
	
	outS := bytes.Buffer{}
	out := bufio.NewWriter(&outS)
	err, status := synacor.Exec(*filename, os.Stdin, out) // os.Stdout
	if err != nil {
		println("Error:", err.Error())
	}
	out.Flush()
	println(outS.String())
	println("Ended with status:", status)
}
