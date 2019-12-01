#!/usr/local/bin/fish

set cookie "session=53616c7465645f5fda690118b76c47ec72c69a318919bb74ed2111a7c32d0bdb58259c4d827b44e7c0bb234f1e92fe2f"

for day in (seq 1 (date +%d))
  echo $day
  set f ~/code/aoc2019/day$day/input
  if test -e $f; continue; end
  curl -o "$f" --cookie "$cookie" https://adventofcode.com/2019/day/$day/input
end
