import strutils, tables, parseutils, sequtils

# proc echoTree(n: Node, level: int = 0) =
#   echo indent($n.label, level, "  ")

type NodeTable = Table[int, seq[int]]

proc newTableFromInput(input: string): NodeTable =
  result = initTable[int, seq[int]]()
  # e.g.: '2 <-> 0, 3, 4'
  for line in input.splitLines:
    let kv = line.split(" <-> ")
    result[kv[0].parseInt] = kv[1].split(", ").map(parseInt)

# Recursivly traverse the children of the table and count all the nodes, depth-first
proc countNodes(t: NodeTable, label: int = 0, found: var seq[int]): int =
  #echo "got $# (c: $#) -> $# (found: $#)".format(label, t[label].len, t[label], found)
  if t[label].len > 0 and not found.contains(label):
    #echo "found $#".format(label)
    found.add label
    inc result
    for child in t[label]:
      result += t.countNodes(child, found)

proc day12DigitalPlumberA*(input: string): string =
  var table = newTableFromInput(input)
  # for k, v in table:
  #   echo k, " ", v

  var found: seq[int] = @[]
  $ table.countNodes(0, found)

proc day12DigitalPlumberB*(input: string): string =
  ""

when isMainModule:
  discard
