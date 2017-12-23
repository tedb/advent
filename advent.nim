#import nimprof
import strutils
import adventpkg/day1
import adventpkg/day2
import adventpkg/day3
import adventpkg/day4
import adventpkg/day5
import adventpkg/day7
import adventpkg/day12
import adventpkg/day16

proc dispatch*(day: string, input: string): string =
  if day.allCharsInSet(Digits):
    let a = dispatch(day & "a", input)
    let b = dispatch(day & "b", input)
    return a & "\n" & b

  case day
  of "1a":
    return day1.day1InverseCaptchaA(input)
  of "1b":
    return day1.day1InverseCaptchaB(input)
  of "2a":
    return day2.day2CorruptionChecksumA(input)
  of "2b":
    return day2.day2CorruptionChecksumB(input)
  of "3a":
    return day3.day3SpiralMemoryA(input)
  of "3b":
    return day3.day3SpiralMemoryB(input)
  of "4a":
    return day4.day4HighEntropyPassphrasesA(input)
  of "4b":
    return day4.day4HighEntropyPassphrasesB(input)
  of "5a":
    return day5.day5MazeTwistyTrampolinesA(input)
  of "5b":
    return day5.day5MazeTwistyTrampolinesB(input)
  of "7a":
    return day7.day7RecursiveCircusA(input)
  of "7b":
    return day7.day7RecursiveCircusB(input)
  of "12a":
    return day12.day12DigitalPlumberA(input)
  of "12b":
    return day12.day12DigitalPlumberB(input)
  of "16a":
    return day16.day16PermutationPromenadeA(input)
  of "16b":
    return day16.day16PermutationPromenadeB(input)
  else:
    quit("Day " & day & " is not implemented")

when isMainModule:
  import docopt, tables, strutils

  let doc = """
Advent of Code 2017

Usage:
  advent <day> [<input>]

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

  echo dispatch($ args["<day>"], v)
