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
  if level == 0:
    return 1
  (level*8) + Level(level-1).last

proc first(level: Level): int =
  if level == 0:
    return 1
  Level(level-1).last + 1

proc xOffset(level: Level, n: int): int =
  if n == 1:
    return 0
  elif n == level.first:
    return int(level) - 1
  elif n == level.last:
    return int(level)
  else:
    raise newException(OSError, "that int isn't supported yet")

proc yOffset(level: Level, n: int): int =
  if n == 1:
    return 0
  elif n == level.first:
    return int(level) - 1
  elif n == level.last:
    return int(level)
  else:
    raise newException(OSError, "that int isn't supported yet")

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

  assert Level(0).xOffset(1) == 0
  assert Level(0).yOffset(1) == 0

  assert Level(1).xOffset(8) == 0
  assert Level(1).yOffset(8) == 1

  assert Level(1).xOffset(9) == 1
  assert Level(1).yOffset(9) == 1

  assert Level(1).xOffset(7) == 1
  assert Level(1).yOffset(7) == 1

  assert Level(1).xOffset(5) == 1
  assert Level(1).yOffset(5) == 1

  assert Level(2).xOffset(10) == 2
  assert Level(2).yOffset(10) == 1

  assert Level(2).xOffset(11) == 2
  assert Level(2).yOffset(11) == 0

  assert Level(2).xOffset(13) == 2
  assert Level(2).yOffset(13) == 2

  assert Level(3).xOffset(30) == 3
  assert Level(3).yOffset(30) == 2

  assert Level(3).xOffset(37) == 3
  assert Level(3).yOffset(37) == 3


  assert findLevel(1) == 0
  assert findLevel(2) == 1
  assert findLevel(9) == 1
  assert findLevel(10) == 2
  assert findLevel(25) == 2
  assert findLevel(26) == 3
  assert findLevel(49) == 3
