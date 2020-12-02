#!/usr/local/bin/fish

# File contents: session=abcd
set cookie $(cat .cookie)

for day in (seq 1 (date +%d))
  echo $day
  set f ~/code/aoc2019/day$day/input
  if test -e $f; continue; end
  curl -o "$f" --cookie "$cookie" https://adventofcode.com/2019/day/$day/input
end
