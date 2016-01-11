package synacor

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"reflect"
	"regexp"
	"runtime"
)

// NewVM initializes a new VM and defines pointers to functions that implement
// opcodes.
func NewVM(r io.Reader, w io.Writer) (vm *VM) {
	vm = &VM{
		Status:             "Ready",
		ExtractStringsWhen: -1,
		memory:             make([]uint16, 65536),
		stack:              make([]uint16, 0),
		r:                  bufio.NewReader(r),
		w:                  bufio.NewWriter(w),
	}

	vm.opcodes = map[uint16]func([]uint16) int{
		// halt: 0
		//   stop execution and terminate the program
		0: vm.opHalt,

		// set: 1 a b
		//   set register <a> to the value of <b>
		1: vm.opSet,

		// push: 2 a
		//   push <a> onto the stack
		2: vm.opPush,

		// pop: 3 a
		//   remove the top element from the stack and write it into <a>; empty stack = error
		3: vm.opPop,

		// eq: 4 a b c
		//   set <a> to 1 if <b> is equal to <c>; set it to 0 otherwise
		4: vm.opEq,

		// gt: 5 a b c
		//   set <a> to 1 if <b> is greater than <c>; set it to 0 otherwise
		5: vm.opGt,

		// jmp: 6 a
		//   jump to <a>
		6: vm.opJmp,

		// jt: 7 a b
		//   if <a> is nonzero, jump to <b>
		7: vm.opJt,

		// jf: 8 a b
		//   if <a> is zero, jump to <b>
		8: vm.opJf,

		// add: 9 a b c
		//   assign into <a> the sum of <b> and <c> (modulo 32768)
		9: vm.opAdd,

		// mult: 10 a b c
		//   store into <a> the product of <b> and <c> (modulo 32768)
		10: vm.opMult,

		// mod: 11 a b c
		//   store into <a> the remainder of <b> divided by <c>
		11: vm.opMod,

		// and: 12 a b c
		//   stores into <a> the bitwise and of <b> and <c>
		12: vm.opAnd,

		// or: 13 a b c
		//   stores into <a> the bitwise or of <b> and <c>
		13: vm.opOr,

		// not: 14 a b
		//   stores 15-bit bitwise inverse of <b> in <a>
		14: vm.opNot,

		// rmem: 15 a b
		//   read memory at address <b> and write it to <a>
		15: vm.opRmem,

		// wmem: 16 a b
		//   write the value from <b> into memory at address <a>
		16: vm.opWmem,

		// call: 17 a
		//   write the address of the next instruction to the stack and jump to <a>
		// (opCall uses a different signature and is not invoked via the map)
		//17: vm.opCall,

		// ret: 18
		//   remove the top element from the stack and jump to it; empty stack = halt
		18: vm.opRet,

		// out: 19 a
		//   write the character represented by ascii code <a> to the terminal
		19: vm.opOut,

		// in: 20 a
		//   read a character from the terminal and write its ascii code to <a>; it can be assumed that once input starts, it will continue until a newline is encountered; this means that you can safely read whole lines from the keyboard and trust that they will be fully read
		20: vm.opIn,

		// noop: 21
		//   no operation
		21: vm.opNoop,
	}

	return vm
}

// VM implements the memory, registers, stack, and opcodes of a Synacor Challenge
// virtual machine.  https://challenge.synacor.com/
type VM struct {
	// Status is a human readable indicator of the phase of execution
	// Specific status strings are poorly defined, other than "Ready"
	Status string

	// ExtractStringsWhen defines the instruction pointer after which
	// we will call ExtractStrings.  -1 to disable (default)
	ExtractStringsWhen int

	// memory with 15-bit address space storing 16-bit values.
	// Program is loaded into approx. the first half of this.
	// This is a slice but in practice is initialized as length 65,536 (i.e. 128 kb RAM)
	memory []uint16

	// eight registers
	registers [8]uint16

	// an unbounded stack which holds individual 16-bit values
	stack []uint16

	programLen int

	opcodes map[uint16]func([]uint16) int

	r *bufio.Reader
	w *bufio.Writer
}

func (vm *VM) Load(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	bytes_per_int := int64(2)
	program := make([]uint16, stat.Size()/bytes_per_int)
	err = binary.Read(f, binary.LittleEndian, &program)
	if err != nil {
		return err
	}

	vm.programLen = len(program)
	if vm.programLen == 0 {
		return errors.New("program len == 0")
	}

	// copy the program into memory
	copy(vm.memory, program)

	return nil
}

//search through program for strings, which is 1 or more consequtive instances
// of opcode 19 (opOut)
// TODO: consider making this search all memory, not sure if necessary
func (vm *VM) ExtractStrings() (list []string) {
	var b bytes.Buffer
	var c int
	
	list = append(list, fmt.Sprintf("----\nDumping strings for programLen %d\n----", vm.programLen))
	
	for i := 0; i < vm.programLen; i++ {
		op := vm.memory[i]

		if op == 19 {
			// append this char to buffer
			if r := vm.memory[i+1]; r > 127 {
				b.WriteRune('_')
			} else {
				b.WriteRune(rune(r))
			}
			c++
		} else if b.Len() > 0 {
			//
			list = append(list, fmt.Sprintf("%d: '%s'", i, strings.TrimSuffix(b.String(), "\n")))
			b.Reset()
		}

		// opcodes have variable length
		if op == 2 || op == 3 || op == 6 || op == 17 || op == 19 || op == 21 {
			i += 1
		} else if op == 1 || op == 7 || op == 8 || op == 14 || op == 15 || op == 16 {
			i += 2
		} else if op == 4 || op == 5 || (op >= 9 && op <= 13) {
			i += 3
		}
	}
	list = append(list, fmt.Sprintf("---\n%d chars found", c))
	return list
}

// Run executes Program (stored in memory starting at position 0)
// against the internal data structures of Memory, Registers,
// and Stack, executing a member of Opcodes for each instruction.
func (vm *VM) Run() (err error) {
	// i is basically our "instruction pointer"
	for i := 0; i < vm.programLen; {
		instr := vm.memory[i]

		var offset int
		var fn func([]uint16) int
		// opCall needs to be invoked differently
		if instr == 17 {
			//fmt.Printf("\n%d %s: %v (i=%d)\n", i, "opCall", vm.memory[i:i+4], i)
			offset = vm.opCall(vm.memory[i+1:], i)
		} else {
			fn = vm.opcodes[instr]
			//fmt.Printf("\n%d %s: %v\n", i, GetFunctionName(fn), vm.memory[i:i+4])
			if fn == nil {
				vm.Status = "ERROR"
				return errors.New(fmt.Sprintf("bad function: %d", instr))
			}
			offset = fn(vm.memory[i+1:])
		}

		if offset == 0 {
			// program ends on halt or errors from ret or pop
			if vm.Status == "Ready" {
				vm.Status = "Not Implemented: " + GetFunctionName(fn)
			}
			break
		} else if x := instr; offset > 3 && (x == 6 || x == 7 || x == 8 || x == 17 || x == 18) {
			// jump instructions return the exact instruction pointer to run next,
			// or on failure, just an incremental offset
			i = offset
			//fmt.Println("jump", offset)
		} else {
			// instructions other than jump return the number of positions to skip
			i += offset
			//fmt.Println("offset", offset)
		}

		// if i == 524 {
		// 	println("\n---- cleared register read checks ----")
		// }
		// if i == 536 {
		// 	println("\n---- cleared set register checks ----")
		// }
		// if i == 612 {
		// 	println("\n---- cleared push/pop ----")
		// }
		
		if vm.ExtractStringsWhen != -1 && i >= vm.ExtractStringsWhen {
			for _, s := range vm.ExtractStrings() {
				vm.w.WriteString(s)
				vm.w.WriteString("\n")
			}
			vm.w.Flush()

			break
		}

	}
	return nil
}

// - each number is stored as a 16-bit little-endian pair (low byte, high byte)
// - numbers 0..32767 mean a literal value
// - numbers 32768..32775 instead mean registers 0..7
// - numbers 32776..65535 are invalid
func (vm *VM) get(n uint16) uint16 {
	if n <= 32767 {
		//println("get literal:", n)
		return n
	} else if n >= 32768 && n <= 32775 {
		//println("get reg:", n-32768, "=", vm.registers[n-32768])
		return vm.registers[n-32768]
	} else {
		panic("Invalid number")
	}
}

func (vm *VM) getAbc(instr []uint16) (a, b, c uint16) {
	a = instr[0] - 32768
	b = vm.get(instr[1])
	c = vm.get(instr[2])
	return
}

func GetFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	name := regexp.MustCompile("synacor.([a-zA-Z]+)").FindStringSubmatch(fullName)
	if len(name) > 1 {
		return name[1]
	} else {
		return fullName
	}
}
