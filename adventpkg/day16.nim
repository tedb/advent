import sequtils, parseutils, strutils

type Dancers = seq[char]

proc initDancers(length: Natural = 16): Dancers =
  result = newSeq[char](length)
  for i, x in result.mpairs:
    x = char(int8('a') + i)

proc find[T](d: var seq[T], item: T): Natural {.raises: [IndexError].}=
  for i, x in d:
    if x == item:
      return i
  raise newException(IndexError, "item " & $item & " was not found in seq")

proc spin[T](s: var seq[T]; n: Natural) {.inline.} =
  let prefix = s[0..<s.len-n]
  let suffix = s[s.len-n..<s.len]
  s = concat(suffix, prefix)

proc exchange[T](d: var seq[T], pos1, pos2: Natural) {.inline.} =
  swap(d[pos1], d[pos2])

proc partner[T](d: var seq[T], name1, name2: char) {.inline.} =
  d.exchange(d.find(name1), d.find(name2))

proc day16PermutationPromenadeA*(input: string, length: Natural = 16): string =
  var d = initDancers(length)
  var i, n, m: Natural = 0
  while i < <input.len:
    inc(i)
    case input[i-1]:
    of 's':  # s1
      i += parseInt(input, n, i)

      d.spin(n)
    of 'x':  # x3/4
      i += parseInt(input, n, i)
      inc(i) # skip slash
      i += parseInt(input, m, i)

      d.exchange(n, m)
    of 'p':  # pe/b
      d.partner(input[i], input[i+2])
      inc(i, 3)
    else:
      quit "got invalid char " & input[i]
    if i < input.len and input[i] != ',':
      quit "next char for " & $i & " isn't comma, got " & input[i]
    inc(i)

  result = d.join()

proc day16PermutationPromenadeB*(input: string): string =
  ""

when isMainModule:
  var d = initDancers()
  echo repr(d)
  assert d[0] == 'a'
  assert d[15] == 'p'

  try:
    assert d.find('a') == 0
    assert d.find('p') == 15
  except IndexError:
    echo "didn't find letter"
    assert false

  # Make sure it raises an exception for a letter not found
  try:
    assert d.find('z') == 0
    assert false
  except IndexError:
    discard

  d.spin(1)
  echo repr(d)
  assert d[0] == 'p'
  assert d[1] == 'a'
  assert d[15] == 'o'

  d.exchange(2, 5)
  echo repr(d)
  assert d[0] == 'p'
  assert d[15] == 'o'
  assert d[2] == 'e'
  assert d[5] == 'b'

  d.partner('g', 'n')
  assert d[7] == 'n'
  assert d[14] == 'g'
