import strutils, parseutils, sequtils

proc day2CorruptionChecksumA*(input: string): string =
  var sum = 0

  for line in splitLines(input):
    var ints = line.splitWhitespace().map(parseInt)
    sum += ints.max - ints.min

  $ sum

proc day2CorruptionChecksumB*(input: string): string =
  ""
