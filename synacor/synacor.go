package synacor

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	
	"runtime"
	"reflect"
	"regexp"
)

func Exec(filename string, r io.Reader, w io.Writer) (err error, status string) {
	f, err := os.Open(filename)
	if err != nil {
		return err, ""
	}

	stat, err := f.Stat()
	if err != nil {
		return err, ""
	}

	bytes_per_int := int64(2)
	program := make([]uint16, stat.Size()/bytes_per_int)
	err = binary.Read(f, binary.LittleEndian, &program)
	if err != nil {
		return err, ""
	}

	if len(program) == 0 {
		return errors.New("program len == 0"), ""
	}

	vm := NewVM(program, r, w)

	err = vm.Run()
	return err, vm.status
}

func NewVM(program []uint16, r io.Reader, w io.Writer) (vm *VM) {
	vm = &VM{program: program,
		memory: make(map[uint16]uint16),
		stack:  make([]uint16, 0),
		status: "Ready",
		r:      r,
		w:      w,
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
		17: vm.opCall,
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

type VM struct {
	program []uint16

	// memory with 15-bit address space storing 16-bit values
	memory map[uint16]uint16

	// eight registers
	registers [8]uint16

	// an unbounded stack which holds individual 16-bit values
	stack []uint16

	opcodes map[uint16]func([]uint16) int

	status string
	r      io.Reader
	w      io.Writer
}

// Run executes Program against the internal data structures of Memory, Registers,
// and Stack, executing a member of Opcodes for each instruction.
// The opcode functions return the position to jump to
func (vm *VM) Run() (err error) {
	for i := 0; i < len(vm.program); {
		instr := vm.program[i]
		fn := vm.opcodes[instr]
		fmt.Printf("\n%d %s: %v\n", i, GetFunctionName(fn), vm.program[i:i+4])
		if fn == nil {
			return errors.New(fmt.Sprintf("bad function: %d", instr))
		}
		offset := fn(vm.program[i+1:])
		if offset == 0 {
			// program ends on halt or errors from ret or pop
			break
		} else if x := instr; offset > 3 && (x == 6 || x == 7 || x == 8 || x == 17 || x == 18) {
			// jump instructions return the exact instruction to run next
			i = offset
		} else {
			// non-jump instructions return the number of positions to skip
			i += offset
			fmt.Println("offset", offset)
		}
		// 		if i > 530 {
		// 			vm.status = "Ended early for safety"
		// 			break
		// 		}
		if i == 524 {
		    println("\n---- cleared register checks ----")
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
	    println("get literal:", n)
		return n
	} else if n >= 32768 && n <= 32775 {
		println("get reg:", n-32768, "=", vm.registers[n-32768])
		return vm.registers[n-32768]
	} else {
		panic("Invalid number")
	}
}

func GetFunctionName(i interface{}) string {
    fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
    name := regexp.MustCompile("synacor.([a-zA-Z]+)").FindStringSubmatch(fullName)
    return name[1]
}