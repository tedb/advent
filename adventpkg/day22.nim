import sequtils, sets, strutils, math

type
  Node = tuple[x, y: int]
  Facing = enum fUp, fRight, fDown, fLeft

proc mapToInfectedGrid(input: string): HashSet[Node] =
  result = initSet[Node](input.len.nextPowerOfTwo)
  var x, y: int
  for line in input.splitLines:
    for i, c in line:
      if c == '#':
        result.incl((x, y))
      inc x
    inc y
    x = 0

# top left is 0, 0
proc mapInitialPosition(input: string): Node =
  let lines = input.splitLines
  result.x = lines[0].len div 2
  result.y = lines.len div 2

proc turn(f: Facing, infected: bool): Facing =
  return if infected:
    if f == fLeft: fUp
      else: Facing(int(f) + 1)
    else:
      if f == fUp: fLeft
        else: Facing(int(f) - 1)

proc forward(n: Node, f: Facing): Node =
  result = n
  case f
  of fUp:
    dec result.y
  of fRight:
    inc result.x
  of fDown:
    inc result.y
  of fLeft:
    dec result.x

# traverse the grid, starting at the given node, for the given number of activity bursts
# return number of bursts that caused an infection
proc traverseGrid(grid: var HashSet[Node], thisNode: Node, bursts: int, direction: Facing = fUp): int =
  if bursts == 0:
    return 0

  let isInfected: bool = grid.contains(thisNode)
  let newDirection = direction.turn(isInfected)

  #echo "pos: $#, infected: $#, newDirection: $#".format(thisNode, isInfected, newDirection)

  var causedInfection: int
  if not isInfected:
    grid.incl thisNode
    inc causedInfection
  else:
    grid.excl thisNode

  return causedInfection + traverseGrid(grid, thisNode.forward(newDirection), bursts-1, newDirection)

when isMainModule:
  var sample = "..#\n#..\n..."
  var grid = mapToInfectedGrid((sample))
  assert grid.len == 2
  assert grid.contains((2, 0))
  assert grid.contains((0, 1))

  var pos = mapInitialPosition(sample)
  assert pos.x == 1
  assert pos.y == 1

  var f: Facing = fUp
  assert f.turn(true) == fRight
  f = fLeft
  assert f.turn(true) == fUp
  f = fUp
  assert f.turn(false) == fLeft

  var node: Node = (1, 1)
  assert node.forward(fUp) == Node((1, 0))
  assert node.forward(fRight) == Node((2, 1))
  assert node.forward(fDown) == Node((1, 2))
  assert node.forward(fLeft) == Node((0, 1))

  grid = mapToInfectedGrid((sample))
  assert grid.traverseGrid(pos, 7) == 5
  grid = mapToInfectedGrid((sample))
  assert grid.traverseGrid(pos, 70) == 41
  grid = mapToInfectedGrid((sample))
  assert grid.traverseGrid(pos, 10000) == 5587
