import adventpkg/day1
import adventpkg/day2

proc dispatch*(day: string, a: bool, b: bool, input: string): string =
  var a = a
  var b = b
  if not a and not b:
    a = true
    b = true

  case day
  of "1":
    if a:
      result = day1.day1InverseCaptchaA(input)
    if b:
      result = day1.day1InverseCaptchaB(input)
  else:
    echo "Day", day, "is not implemented"

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

  echo dispatch($ args["<day>"], args["a"], args["b"], v)
