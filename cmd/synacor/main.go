package main

import (
	"flag"
	"github.com/tedb/advent/synacor"
	"os"
	//"bytes"
	//"bufio"
)

var filename = flag.String("f", "/home/ubuntu/workspace/src/github.com/tedb/advent/synacor/3rd_party/challenge.bin", "execute program from `filename`")

// todo: change to Integer, IP after which strings are extracted (after encryption)
var extract = flag.Int("e", -1, "extract strings from source (opcode 19; like Linux 'strings' utility), starting after passing instruction pointer `e` to skip past code self-decryption.  65535 to disable.")

func main() {
	flag.Parse()

	vm := synacor.NewVM(os.Stdin, os.Stdout)
	vm.ExtractStringsWhen = *extract

	err := vm.Load(*filename)
	if err != nil {
		panic(err)
	}

	err = vm.Run()
	if err != nil {
		println("Error:", err.Error())
	}
	println("Ended with status:", vm.Status)

}

// SPOILER ALERT
// Game solved with Ruby:
//   ruby -e 'puts [2,3,5,7,9].permutation.select {|x| x[0] + x[1] * x[2]**2 + x[3]**3 - x[4] == 399}'
