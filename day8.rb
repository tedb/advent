#!/usr/bin/ruby

# DSL for a line like the following, which is *almost* Ruby
# b inc 5 if a > 1
class Day8
  attr_accessor :vars, :max

  def initialize
    @vars = Hash.new(0)
    @max = 0
  end

  def inc x
    x
  end

  def dec x
    -x
  end

  def method_missing m, *args
    if !args.empty?
      vars[m] += args[0]
      @max = vars[m] if vars[m] > @max
    end
    vars[m]
  end
end

day8 = Day8.new
day8.instance_eval STDIN.read

puts day8.vars.values.max
puts day8.max
