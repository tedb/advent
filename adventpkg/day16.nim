import sequtils, parseutils, strutils

proc initDancers(length: Natural = 16): seq[char] =
  result = newSeq[char](length)
  for i, x in result.mpairs:
    x = char(int8('a') + i)

# proc find[T](d: var seq[T], item: T): Natural {.inline.}=
#   for i, x in d:
#     if x == item:
#       return i
#   quit "item " & $item & " was not found in seq"

proc find2[T](d: var seq[T], item1, item2: T): tuple[m, n: Natural] {.inline.}=
  if d[0] == item1:
    result.m = 0
  elif d[0] == item2:
    result.n = 0

  if d[1] == item1:
    result.m = 1
  elif d[1] == item2:
    result.n = 1

  if d[2] == item1:
    result.m = 2
  elif d[2] == item2:
    result.n = 2

  if d[3] == item1:
    result.m = 3
  elif d[3] == item2:
    result.n = 3

  if d[4] == item1:
    result.m = 4
  elif d[4] == item2:
    result.n = 4

  if d[5] == item1:
    result.m = 5
  elif d[5] == item2:
    result.n = 5

  if d[6] == item1:
    result.m = 6
  elif d[6] == item2:
    result.n = 6

  if d[7] == item1:
    result.m = 7
  elif d[7] == item2:
    result.n = 7

  if d[8] == item1:
    result.m = 8
  elif d[8] == item2:
    result.n = 8

  if d[9] == item1:
    result.m = 9
  elif d[9] == item2:
    result.n = 9

  if d[10] == item1:
    result.m = 10
  elif d[10] == item2:
    result.n = 10

  if d[11] == item1:
    result.m = 11
  elif d[11] == item2:
    result.n = 11

  if d[12] == item1:
    result.m = 12
  elif d[12] == item2:
    result.n = 12

  if d[13] == item1:
    result.m = 13
  elif d[13] == item2:
    result.n = 13

  if d[14] == item1:
    result.m = 14
  elif d[14] == item2:
    result.n = 14

  if d[15] == item1:
    result.m = 15
  elif d[15] == item2:
    result.n = 15

proc spin[T](s: var seq[T]; n: Natural) {.inline.} =
  #s.insert(s[s.len-n..<s.len])
  #s.setLen(s.len-n)

  # let prefix = s[0..<s.len-n]
  # s = s[s.len-n..<s.len]
  # s.add(prefix)

  let oldEnd = s.len-n
  s.add(s[0..<oldEnd])
  #echo s
  s.delete(0, <oldEnd)
  #echo s

proc exchange[T](d: var seq[T], pos1, pos2: Natural) {.inline.} =
  swap(d[pos1], d[pos2])

proc partner[T](d: var seq[T], name1, name2: T) {.inline.} =
  let (m, n) = d.find2(name1, name2)
  d.exchange(m, n)

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

  assert d.find('a') == 0
  assert d.find('p') == 15

  d.spin(1)
  assert d[0] == 'p'
  assert d[1] == 'a'
  assert d[15] == 'o'

  d.spin(2)
  assert d[0] == 'n'
  assert d[1] == 'o'
  assert d[15] == 'm'
  d.spin(13)

  d.exchange(2, 5)
  assert d[0] == 'a'
  assert d[15] == 'p'
  assert d[2] == 'f'
  assert d[5] == 'c'

  d.partner('g', 'n')
  assert d[6] == 'n'
  assert d[13] == 'g'
