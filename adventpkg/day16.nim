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

proc find2[T](d: var seq[T], item1, item2: T): tuple[m, n: Natural] {.inline.}=
  for i, x in d:
    if x == item1:
      result.m = i
    elif x == item2:
      result.n = i

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
