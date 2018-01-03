#import nimprof
import strutils
import adventpkg/day1
import adventpkg/day2
import adventpkg/day3
import adventpkg/day4
import adventpkg/day5
import adventpkg/day6
import adventpkg/day7
import adventpkg/day8
import adventpkg/day9
import adventpkg/day10
import adventpkg/day11
import adventpkg/day12
import adventpkg/day13
import adventpkg/day14
import adventpkg/day15
import adventpkg/day16
import adventpkg/day17
import adventpkg/day18
import adventpkg/day19
import adventpkg/day20
import adventpkg/day21
import adventpkg/day22
import adventpkg/day23
import adventpkg/day24
import adventpkg/day25

proc dispatch*(day: string, input: string): string =
  if day.allCharsInSet(Digits):
    let a = dispatch(day & "a", input)
    let b = dispatch(day & "b", input)
    return a & "\n" & b

  case day
  of "1a":
    return day1InverseCaptchaA(input)
  of "1b":
    return day1InverseCaptchaB(input)
  of "2a":
    return day2CorruptionChecksumA(input)
  of "2b":
    return day2CorruptionChecksumB(input)
  of "3a":
    return day3SpiralMemoryA(input)
  of "3b":
    return day3SpiralMemoryB(input)
  of "4a":
    return day4HighEntropyPassphrasesA(input)
  of "4b":
    return day4HighEntropyPassphrasesB(input)
  of "5a":
    return day5MazeTwistyTrampolinesA(input)
  of "5b":
    return day5MazeTwistyTrampolinesB(input)
  of "7a":
    return day7RecursiveCircusA(input)
  of "7b":
    return day7RecursiveCircusB(input)
  of "9a":
    return day9StreamProcessingA(input)
  of "9b":
    return day9StreamProcessingB(input)
  of "12a":
    return day12DigitalPlumberA(input)
  of "12b":
    return day12DigitalPlumberB(input)
  of "15a":
    return day15DuelingGeneratorsA(input)
  of "15b":
    return day15DuelingGeneratorsB(input)
  of "16a":
    return day16PermutationPromenadeA(input)
  of "16b":
    return day16PermutationPromenadeB(input)
  of "17a":
    return day17SpinlockA(input)
  of "17b":
    return day17SpinlockB(input)
  of "22a":
    return day22SporificaVirusA(input)
  of "22b":
    return day22SporificaVirusB(input)
  of "24a":
    return day24ElectromagneticMoatA(input)
  of "24b":
    return day24ElectromagneticMoatB(input)
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
