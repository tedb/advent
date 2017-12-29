import lists, strutils

proc insert2017(forward: int, insertions: int = 2017, valueAfter: int = 2017): int =
  var ring = initSinglyLinkedRing[int]()
  for i in 0..insertions:
    # if i mod 10000 == 0:
    #   echo i

    ring.append i
    for i in 0..<forward:
      ring.head = ring.head.next
      ring.tail = ring.tail.next

  return ring.find(valueAfter).next.value

proc day17SpinlockA*(input: string): string =
  var inputVal = 328
  $ insert2017(inputVal)

proc day17SpinlockB*(input: string): string =
  var inputVal = 328
  $ insert2017(inputVal, 50_000_000, 0)

when isMainModule:
  assert insert2017(3, 2017) == 638
