import unittest, ../advent
import ../adventpkg/day1
import ../adventpkg/day2
import ../adventpkg/day3
import ../adventpkg/day4
import ../adventpkg/day5

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

suite "day2a":
  test "example":
    let input = "5 1 9 5\n7 5 3\n2 4 6 8"
    check day2.day2CorruptionChecksumA(input) == "18"

suite "day2b":
  test "example":
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

suite "day4a":
  test "example":
    let input = "aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa"
    check day4.day4HighEntropyPassphrasesA(input) == "2"

suite "day4b":
  test "example":
    let input = "abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii ooii oooi oooo\noiii ioii iioi iiio"
    check day4.day4HighEntropyPassphrasesB(input) == "3"

suite "day5a":
  test "example":
    let input = "0\n3\n0\n1\n-3"
    check day5.day5MazeTwistyTrampolinesA(input) == "5"
