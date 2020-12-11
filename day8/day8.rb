Instr = Struct.new(:op, :arg, :seen) do
    def flip_nop_jmp!
        if self.op == :nop
            self.op = :jmp
        elsif self.op == :jmp
            self.op = :nop
        end
    end
end

rawProg = []
ARGF.read.scan(/(...) (\S+)/) { |op, arg|
    rawProg << Instr.new(op.to_sym, arg.to_i, false)
}

def runProg(prog, runToEnd=false)
    acc = 0
    ip = 0
    prog.each{|i| i.seen = false}

    while true do
        return acc if prog[ip].nil?
        
        break if prog[ip].seen
        prog[ip].seen = true

        case prog[ip].op
        when :nop
            ip += 1
        when :acc
            acc += prog[ip].arg
            ip += 1
        when :jmp
            ip += prog[ip].arg
        else
            raise ArgumentError
        end
    end

    return runToEnd ? 0 : acc
end

puts "part1", runProg(rawProg)

rawProg.each do |i|
    i.flip_nop_jmp!
    acc = runProg(rawProg, true)
    i.flip_nop_jmp!
    
    if acc != 0
        puts "part 2", acc
        break
    end
end