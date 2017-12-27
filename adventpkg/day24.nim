import sequtils, parseutils, strutils, future, sets, math

type
  Part = tuple[a, b: int]

# Convert seq of 2 items into a tuple with those items
proc xySort[T](input: seq[T]): tuple[a, b: T] =
  if input.len != 2:
    quit "xySort needs length 2"

  if input[0] < input[1]:
    result = (input[0], input[1])
  else:
    result = (input[1], input[0])

proc max[T](input: seq[T]): T =
  for x in input:
    if x > result:
      result = x

proc bridgeStrength(parts: seq[Part]): int =
  parts.map((p) => p.a + p.b).foldl(a + b)

proc maxChild(parts: seq[Part], used: var HashSet[Part], rightPort: int = 0): int =
  let usedCopy = used
  var children: seq[Part] = parts.filter((p) => not usedCopy.contains(p) and (p.a == rightPort or p.b == rightPort))
  echo "children for $#: $#".format(rightPort, children)

  if children.len == 0:
    return 0

  for c in children:
    var otherPort = if c.a == rightPort: c.b
      else: c.a

    used.incl xySort(@[rightPort, otherPort])

    var thisChild = rightPort + otherPort + parts.maxChild(used, otherPort)
    echo "thisChild ($# $#) = $#".format(rightPort, otherPort, thisChild)
    if thisChild > result:
      result = thisChild

proc day24ElectromagneticMoatA*(input: string): string =
  let parts: seq[Part] = input.splitLines.map((line) => line.split("/").map((port) => port.parseInt).xySort)
  var used = initSet[Part](nextPowerOfTwo(parts.len))

  return $ parts.maxChild(used)

proc day24ElectromagneticMoatB*(input: string): string =
  ""

when isMainModule:
  assert bridgeStrength(@[(1, 2), (2, 3)]) == 8
  assert bridgeStrength(@[(0, 0)]) == 0

  assert xySort(@[1,2]) == (1, 2)
  assert xySort(@[2, 0]) == (0, 2)

  assert max(@[1]) == 1
  assert max(@[3, 1, 2]) == 3
  assert max(@[1, 2, 3]) == 3
