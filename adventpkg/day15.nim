import strutils, sequtils, future, times, os, threadpool, nre

let factorA: uint64 = 16807
let factorB: uint64 = 48271

proc generate(initial: int, factor: uint64, chan: ptr Channel[uint64], filterMultiples: uint64 = 1) {.thread.} =
  var last = uint64(initial)
  while true:
    if chan[].peek == -1:
      break
    last = (last * factor) mod 2147483647
    if last mod filterMultiples == 0:
      chan[].send last

proc threadedResultA(starterA, starterB: int): int =
  var chan1, chan2: Channel[uint64]
  open(chan1)
  open(chan2)

  spawn generate(starterA, factorA, addr(chan1))
  spawn generate(starterB, factorB, addr(chan2))

  for i in 0..<40_000_000:
    if i mod 1000 == 0:
      echo i, " ", result
    if chan1.recv shl 48 == chan2.recv shl 48:
      inc result

  close(chan1)
  close(chan2)

proc day15DuelingGeneratorsAThreaded*(input: string): string =
  let inputs: seq[int] = input.findAll(re"(\d+)").map(parseInt)

  $ threadedResultA(inputs[0], inputs[1])

proc day15DuelingGeneratorsASimple*(input: string): string =
  let inputs: seq[int] = input.findAll(re"(\d+)").map(parseInt)
  var matches: int

  var valA = uint64(inputs[0])
  var valB = uint64(inputs[1])

  for i in 0..<40_000_000:
    valA = (valA * factorA) mod 2147483647
    valB = (valB * factorB) mod 2147483647
    if valA shl 48 == valB shl 48:
      inc matches

  $ matches

proc day15DuelingGeneratorsA*(input: string): string =
  day15DuelingGeneratorsASimple(input)

proc day15DuelingGeneratorsB*(input: string): string =
  let inputs: seq[int] = input.findAll(re"(\d+)").map(parseInt)
  var matches: int

  var chan1, chan2: Channel[uint64]
  open(chan1)
  open(chan2)

  spawn generate(inputs[0], factorA, addr(chan1), 4)
  spawn generate(inputs[1], factorB, addr(chan2), 8)

  for i in 0..<5_000_000:
    # if i mod 1000 == 0:
    #   echo i, " ", matches
    if chan1.recv shl 48 == chan2.recv shl 48:
      inc matches

  close(chan1)
  close(chan2)

  $ matches
