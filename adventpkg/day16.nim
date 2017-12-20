import sequtils, parseutils, strutils

proc initDancers(length: Natural = 16): seq[char] =
  result = newSeq[char](length)
  for i, x in result.mpairs:
    x = char(int8('a') + i)

proc find[T](d: var seq[T], item: T): Natural {.inline.}=
  for i, x in d:
    if x == item:
      return i
  quit "item " & $item & " was not found in seq"

proc spin[T](s: var seq[T]; n: Natural) {.inline.} =
  s.insert(s[s.len-n..<s.len])
  s.setLen(s.len-n)

proc exchange[T](d: var seq[T], pos1, pos2: Natural) {.inline.} =
  swap(d[pos1], d[pos2])

proc partner[T](d: var seq[T], name1, name2: char) {.inline.} =
  d.exchange(d.find(name1), d.find(name2))

proc dance(d: var seq[char], input: string) =
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

proc day16PermutationPromenadeA*(input: string, length: Natural = 16): string =
  var d = initDancers(length)
  d.dance(input)
  d.join()

proc day16PermutationPromenadeB*(input: string, length: Natural = 16): string =
  var d = initDancers(length)
  for i in 0..<1_000_000_000:
    if i mod 100 == 0:
      echo $i
    d.dance(input)

  d.join()

when isMainModule:
  var d = initDancers()
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
  assert d[0] == 'p'
  assert d[1] == 'a'
  assert d[15] == 'o'

  d.exchange(2, 5)
  assert d[0] == 'p'
  assert d[15] == 'o'
  assert d[2] == 'e'
  assert d[5] == 'b'

  d.partner('g', 'n')
  assert d[7] == 'n'
  assert d[14] == 'g'
