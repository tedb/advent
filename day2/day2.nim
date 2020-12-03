import strscans
import strutils

var part1, part2, lines: int

for line in open("day2/input").lines:
    var min, max: int
    var letter, password: string
    lines+=1

    # 1-3 a: abcde
    if scanf(line, "$i-$i $w: $w$.", min, max, letter, password):
        var appears = strutils.count(password, letter[0])
        if appears >= min and appears <= max:
            part1+=1

        if password[min-1] == letter[0] xor password[max-1] == letter[0]:
            echo "part2"
            part2+=1

echo "lines ", lines
echo part1
echo part2