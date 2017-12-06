import strutils, parseutils, sequtils

proc boringLoopJump(input: seq[int], partB: bool = false): int =
  var v = input
  var i = 0
  while i < v.len and i >= 0:
    inc result

    var incr = 1
    if partB and v[i] >= 3:
      incr = - 1

    v[i] += incr
    i += v[i] - incr # jmp by the non-incremented amount

proc day5MazeTwistyTrampolinesA*(input: string): string =
  $ boringLoopJump input.splitLines.map(parseInt)

proc day5MazeTwistyTrampolinesB*(input: string): string =
  $ boringLoopJump(input.splitLines.map(parseInt), true)
