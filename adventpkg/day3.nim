import strutils, parseutils, sequtils

# Level is number of rings distant from 1 at center
type Level = int

proc width(level: Level): int =
  (level * 2) + 1

proc perimeter(level: Level): int =
  level * 8

proc day3SpiralMemoryA*(input: string): string =
  let L = Level(parseInt(input))
  $ L

proc day3SpiralMemoryB*(input: string): string =
  let L = Level(parseInt(input))
  $ L
