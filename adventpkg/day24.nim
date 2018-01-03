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
  if parts.len == 0:
    return 0
  parts.map((p) => p.a + p.b).foldl(a + b)

proc formatBridge(parts: seq[Part]): string =
  parts.map((p) => "$#/$#".format(p.a, p.b)).join(" ")

proc pad(parts: seq[Part]): string =
  result = ""
  for p in parts:
    result &= " "

proc maxChild(parts: seq[Part], rightPort: int = 0, parents: seq[Part] = @[]): int =
  var children: seq[Part] = parts.filter((p) => (p.a == rightPort or p.b == rightPort) and not parents.contains(p))

  #echo "$#children for $# ($#): $#".format(pad(parents), rightPort, formatBridge(parents), children)

  if children.len == 0:
    return 0

  for c in children:
    var otherPort = if c.a == rightPort: c.b
      else: c.a

    var thisStrength = rightPort + otherPort + parts.maxChild(otherPort, parents & c)
    #echo "$#thisStrength ($# $#) = $#".format(pad(parents), rightPort, otherPort, thisStrength)
    if thisStrength > result:
      result = thisStrength

proc maxLength(parts: seq[Part], longestBridge: var seq[Part], rightPort: int = 0, parents: seq[Part] = @[]) =
  var children: seq[Part] = parts.filter((p) => (p.a == rightPort or p.b == rightPort) and not parents.contains(p))

  if children.len == 0:
    if len(parents) >= len(longestBridge) and parents.bridgeStrength > longestBridge.bridgeStrength:
      longestBridge = parents

  for c in children:
    var otherPort = if c.a == rightPort: c.b
      else: c.a

    parts.maxLength(longestBridge, otherPort, parents & c)

proc day24ElectromagneticMoatA*(input: string): string =
  let parts: seq[Part] = input.splitLines.map((line) => line.split("/").map((port) => port.parseInt).xySort)

  return $ parts.maxChild()

proc day24ElectromagneticMoatB*(input: string): string =
  let parts: seq[Part] = input.splitLines.map((line) => line.split("/").map((port) => port.parseInt).xySort)
  var longestBridge: seq[Part] = @[]
  parts.maxLength(longestBridge)
  return $longestBridge.bridgeStrength

when isMainModule:
  assert bridgeStrength(@[(1, 2), (2, 3)]) == 8
  assert bridgeStrength(@[(0, 0)]) == 0

  assert xySort(@[1,2]) == (1, 2)
  assert xySort(@[2, 0]) == (0, 2)

  assert max(@[1]) == 1
  assert max(@[3, 1, 2]) == 3
  assert max(@[1, 2, 3]) == 3

  assert pad(@[]) == ""
  assert pad(@[(1,2)]) == " "
  assert pad(@[(1, 2), (2, 3)]) == "  "
