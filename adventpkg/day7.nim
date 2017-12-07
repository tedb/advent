import nre, strutils, sets, sequtils, tables, future

type
  Node = tuple[name: string, weight: int, children: seq[string]]

proc selfAndChildrensWeight(n: Node, nodes: Table[string, Node]): int =
  result = n.weight

  let weights = n.children.map((c: string) => nodes[c].selfAndChildrensWeight(nodes))
  if weights.len > 0:
    if weights.min != weights.max:
      raise newException(OSError, "min $# != max $#" % [$weights.min, $weights.max])
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
