#!env ruby

require 'byebug'

lines = ARGF.read.lines
data = lines.map { |line| line.split }
summed = data.transpose.sum do |col|
  op = col.pop.to_sym
  col.map(&:to_i).reduce(&op)
end

puts "Part 1", summed

chunks = lines.map{|line| line.chomp.split('')}.transpose.slice_when {|before, _| before.join.strip == "" }
sum = chunks.map do |chunk|
  chunk.reject!{|c| c.join.strip == ""}
  op = chunk.first.last.to_sym
  chunk.map(&:join).map(&:to_i).reduce(&op)
end.sum

puts "Part 2", sum
