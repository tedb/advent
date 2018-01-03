import lists, strutils, parseutils, sequtils

proc newRing(len: int = 256): DoublyLinkedRing[int] =
  result = initDoublyLinkedRing[int]()
  for i in 0..<len:
    result.append i

proc day10KnotHashA*(input: string): string =
  #let input.split(",").map(parseInt)
  discard

when isMainModule:
  #assert day10KnotHashA()
  echo repr newRing()
  let input = "102,255,99,252,200,24,219,57,103,2,226,254,1,0,69,216"
  echo day10KnotHashA(input)
