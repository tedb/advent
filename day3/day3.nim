proc treesForSlope(x, y: int): int =
    var posX: int
    var posY: int

    for line in open("day3/input").lines:
        if posY mod y != 0:
            posY+=1
            continue
        if line[posX mod len(line)] == '#':
            result+=1
        posX+=x
        posY+=1

echo treesForSlope(3,1)
echo treesForSlope(1,1)*treesForSlope(3,1)*treesForSlope(5,1)*treesForSlope(7,1)*treesForSlope(1,2)