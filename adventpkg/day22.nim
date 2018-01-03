import sequtils, sets, strutils, math, tables

type
  Node = tuple[x, y: int]
  Facing = enum fUp, fRight, fDown, fLeft
  NodeState = enum nsWeakened, nsInfected, nsFlagged

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

proc reverse(f: Facing): Facing =
  return case f
  of fUp:
    fDown
  of fRight:
    fLeft
  of fDown:
    fUp
  of fLeft:
    fRight

proc forward(n: var Node, f: Facing) =
  case f
  of fUp:
    dec n.y
  of fRight:
    inc n.x
  of fDown:
    inc n.y
  of fLeft:
    dec n.x

# traverse the grid, starting at the given node, for the given number of activity bursts
# return number of bursts that caused an infection
proc traverseGrid(grid: HashSet[Node], node: Node, bursts: int, direction: Facing = fUp): int =
  # shadow all the arg variables
  var grid = grid
  var node = node
  var bursts = bursts
  var direction = direction

  while bursts > 0:
    var isInfected = grid.contains(node)

    #echo "pos: $#, infected: $#, direction: $#".format(node, isInfected, direction)

    direction = direction.turn(isInfected)
    if not isInfected:
      grid.incl node
      inc result
    else:
      grid.excl node

    node.forward(direction)
    dec bursts

proc traverseGridWeaken(infected: HashSet[Node], node: Node, bursts: int, direction: Facing = fUp): int =
  var node = node
  var bursts = bursts
  var direction = direction

  # States holds non-clean node states; absent nodes are "clean"
  var states = initTable[Node, NodeState]()
  for n in infected:
    states[n] = nsInfected

  while bursts > 0:
    if not states.hasKey(node):
      # node is clean
      direction = direction.turn(false)
      states[node] = nsWeakened
    else:
      case states[node]
      of nsWeakened:
        # no direction change
        states[node] = nsInfected
        inc result
      of nsInfected:
        direction = direction.turn(true)
        states[node] = nsFlagged
      of nsFlagged:
        direction = direction.reverse()
        states.del(node)

    #echo "pos: $#, infected: $#, direction: $#".format(node, isInfected, direction)

    node.forward(direction)
    dec bursts

proc day22SporificaVirusA*(input: string, bursts: int = 10000): string =
  $ mapToInfectedGrid(input).traverseGrid(mapInitialPosition(input), bursts)

proc day22SporificaVirusB*(input: string, bursts: int = 10000000): string =
  $ mapToInfectedGrid(input).traverseGridWeaken(mapInitialPosition(input), bursts)

when isMainModule:
  var sample = "..#\n#..\n..."
  var grid = mapToInfectedGrid(sample)
  assert grid.len == 2
  assert grid.contains((2, 0))
  assert grid.contains((0, 1))

  var pos = mapInitialPosition(sample)
  assert pos.x == 1
  assert pos.y == 1

  assert fUp.turn(true) == fRight
  assert fLeft.turn(true) == fUp
  assert fUp.turn(false) == fLeft

  var node: Node = (1, 1)
  node.forward(fUp)
  assert node == Node((1, 0))
  node = (1, 1)
  node.forward(fRight)
  assert node == Node((2, 1))
  node = (1, 1)
  node.forward(fDown)
  assert node == Node((1, 2))
  node = (1, 1)
  node.forward(fLeft)
  assert node == Node((0, 1))
