require 'bundler/inline'

gemfile do
  source 'https://rubygems.org'
  gem "scanf", "~> 1.0"
  gem 'pry', '~> 0.15.0'
end

require 'scanf'
puts "part 1"
# puts ARGF.file.scanf("%d %d") { |a, b| [a, b]}.transpose.map(&:sort!).transpose.map{|a, b| (a-b).abs}.sum
data = ARGF.file.read

# data = '7 6 4 2 1
# 1 2 7 8 9
# 9 7 6 2 1
# 1 3 2 4 5
# 8 6 4 4 1
# 1 3 6 7 9'

def is_safe(nums, up=true)
    nums.each_cons(2) do |a, b|
        diff = up ? b-a : a-b
        return false unless diff.between?(1, 3)
    end
end

puts data.lines.map{ |line| line.scan(/\d+/).map(&:to_i) }.count{|nums| is_safe(nums, true) || is_safe(nums, false)}
# .scan(/mul\((\d+),(\d+)\)/).to_a.map{|a, b| a.to_i*b.to_i}.sum

puts "part 2"
puts data.lines.map{ |line| line.scan(/\d+/).map(&:to_i) }.count{ |nums|
    is_safe(nums, true) || is_safe(nums, false) || 
        (0..nums.size-1).any? { |i| cand = nums.dup; cand.delete_at(i); is_safe(cand, true) || is_safe(cand, false) }
}
