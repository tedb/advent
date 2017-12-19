import sequtils

type Dancers = seq[char]

proc day16PermutationPromenadeA*(input: string): string =
  ""

proc day16PermutationPromenadeB*(input: string): string =
  ""

proc initDancers(): Dancers =
  result = newSeq[char](16)
  for i, x in result.mpairs:
    x = char(int8('a') + i)

proc find[T](d: var seq[T], item: T): Natural {.raises: [IndexError].}=
  for i, x in d:
    if x == item:
      return i
  raise newException(IndexError, "item " & $item & " was not found in seq")

proc spin[T](s: var seq[T]; n: Natural) =
  let prefix = s[0..<s.len-n]
  let suffix = s[s.len-n..<s.len]
  s = concat(suffix, prefix)

proc exchange[T](d: var seq[T], pos1, pos2: Natural) =
  # TODO: use integer swap trick with xor
  let tmp = d[pos1]
  d[pos1] = d[pos2]
  d[pos2] = tmp

proc partner[T](d: var seq[T], name1, name2: char) =
  d.exchange(d.find(name1), d.find(name2))

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
