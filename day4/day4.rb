data = ARGF.file.read

# data = 'MMMSXXMASM
# MSAMXMSMSA
# AMXSXMAAMM
# MSAMASMSMX
# XMASAMXAMM
# XXAMMXXAMA
# SMSMSASXSS
# SAXAMASAAA
# MAMMMXMMMM
# MXMXAXMASX'

puts 'part 1'

w = data.lines.first.size - 1
puts "width: #{w}"

# I couldn't use regex alternation as intended, because some "starting letters" (X or S) do double duty.
# TIL, with regex lookahead, even with a zero-width match, the engine won't consider the same letter twice.
# So, we break up what would have been a larger regex into multiple expressions, and just match each of them.
re = {
    straight: /(?=XMAS)/m,
    vertical: /(?=X.{#{w}}M.{#{w}}A.{#{w}}S)/m,
    right_diag: /(?=X.{#{w+1}}M.{#{w+1}}A.{#{w+1}}S)/m,
    left_diag: /(?=X.{#{w-1}}M.{#{w-1}}A.{#{w-1}}S)/m,

    straight_rev: /(?=SAMX)/m,
    vertical_rev: /(?=S.{#{w}}A.{#{w}}M.{#{w}}X)/m,
    right_diag_rev: /(?=S.{#{w+1}}A.{#{w+1}}M.{#{w+1}}X)/m,
    left_diag_rev: /(?=S.{#{w-1}}A.{#{w-1}}M.{#{w-1}}X)/m,
}

puts re.sum{|k, v| puts k, v, data.scan(v).inspect, ""; data.scan(v).length}
