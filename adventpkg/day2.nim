import strutils, parseutils, sequtils

proc findDivisibles(ints: seq[int]): int =
  for i in ints:
    for j in ints:
      if i != j and i mod j == 0:
        return i div j

proc maxDiff(ints: seq[int]): int =
  ints.max - ints.min

proc checksum(input: string, linechecker: proc): int =
  for line in splitLines(input):
    var ints = line.splitWhitespace().map(parseInt)
    result += linechecker(ints)

proc day2CorruptionChecksumA*(input: string): string =
  $ checksum(input, maxDiff)

proc day2CorruptionChecksumB*(input: string): string =
  $ checksum(input, findDivisibles)
