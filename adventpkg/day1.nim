import strutils, parseutils

type CircularDigitString = object
  digits: string
  offset: int

proc rightindexforpair(digitstring: CircularDigitString, left: int): int =
  (left + digitstring.offset) mod digitstring.digits.len

iterator pairs(digitstring: CircularDigitString): tuple[a: char, b: char] =
  for i in 0..digitstring.digits.len - 1:
    yield (digitstring.digits[i], digitstring.digits[digitstring.rightindexforpair(i)])

proc day1InverseCaptchaA*(input: string, offset: int = 1): string =
  var sum = 0
  for a, b in CircularDigitString(digits: input, offset: offset):
    if a == b:
      sum += parseInt($a)
  result = $sum

proc day1InverseCaptchaB*(input: string): string =
  day1InverseCaptchaA(input, int(input.len/2))
