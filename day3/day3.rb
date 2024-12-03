require 'bundler/inline'

gemfile do
  source 'https://rubygems.org'
  gem "scanf", "~> 1.0"
  gem 'pry', '~> 0.15.0'
end

require 'scanf'
puts "part 1"
# puts ARGF.file.scanf("%d %d") { |a, b| [a, b]}.transpose.map(&:sort!).transpose.map{|a, b| (a-b).abs}.sum
puts ARGF.file.read.scan(/mul\((\d+),(\d+)\)/).to_a.map{|a, b| a.to_i*b.to_i}.sum

puts "part 2"
ARGF.rewind
data = ARGF.file.read

# data = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
re = %r{
  (don't\(\))|
  (do\(\))|
  mul\((\d+),(\d+)\)
}xm

doing = true
sum = 0
data.scan(re) do |dont_tok, do_tok, a, b|
  # puts [dont_tok, do_tok, a, b, doing].inspect

  if dont_tok
    doing = false
  elsif do_tok
    doing = true
  elsif doing
    sum += a.to_i * b.to_i
  end
end
puts sum