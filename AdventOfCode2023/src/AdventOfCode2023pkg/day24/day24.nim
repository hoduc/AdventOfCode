import std/unittest
import std/math
import std/strutils
import std/sequtils
import std/sugar

type Scalar = tuple
    x, y, z: float64

type Line = object
    x0, y0: float64
    slope, yIntercept: float64

type Stone = object
    x, vx: float64
    y, vy: float64 


proc lineFromTwoPoints(x1,y1,x2,y2: float64): Line =
    var slope = ( y2 - y1 ) / (x2 - x1)
    var yIntercept = y2 - slope * x2
    return Line(slope: slope, yIntercept: yIntercept)

proc lineIntersect(l1, l2: Line): (float64, float64) =
    var slope = l1.slope - l2.slope
    var yIntercept = l2.yIntercept - l1.yIntercept
    var x = yIntercept / slope
    var y = l1.slope*x + l1.yIntercept
    return (x, y)

proc day24*(fileName: string, min, max: float64): int =
    var stones: seq[Stone]
    for line in lines(fileName):
        var lineSplit = line.split("@")
        var ps = lineSplit[0].split(",").map(e => float64(parseInt(e.strip())))
        var vs = lineSplit[1].split(",").map(e => float64(parseInt(e.strip())))
        var p: Scalar = (ps[0], ps[1], ps[2])
        var v: Scalar = (vs[0], vs[1], vs[2])
        
        var stone = Stone(x: p.x, y: p.y, vx: v.x, vy: v.y)

        # var (x2, y2) = (p.x + v.x, p.y + v.y)
        # var l = lineFromTwoPoints(p.x, p.y, x2, y2)
        # l.x0 = p.x
        # l.y0 = p.y
        # echo (p.x, p.y), " => " , (x2, y2), "|", l        
        # stoneLines.add(l)
        stones.add(stone)
    var intersections = 0
    for i in 0 ..< len(stones):
        var s1 = stones[i]
        var l1 = lineFromTwoPoints(s1.x, s1.y, s1.x + s1.vx, s1.y + s1.vy)
        for j in i+1 ..< len(stones):
            var s2 = stones[j]
            var l2 = lineFromTwoPoints(s2.x, s2.y, s2.x + s2.vx, s2.y + s2.vy)
            var (x,y) = lineIntersect(l1, l2)
            # echo "---", (x,y)
            # check if it is in the range
            # and t not negative
            if x >= min and x <= max and y >= min and y <= max and (x - s1.x) / s1.vx > 0 and (y - s1.y) / s1.vy > 0 and (x - s2.x) / s2.vx > 0 and (y - s2.y) / s2.vy > 0:
                # echo ":in:"
                intersections += 1

    return intersections




test "logic":
    check lineFromTwoPoints(2, 3, 6, 4) == Line(slope: 1 / 4, yIntercept: 5 / 2)
    var (x, y) = lineIntersect(lineFromTwoPoints(19, 13, 17, 14), lineFromTwoPoints(18, 19, 17, 18))
    check round(x, 2) == 14.33
    check round(y, 2) == 15.33

    echo lineIntersect(lineFromTwoPoints(19, 13, 17, 14), lineFromTwoPoints(20, 19, 21, 14))
    echo lineIntersect(lineFromTwoPoints(18, 19, 17, 18), lineFromTwoPoints(12, 31, 11, 29))
    echo lineIntersect(lineFromTwoPoints(12, 31, 11, 29), lineFromTwoPoints(20, 19, 21, 14))