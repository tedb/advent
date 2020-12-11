require 'set'

BagQuantity = Struct.new(:q, :color)

containedBy = Hash.new {|h,k| h[k] = [] }
contains = Hash.new {|h,k| h[k] = [] }


ARGF.read.scan(/(\w+ \w+) bags contain (?:no other bags|(\d+) (\w+ \w+) bags?(?:|, (\d+) (\w+ \w+) bags?(?:|, (\d+) (\w+ \w+) bags?(?:|, (\d+) (\w+ \w+) bags?(?:|, (\d+) (\w+ \w+) bags?(?:|, (\d+) (\w+ \w+) bags?(?:|, (\d+) (\w+ \w+) bags?)))))))\./) { |terms|
    outer = terms.shift
    terms.each_slice(2) { |_, inner|
        containedBy[inner] << outer if inner
    }

    terms.each_slice(2) { |q, inner|
        contains[outer] << BagQuantity.new(q.to_i, inner) if inner
    }
}

def children(hash, key)
    hash[key] | hash[key].map{ |child| children(hash, child)}.flatten
end

puts "part 1", children(containedBy, "shiny gold").to_set.length

def score(hash, key)
    hash[key].reduce(1) { |acc, child| acc + score(hash, child.color) * child.q}
end

puts "part 2", score(contains, "shiny gold") - 1