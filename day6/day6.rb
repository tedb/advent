#!env ruby

require 'byebug'
data = ARGF.read.lines.map { |line| line.split }
summed = data.transpose.sum do |col|
  op = col.pop.to_sym
  col.map(&:to_i).reduce(&op)
end

puts "Part 1", summed
