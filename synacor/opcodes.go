package synacor

//import "fmt"

// halt: 0
//   stop execution and terminate the program
func (vm *VM) opHalt(instr []uint16) int {
	vm.status = "Halted"
	return 0
}

// set: 1 a b
//   set register <a> to the value of <b>
func (vm *VM) opSet(instr []uint16) int {
	a, b, _ := vm.getAbc(instr)

	vm.registers[a] = b
	//println("set reg", a, "=", vm.registers[a])
	return 3
}

// push: 2 a
//   push <a> onto the stack
func (vm *VM) opPush(instr []uint16) int {
	a := vm.get(instr[0])
	vm.stack = append(vm.stack, a)
	return 2
}

// pop: 3 a
//   remove the top element from the stack and write it into <a>; empty stack = error
func (vm *VM) opPop(instr []uint16) int {
	if len(vm.stack) == 0 {
		// error
		vm.status = "opPop has empty stack!"
		return 0
	}

	v := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	a, _, _ := vm.getAbc(instr)
	vm.registers[a] = v
	return 2
}

// eq: 4 a b c
//   set <a> to 1 if <b> is equal to <c>; set it to 0 otherwise
func (vm *VM) opEq(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	if b == c {
		vm.registers[a] = 1
	} else {
		vm.registers[a] = 0
	}
	//println("eq:", b, c, "->", a)
	return 4
}

// gt: 5 a b c
//   set <a> to 1 if <b> is greater than <c>; set it to 0 otherwise
func (vm *VM) opGt(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	if b > c {
		vm.registers[a] = 1
	} else {
		vm.registers[a] = 0
	}
	return 4
}

// jmp: 6 a
//   jump to <a>
func (vm *VM) opJmp(instr []uint16) int {
	a := vm.get(instr[0])
	return int(a)
}

// jt: 7 a b
// (Jump if True)
//   if <a> is nonzero, jump to <b>
func (vm *VM) opJt(instr []uint16) int {
	a := vm.get(instr[0])
	b := vm.get(instr[1])
	if a != 0 {
		return int(b)
	}
	return 3
}

// jf: 8 a b
// (Jump if False)
//   if <a> is zero, jump to <b>
func (vm *VM) opJf(instr []uint16) int {
	a := vm.get(instr[0])
	b := vm.get(instr[1])
	if a == 0 {
		return int(b)
	}
	return 3
}

// add: 9 a b c
//   assign into <a> the sum of <b> and <c> (modulo 32768)
func (vm *VM) opAdd(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	vm.registers[a] = (b + c) % 32768
	return 4
}

// mult: 10 a b c
//   store into <a> the product of <b> and <c> (modulo 32768)
func (vm *VM) opMult(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	vm.registers[a] = (b * c) % 32768
	return 4
}

// mod: 11 a b c
//   store into <a> the remainder of <b> divided by <c>
func (vm *VM) opMod(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	vm.registers[a] = b % c
	return 4
}

// and: 12 a b c
//   stores into <a> the bitwise and of <b> and <c>
func (vm *VM) opAnd(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	vm.registers[a] = b & c
	return 4
}

// or: 13 a b c
//   stores into <a> the bitwise or of <b> and <c>
func (vm *VM) opOr(instr []uint16) int {
	a, b, c := vm.getAbc(instr)

	vm.registers[a] = b | c
	return 4
}

// not: 14 a b
//   stores 15-bit bitwise inverse of <b> in <a>
func (vm *VM) opNot(instr []uint16) int {
	a, b, _ := vm.getAbc(instr)

	// this is tricky because we want to do the NOT only against 15 bits, not 16
	// first bit is untouched
	first := b >> 15 << 15
	last15 := b << 1
	last15Not := ^last15 >> 1
	v := first | last15Not
	vm.registers[a] = v
	return 3
}

// rmem: 15 a b
//   read memory at address <b> and write it to <a>
func (vm *VM) opRmem(instr []uint16) int {
	a, b, _ := vm.getAbc(instr)
	vm.registers[a] = vm.memory[b]
	//println("rmem: reg", a, "= mem", b, "==", vm.registers[a])
	return 3
}

// wmem: 16 a b
//   write the value from <b> into memory at address <a>
func (vm *VM) opWmem(instr []uint16) int {
	a := vm.get(instr[0])
	b := vm.get(instr[1])
	//memAddr := vm.registers[a]
	//println("wmem: memAddr", memAddr)
	vm.memory[a] = b
	return 3
}

// call: 17 a
//   write the address of the next instruction to the stack and jump to <a>
func (vm *VM) opCall(instr []uint16, nextI int) int {
	a := vm.get(instr[0])
	vm.stack = append(vm.stack, uint16(nextI)+2)
	//fmt.Printf("stack after call: %v\n", vm.stack)

	return int(a)
}

// ret: 18
//   remove the top element from the stack and jump to it; empty stack = halt
func (vm *VM) opRet(instr []uint16) int {
	if len(vm.stack) == 0 {
		// error
		vm.status = "opRet has empty stack!"
		return 0
	}

	v := vm.stack[len(vm.stack)-1]
	//println("opRet stack len:", len(vm.stack), "val", v)

	vm.stack = vm.stack[:len(vm.stack)-1]
	//fmt.Printf("stack after ret: %v\n", vm.stack)

	return int(v)
}

// out: 19 a
//   write the character represented by ascii code <a> to the terminal
func (vm *VM) opOut(instr []uint16) int {
	b := vm.get(instr[0])
	_, err := vm.w.Write([]byte{byte(b)})
	vm.w.Flush()
	if err != nil {
		vm.status = err.Error()
		return 0
	}
	return 2
}

// in: 20 a
//   read a character from the terminal and write its ascii code to <a>; it can be
//   assumed that once input starts, it will continue until a newline is encountered;
//   this means that you can safely read whole lines from the keyboard and trust that
//   they will be fully read
func (vm *VM) opIn(instr []uint16) int {
	a := instr[0] - 32768
	chr, err := vm.r.ReadByte()
	if err != nil {
		vm.status = err.Error()
		return 0
	}
	vm.registers[a] = uint16(chr)
	return 2
}

// noop: 21
//   no operation
func (vm *VM) opNoop(instr []uint16) int {
	return 1
}
