import lists, strutils, parseutils, sequtils

proc newRing(len: int = 256): DoublyLinkedRing[int] =
  result = initDoublyLinkedRing[int]()
  for i in 0..<len:
    result.append i

echo repr newRing()

proc day10KnotHashA*(input: string): string =
  let input.split(",").map(parseInt)

let input = "102,255,99,252,200,24,219,57,103,2,226,254,1,0,69,216"
echo day10KnotHashA(input)

when isMainModule:
    assert day10KnotHashA()


 # To achieve this, begin with a list of numbers from 0 to 255, a current position which
 #begins at 0 (the first element in the list), a skip size (which starts at 0), and a
 #sequence of lengths (your puzzle input). Then, for each length:
 #
 # Reverse the order of that length of elements in the list, starting with the element
 #at the current position.
 # Move the current position forward by that length plus the skip size.
 # Increase the skip size by one.
