import strutils, parseutils, sequtils

# Level is number of rings distant from 1 at center
type Level = int

proc width(level: Level): int =
  (level * 2) + 1

proc perimeter(level: Level): int =
  if level == 0:
    return 1
  level * 8

proc last(level: Level): int =
  # x squared; math.pow requries floats
  let x = (2*level + 1)
  x * x

proc first(level: Level): int =
  if level == 0:
    return 1
  Level(level-1).last + 1


proc findLevel(n: int): Level =
  while Level(result).last < n:
    inc result

proc day3SpiralMemoryA*(input: string): string =
  let n = parseInt(input)

  let L = findLevel(n)
  echo "$# is level $# (min $#, max $#, width $#), $# from beginning." % [
    $n, $L, $L.first, $L.last, $L.width, $(n - L.first)]

  $ (L.xOffset(n) + L.yOffset(n))

proc day3SpiralMemoryB*(input: string): string =
  let L = findLevel(parseInt(input))
  $ L

when isMainModule:
  assert Level(0).width == 1
  assert Level(1).width == 3
  assert Level(2).width == 5

  assert Level(0).perimeter == 1
  assert Level(1).perimeter == 8
  assert Level(2).perimeter == 16

  assert Level(0).first == 1
  assert Level(1).first == 2
  assert Level(2).first == 10

  assert Level(0).last == 1
  assert Level(1).last == 9
  assert Level(2).last == 25

  assert findLevel(1) == 0
  assert findLevel(2) == 1
  assert findLevel(9) == 1
  assert findLevel(10) == 2
  assert findLevel(25) == 2
  assert findLevel(26) == 3
  assert findLevel(49) == 3
