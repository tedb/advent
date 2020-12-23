import strutils
import bitops

type 
    DualUint64 = tuple[L, R: uint64; LenL, LenR: Natural]
    Board = seq[DualUint64]
    
# Turn 128-byte string like "L..L..etc..1..1" into ints with bit pattern "100100" and "001001"
# This puts a margin of one on the left, and spacing at the end of the right int,
# so the ints can be pretty-printed in binary that draws the puzzle map
proc strToDualUint64(line: string; one: char): DualUint64 =
    let intLen = 64
    
    assert(len(line) < (intLen*2)-2)
    assert(len(line) > intLen)

    # index 62 is one short of full int length, so the left bit stays zero
    for ch in line[0..intLen-2]:
        result.L = result.L shl 1 or (if ch == one: 1 else: 0)

    for ch in line[intLen-1..<line.len]:
        result.R = result.R shl 1 or (if ch == one: 1 else: 0)

    result.LenL = intLen - 1
    result.LenR = line.len - result.LenL
    let unassignedBits = intLen - result.LenR
    result.R = result.R shl unassignedBits

proc `$`(d64: DualUint64): string =
    int64(d64.L).toBin(64) & int64(d64.R).toBin(64)[0..d64.LenR]

# Extract the int value of the k bits at position p from uint64 i
# Algo from http://www.geeksforgeeks.org/extract-k-bits-given-position-number
proc extractBits(i: uint64, k, p: Natural): int64 =
    let setBits = int64((1 shl k) - 1)
    setBits and int64(i shr (p - 1))

proc extractTwoBits(i: uint64; a, b: Natural): int64 =
    #TODO
    return 0

assert extractBits(1, 1, 1) == 1
assert extractBits(128, 1, 1) == 0
assert extractBits(127, 3, 5) == 7
assert extractBits(171, 5, 2) == 21

var
    seats: Board
    occ: ref Board
    occNext: ref Board
    iter: Natural

for line in open("day11/input").lines:
    seats.add strToDualUint64(line, 'L')    

let zeroRow: DualUint64 = (uint64(0), uint64(0), seats[0].LenL, seats[0].LenR)
seats.insert(zeroRow, 0)
seats.add(zeroRow)

for i, row in seats:
    if i == 0: continue
    echo i, ": ", row
    for j in 1..<row.LenL:
        var aboveBits = extractBits(seats[i-1].L, 3, j+1)
        var belowBits = extractBits(seats[i+1].L, 3, j+1)
        echo aboveBits.toBin(3)

        seats[i-1].L.bitslice(j+1, j+4)
    #for j in 0..<row.LenR:

    break

echo iter, occ[] == occNext[]
