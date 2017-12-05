import strutils, parseutils, sequtils

proc boringLoopJump(input: seq[int]): int =
  var v = input
  var i = 0
  while i < v.len and i >= 0:
    inc result

    inc v[i]
    i += v[i] - 1 # jmp by the non-incremented amount

    # echo "result $#: i = $#" % [$result, $i]
    # echo $v

proc day5MazeTwistyTrampolinesA*(input: string): string =
  $ boringLoopJump input.splitLines.map(parseInt)

proc day5MazeTwistyTrampolinesB*(input: string): string =
  ""
