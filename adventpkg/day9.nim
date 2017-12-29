proc measureStream(input: string): tuple[totalScore, countGarbage: int] =
  var i, currentScore: int
  var inGarbage: bool

  while i < input.len:
    if not inGarbage:
      case input[i]
      of '{':
        currentScore.inc
        result.totalScore.inc currentScore
      of '}':
        currentScore.dec
      of '<':
        inGarbage = true
      else:
        discard
    else:
      case input[i]
      of '!':
        i.inc
      of '>':
        inGarbage = false
      else:
        result.countGarbage.inc

    i.inc

proc day9StreamProcessingA*(input: string): string =
  $ measureStream(input).totalScore

proc day9StreamProcessingB*(input: string): string =
  $ measureStream(input).countGarbage

when isMainModule:
  assert day9StreamProcessingA("{}") == "1"
  assert day9StreamProcessingA("{{{}}}") == "6"
  assert day9StreamProcessingA("{{},{}}") == "5"
  assert day9StreamProcessingA("{{{},{},{{}}}}") == "16"
  assert day9StreamProcessingA("{<a>,<a>,<a>,<a>}") == "1"
  assert day9StreamProcessingA("{{<ab>},{<ab>},{<ab>},{<ab>}}") == "9"
  assert day9StreamProcessingA("{{<!!>},{<!!>},{<!!>},{<!!>}}") == "9"
  assert day9StreamProcessingA("{{<a!>},{<a!>},{<a!>},{<ab>}}") == "3"

  assert day9StreamProcessingB("<>") == "0"
  assert day9StreamProcessingB("<random characters>") == "17"
  assert day9StreamProcessingB("<<<<>") == "3"
  assert day9StreamProcessingB("<{!>}>") == "2"
  assert day9StreamProcessingB("<!!>") == "0"
  assert day9StreamProcessingB("<!!!>>") == "0"
  assert day9StreamProcessingB("<{o\"i!a,<{i<a>") == "10"
