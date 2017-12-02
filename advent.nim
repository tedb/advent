import adventpkg/day1

proc dispatch*(day: string, a: bool, b: bool, input: string = nil): string =
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

  echo dispatch($ args["<day>"], args["a"], args["b"], $ args["<input>"])
