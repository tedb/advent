import unittest, ../advent, ../adventpkg/day1

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
