ints = File.read("day5/input").tr("FLBR", "0011").lines.map{|x| x.to_i(2)}.sort

puts "part1:", ints.last

# xor'ing all the ints against the list of what all the ints *should* be will reveal the missing one,
# because A xor A = 0 (a number XOR'ed against itself cancels out)
puts "part2:", ints.reduce(&:^) ^ (ints.first..ints.last).reduce(&:^)
