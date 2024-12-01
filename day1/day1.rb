require 'bundler/inline'

gemfile do
  source 'https://rubygems.org'
  gem "scanf", "~> 1.0"
  gem 'pry', '~> 0.15.0'
end

require 'scanf'
puts "part 1"
puts ARGF.file.scanf("%d %d") { |a, b| [a, b]}.transpose.map(&:sort!).transpose.map{|a, b| (a-b).abs}.sum

puts "part 2"
ARGF.rewind
pairs = ARGF.file.scanf("%d %d") { |a, b| [a, b]}
counts = pairs.group_by{|a, b| b}.transform_values{|v| v.count}
counts.default = 0
puts pairs.inject(0) { |memo, pair| left = pair[0]; memo += left * counts[left]}

# binding.pry