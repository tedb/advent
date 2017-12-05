import strutils
import adventpkg/day1
import adventpkg/day2
import adventpkg/day3

proc dispatch*(day: string, a: bool, b: bool, input: string): string =
  if not a and not b:
    return "%s\n%s" % [dispatch(day, true, false, input), dispatch(day, false, true, input)]

  # TODO: convert case statement to lookup table + tuples

  case day
  of "1":
    if a:
      return day1.day1InverseCaptchaA(input)
    if b:
      return day1.day1InverseCaptchaB(input)
  of "2":
    if a:
      return day2.day2CorruptionChecksumA(input)
    if b:
      return day2.day2CorruptionChecksumB(input)
  of "3":
    if a:
      return day3.day3SpiralMemoryA(input)
    if b:
      return day3.day3SpiralMemoryB(input)
  else:
    quit("Day " & day & " is not implemented")

when isMainModule:
  import docopt, tables, strutils

  let doc = """
Advent of Code 2017

Usage:
  advent <day> [a | b] [<input>]

Options:
  -h --help     Show this screen.
  -v --verison  Show version.
"""

  let args = docopt(doc, version = "advent 0.1.0")

  var v = $ args["<input>"]
  if not args["<input>"]:
    v = readAll(stdin)
    echo "read ", v.len, " bytes from stdin"
    if v.len == 0:
      quit("stdin was zero length")
  v.removeSuffix

  echo dispatch($ args["<day>"], args["a"], args["b"], v)
