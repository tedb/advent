import strutils, parseutils, sequtils, algorithm

proc isUnique[T](v: seq[T]): bool =
  result = true
  for i in 1..<v.len:
    if v[i] == v[i-1]:
      return false

proc anagramify(s: string): string =
  s.sorted(cmp).join

proc validPassphraseUnique(s: string): bool =
  s.splitWhitespace.sorted(cmp).isUnique

proc validPassphraseAnagram(s: string): bool =
  s.splitWhitespace.map(anagramify).sorted(cmp).isUnique

proc day4HighEntropyPassphrasesA*(input: string): string =
  $ input.splitLines.filter(validPassphraseUnique).len

proc day4HighEntropyPassphrasesB*(input: string): string =
  $ input.splitLines.filter(validPassphraseAnagram).len

when isMainModule:
  assert anagramify("cadb") == "abcd"
  assert isUnique(@["ab", "ba"])
  assert isUnique(@["ab", "ab", "cb"]) == false
