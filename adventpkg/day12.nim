import strutils, tables

type
  Node = ref NodeObj
  NodeObj = object
    label: int
    children: seq[ref Node]

proc newNode(label: int): Node =
  new(result)
  result.label = label
  result.children = @[]

proc addChild(n: Node, label: int) =
  var c: Node
  new(c)
  c.label = label
  n.children.add(c)

proc echoTree(n: Node, level: int = 0) =
  echo indent($n.label, level, "  ")

proc day12DigitalPlumberA*(input: string): string =
  # var root = initRoot(0)
  # echo repr(root)
  # root.echoTree
  ""

proc day12DigitalPlumberB*(input: string): string =
  ""

when isMainModule:
  var root = newNode(0)
  echo repr(root)
  root.echoTree
