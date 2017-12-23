import sequtils, parseutils, strutils, times, os

proc initDancers(length: Natural = 16): seq[char] =
  result = newSeq[char](length)
  for i, x in result.mpairs:
    x = char(int8('a') + i)

# Find the positions of 2 items in seq
proc find2[T](d: var seq[T], item1, item2: T): tuple[m, n: Natural] {.inline.}=
  for i, x in d:
    if x == item1:
      result.m = i
    elif x == item2:
      result.n = i
  if result.m == 0 and result.n == 0:
    quit "item " & $item1 & " or " & $item2 & " was not found in seq"

proc spin(s: var seq[char]; n: Natural) =
  s.insert(s[s.len-n..<s.len])
  s.setLen(s.len-n)

proc spin2(s: var seq[char]; n: Natural) =
  let prefix = s[0..<s.len-n]
  let suffix = s[s.len-n..<s.len]
  s[0..<s.len-n] = suffix
  s[s.len-n..<s.len] = prefix

proc spin3(s: var seq[char]; n: Natural) =
  let prefix = s[0..<s.len-n]
  s = s[s.len-n..<s.len]
  s.add(prefix)

proc spin4(s: var seq[char]; n: Natural) =
  let oldEnd = s.len-n
  s.add(s[0..<oldEnd])
  s.delete(0, <oldEnd)

proc spinInline(s: var seq[char]; n: Natural) {.inline.} =
  s.insert(s[s.len-n..<s.len])
  s.setLen(s.len-n)

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
  let firstDancers = initDancers(length)
  var d = firstDancers

  var cyclePeriod: int

  for i in 1..1_000_000_000:
    if i mod 100 == 0:
      echo "at $#: $#".format(i, d)

    d.dance(input)
    if d == firstDancers:
      echo "matched at $#, $# == $#".format(i, d, firstDancers)
      cyclePeriod = i
      break

  # Get the positions at the billionth iteration with the redundanat cycles eliminated
  d = firstDancers
  for i in 1..(1_000_000_000 mod cyclePeriod):
    d.dance(input)

  d.join()

# from https://stackoverflow.com/questions/36577570/how-to-benchmark-few-lines-of-code-in-nim
template benchmark(benchmarkName: string, code: untyped) =
  let t0 = epochTime()
  code
  let elapsed = epochTime() - t0
  let elapsedStr = elapsed.formatFloat(format = ffDecimal, precision = 3)
  echo "CPU Time [", benchmarkName, "] ", elapsedStr, "s"

# Results:
# CPU Time [spin] 0.679s
# CPU Time [spin2] 1.417s
# CPU Time [spin3] 1.235s
# CPU Time [spin4] 1.153s
# CPU Time [spinInline] 0.653s
proc benchmarks =
  proc benchDancers(testproc: proc(s: var seq[char]; n: Natural), name: string = "bench") =
    var d = initDancers()
    testproc(d, 2)
    assert d[0] == 'o'
    assert d[1] == 'p'
    assert d[15] == 'n'

    benchmark name:
      for i in 0..<1_000_000:
        testproc(d, 3)

  benchDancers(spin, "spin")
  benchDancers(spin2, "spin2")
  benchDancers(spin3, "spin3")
  benchDancers(spin4, "spin4")

  var d = initDancers()
  spinInline(d, 2)
  benchmark "spinInline":
    for i in 0..<1_000_000:
      spinInline(d, 3)

when isMainModule:
  let firstDancers = initDancers()
  var d = firstDancers
  assert firstDancers == d
  assert d[0] == 'a'
  assert d[15] == 'p'

  var find_result: tuple[m, n: Natural] = (Natural(0), Natural(15))
  assert find_result == d.find2('a', 'p')

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

  assert firstDancers != d

  echo "all asserts passed"

  benchmarks()
