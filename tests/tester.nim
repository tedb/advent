import unittest, ../advent, ../adventpkg/day1

suite "day1":
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
