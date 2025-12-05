#!env ruby

require 'byebug'

ranges = []
while gets && ! $_.chomp.empty?
  /^(?<start>\d+)-(?<stop>\d+)/ =~ $_
  ranges << (start.to_i .. stop.to_i)
end

# Merge ranges
ranges.sort_by!(&:first)
merged = [ranges[0]]

ranges[1..].each do |r|
  prev = merged[-1]
  if r.first > prev.last + 1
    merged << r
  else
    merged[-1] = prev.first .. [prev.last, r.last].max
  end
end

ranges = merged

count = 0
while gets
  count += 1 if ranges.detect { |r| r.include? $_.chomp.to_i }
end

puts "Part 1: #{count}"
puts "Part 2: ", ranges.sum(&:size)



