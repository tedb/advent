import unittest, strutils
import ../advent
import ../adventpkg/day1
import ../adventpkg/day2
import ../adventpkg/day3
import ../adventpkg/day4
import ../adventpkg/day5
import ../adventpkg/day7
import ../adventpkg/day12
import ../adventpkg/day16
import ../adventpkg/day24

suite "day1a":
  test "1122":
    let input = "1122"
    check day1.day1InverseCaptchaA(input) == "3"
  test "1111":
    let input = "1111"
    check day1.day1InverseCaptchaA(input) == "4"
  test "1234":
    let input = "1234"
    check day1.day1InverseCaptchaA(input) == "0"
  test "91212129":
    let input = "91212129"
    check day1.day1InverseCaptchaA(input) == "9"

suite "day1b":
  test "1212":
    let input = "1212"
    check day1.day1InverseCaptchaB(input) == "6"
  test "1221":
    let input = "1221"
    check day1.day1InverseCaptchaB(input) == "0"
  test "123425":
    let input = "123425"
    check day1.day1InverseCaptchaB(input) == "4"
  test "123123":
    let input = "123123"
    check day1.day1InverseCaptchaB(input) == "12"
  test "12131415":
    let input = "12131415"
    check day1.day1InverseCaptchaB(input) == "4"

suite "day2":
  test "day2a":
    let input = "5 1 9 5\n7 5 3\n2 4 6 8"
    check day2.day2CorruptionChecksumA(input) == "18"

  test "day2b":
    let input = "5 9 2 8\n9 4 7 3\n3 8 6 5"
    check day2.day2CorruptionChecksumB(input) == "9"

suite "day3a":
  test "1":
    let input = "1"
    check day3.day3SpiralMemoryA(input) == "0"
  test "12":
    let input = "12"
    check day3.day3SpiralMemoryA(input) == "3"
  test "23":
    let input = "23"
    check day3.day3SpiralMemoryA(input) == "2"
  test "1024":
    let input = "1024"
    check day3.day3SpiralMemoryA(input) == "31"

suite "day4":
  test "day4a":
    let input = "aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa"
    check day4.day4HighEntropyPassphrasesA(input) == "2"

  test "example":
    let input = "abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii ooii oooi oooo\noiii ioii iioi iiio"
    check day4.day4HighEntropyPassphrasesB(input) == "3"

suite "day5":
  let input = "0\n3\n0\n1\n-3"

  test "day5a":
    check day5.day5MazeTwistyTrampolinesA(input) == "5"

  test "day5b":
    check day5.day5MazeTwistyTrampolinesB(input) == "10"

suite "day7":
  var input = """pbga (66)
    xhth (57)
    ebii (61)
    havc (66)
    ktlj (57)
    fwft (72) -> ktlj, cntj, xhth
    qoyq (66)
    padx (45) -> pbga, havc, qoyq
    tknk (41) -> ugml, padx, fwft
    jptl (61)
    ugml (68) -> gyxo, ebii, jptl
    gyxo (61)
    cntj (57)""".unindent

  test "day7a":
    check day7RecursiveCircusA(input) == "tknk"

  test "day7b":
    check day7RecursiveCircusB(input) == "60"

suite "day12":
  var input = """0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5""".unindent

  test "day12a":
    check day12DigitalPlumberA(input) == "6"
  test "day12b":
    check day12DigitalPlumberB(input) == "6"

suite "day16":
  var input = "s1,x3/4,pe/b"
  test "day16a":
    check day16PermutationPromenadeA(input, 5) == "baedc"
  test "day16b":
    check day16PermutationPromenadeB(input, 5) == "ceadb"

suite "day24":
  var input = """0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10"""

  test "day24a":
    check day24ElectromagneticMoatA(input) == "31"
  test "day24b":
    check day24ElectromagneticMoatB(input) == "z"
