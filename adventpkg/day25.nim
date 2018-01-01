import sequtils

proc checksum(tape: seq[bool]): int =
  for _, v in tape:
    if v:
      inc result

proc turing(steps: int): int =
  type States = enum A, B, C, D, E, F
  var state: States = A
  var cursor: int = 0
  var tape: seq[bool] = @[]

  for i in 1..steps:
    if cursor < tape.low:
      tape.insert(false, 0)
      inc cursor
    elif cursor > tape.high:
      tape.add(false)

    case state
    of A:
      if not tape[cursor]:
        tape[cursor] = true
        inc cursor
        state = B
      else:
        tape[cursor] = false
        dec cursor
        state = F
    of B:
      if not tape[cursor]:
        tape[cursor] = false
        inc cursor
        state = C
      else:
        tape[cursor] = false
        inc cursor
        state = D
    of C:
      if not tape[cursor]:
        tape[cursor] = true
        dec cursor
        state = D
      else:
        tape[cursor] = true
        inc cursor
        state = E
    of D:
      if not tape[cursor]:
        tape[cursor] = false
        dec cursor
        state = E
      else:
        tape[cursor] = false
        dec cursor
        state = D
    of E:
      if not tape[cursor]:
        tape[cursor] = false
        inc cursor
        state = A
      else:
        tape[cursor] = true
        inc cursor
        state = C
    of F:
      if not tape[cursor]:
        tape[cursor] = true
        dec cursor
        state = A
      else:
        tape[cursor] = true
        inc cursor
        state = A

  checksum(tape)

if isMainModule:
  assert checksum(@[]) == 0
  assert checksum(@[false]) == 0
  assert checksum(@[true]) == 1
  assert checksum(@[true, false, true]) == 2

  echo turing(12994925)
