import strutils

type
  Pos = tuple[x, y: int]
  Facing = enum fUp, fRight, fDown, fLeft

proc forward(n: var Pos, f: Facing) =
  case f
  of fUp:
    dec n.y
  of fRight:
    inc n.x
  of fDown:
    inc n.y
  of fLeft:
    dec n.x

proc nextMove(current, n, e, s, w: char, facing: Facing): Facing =
  result = facing
  if current == '+':
    case facing
    of fUp, fDown:
      if e != ' ':
        result = fRight
      elif w != ' ':
        result = fLeft
    of fRight, fLeft:
      if n != ' ':
        result = fUp
      elif s != ' ':
        result = fDown

proc get(lines: seq[string], y, x: int): char =
  return if y < lines.low or y > lines.high or x < lines[y].low or x > lines[y].high:
    ' '
    else:
      lines[y][x]

proc followPath(lines: seq[string]): tuple[chars: string, steps: int] =
  result.chars = ""
  var pos = Pos((lines[0].find('|'), 0))
  var current = lines[pos.y][pos.x]
  var facing = fDown

  while current != ' ':
    inc result.steps

    if current in ('A'..'Z'):
      result.chars &= current
    facing = current.nextMove(lines.get(pos.y-1, pos.x), lines.get(pos.y, pos.x+1), lines.get(pos.y+1, pos.x), lines.get(pos.y, pos.x-1), facing)
    pos.forward(facing)

    current = lines.get(pos.y, pos.x)

proc day19ASeriesofTubesA*(input: string): string =
  let lines = input.splitLines
  $ lines.followPath.chars

proc day19ASeriesofTubesB*(input: string): string =
  let lines = input.splitLines
  $ lines.followPath.steps

when isMainModule:
  assert nextMove('|', ' ', ' ', '|', ' ', fDown) == fDown
  assert nextMove('|', '|', ' ', 'A', ' ', fDown) == fDown
  assert nextMove('A', '|', '-', '|', '-', fDown) == fDown
  assert nextMove('|', ' ', ' ', '|', ' ', fDown) == fDown
  assert nextMove('|', ' ', ' ', '|', ' ', fDown) == fDown
  assert nextMove('+', '|', 'B', ' ', ' ', fDown) == fRight
  assert nextMove('B', ' ', '-', '|', '+', fRight) == fRight

  var input = """     |
     |  +--+
     A  |  C
  F---|----E|--+
     |  |  |  D
     +B-+  +--+ """

  assert input.splitLines.get(-1, 1) == ' '
  assert input.splitLines.get(0, 0) == ' '
  assert input.splitLines.get(0, 5) == '|'
