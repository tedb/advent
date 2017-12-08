import nre, strutils, sets, sequtils, tables, future, algorithm

type
  Node = tuple[name: string, weight: int, children: seq[string]]

proc findUnique[T](v: seq[T]): T =
  var v = v
  v.sort(cmp)
  let last = v.len-1

  if v.len == 1:
    return v[0]
  elif v[0] == v[last]:
    return
  elif v.len == 2:
    #echo "2 elements in findUnique; choosing higher"
    return v[1]

  if v[0] != v[last]:
    if v[0] != v[1]:
      result = v[0]
      #echo "low value $# is the oddball for $#" % [$result, $v]
    else:
      result = v[last]
      #echo "high value $# is the oddball for $#" % [$result, $v]

proc selfAndChildrensWeight(n: Node, nodes: Table[string, Node]): int =
  result = n.weight

  let weights = n.children.map((c: string) => nodes[c].selfAndChildrensWeight(nodes))
  if weights.len > 0:
    let uniq = weights.findUnique
    if uniq != 0:
      echo "$#: weight: $#, children: $#, unique: $#" % [n.name, $n.weight, $weights, $uniq]
    result += weights.foldl(a+b)

# lgxjcy (359) -> vsbfmt, kbdcl
proc parseLine(s: string): Node =
  let m = s.match(re"(?<name>\w+) \((?<weight>\d+)\)(?: -> (?<children>.+))?")
  if m.isNone:
    raise newException(OSError, "line not parsable: " & s)

  let c = m.get().captures
  return Node((c["name"], parseInt(c["weight"]), c["children"].split(", ")))

proc day7RecursiveCircusA*(input: string): string =
  var names = initSet[string]()
  var children = initSet[string]()

  for node in input.splitLines.map(parseLine):
    names.incl node.name
    children.incl node.children.toSet

  var root: HashSet[string] = difference(names, children)
  if root.len != 1:
    raise newException(OSError, "not a tree; root len = " & $root.len)
  for i in root.items:
    return i

proc day7RecursiveCircusB*(input: string): string =
  #let root = day7RecursiveCircusA(input)

  var nodes = initTable[string, Node]()
  for node in input.splitLines.map(parseLine):
    nodes[node.name] = node

  let root = day7RecursiveCircusA(input)

  $ nodes[root].selfAndChildrensWeight(nodes)

when isMainModule:
  assert parseLine("lgxjcy (359) -> vsbfmt, kbdcl") == Node(("lgxjcy", 359, @["vsbfmt", "kbdcl"]))
  assert parseLine("z (1)") == Node(("z", 1, @[]))

  assert findUnique(@[1, 1, 2]) == 2
  assert findUnique(@[2, 1, 1]) == 2
  assert findUnique(@[2, 2, 1]) == 1
  assert findUnique(@[2, 1]) == 2
  assert findUnique(@[2]) == 2
  assert findUnique(@[2, 2, 2]) == 0
