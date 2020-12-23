deck1 = %w[8 19 46 11 36 10 35 9 24 22 50 1 34 7 18 28 3 38 43 2 6 42 23 12 20].map(&:to_i)
deck2 = %w[39 27 44 29 5 48 30 32 15 31 14 21 49 17 45 47 16 26 33 25 13 41 4 40 37].map(&:to_i)

while deck1.size > 0 && deck2.size > 0
    c1, c2 = deck1.shift, deck2.shift
    if c1 > c2
        deck1.push(c1, c2)
    else
        deck2.push(c2, c1)
    end
end

win = deck1.size > 0 ? deck1 : deck2

acc = 0
win.reverse.each_with_index do |c, i|
    acc += c*(i+1)
end
puts acc