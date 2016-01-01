package synacor

// halt: 0
//   stop execution and terminate the program
func (vm *VM) opHalt(instr []uint16) int {
	vm.status = "Halted"
	return 0
}

// set: 1 a b
//   set register <a> to the value of <b>
func (vm *VM) opSet(instr []uint16) int {
	a := vm.get(instr[0])
	b := vm.get(instr[1])
	vm.registers[a] = b
	println("set reg", a, "=", b, ":", vm.registers[a])
	return 3
}

func (vm *VM) opPush(instr []uint16) int {
	return 0
}
func (vm *VM) opPop(instr []uint16) int {
	return 0
}
func (vm *VM) opEq(instr []uint16) int {
	return 0
}
func (vm *VM) opGt(instr []uint16) int {
	return 0
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

func (vm *VM) opAdd(instr []uint16) int {
	return 0
}
func (vm *VM) opMult(instr []uint16) int {
	return 0
}
func (vm *VM) opMod(instr []uint16) int {
	return 0
}
func (vm *VM) opAnd(instr []uint16) int {
	return 0
}
func (vm *VM) opOr(instr []uint16) int {
	return 0
}
func (vm *VM) opNot(instr []uint16) int {
	return 0
}
func (vm *VM) opRmem(instr []uint16) int {
	return 0
}
func (vm *VM) opWmem(instr []uint16) int {
	return 0
}
func (vm *VM) opCall(instr []uint16) int {
	return 0
}
func (vm *VM) opRet(instr []uint16) int {
	return 0
}

// out: 19 a
//   write the character represented by ascii code <a> to the terminal
func (vm *VM) opOut(instr []uint16) int {
	b := vm.get(instr[0])
	vm.w.Write([]byte{byte(b)})
	return 2
}

// in: 20 a
//   read a character from the terminal and write its ascii code to <a>; it can be
//   assumed that once input starts, it will continue until a newline is encountered;
//   this means that you can safely read whole lines from the keyboard and trust that
//   they will be fully read
func (vm *VM) opIn(instr []uint16) int {
	return 0
}

// noop: 21
//   no operation
func (vm *VM) opNoop(instr []uint16) int {
	return 1
}
