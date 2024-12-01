#!/usr/bin/env fish

# File contents: session=abcd
set cookie (cat .cookie)
set year 2024

for day in (seq 1 (date +%d))
  echo $day
  set f ./day$day/input
  if test -e $f; continue; end
  curl -o "$f" --cookie "$cookie" https://adventofcode.com/$year/day/$day/input
end
