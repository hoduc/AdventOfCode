#[

--- Day 13: Point of Incidence ---

With your help, the hot springs team locates an appropriate spring which launches you neatly and precisely up to the edge of Lava Island.

There's just one problem: you don't see any lava.

You do see a lot of ash and igneous rock; there are even what look like gray mountains scattered around. After a while, you make your way to a nearby cluster of mountains only to discover that the valley between them is completely full of large mirrors. Most of the mirrors seem to be aligned in a consistent way; perhaps you should head in that direction?

As you move through the valley of mirrors, you find that several of them have fallen from the large metal frames keeping them in place. The mirrors are extremely flat and shiny, and many of the fallen mirrors have lodged into the ash at strange angles. Because the terrain is all one color, it's hard to tell where it's safe to walk or where you're about to run into a mirror.

You note down the patterns of ash (.) and rocks (#) that you see as you walk (your puzzle input); perhaps by carefully analyzing these patterns, you can figure out where the mirrors are!

For example:

#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

To find the reflection in each pattern, you need to find a perfect reflection across either a horizontal line between two rows or across a vertical line between two columns.

In the first pattern, the reflection is across a vertical line between two columns; arrows on each of the two columns point at the line between the columns:

123456789
    ><   
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.
    ><   
123456789

In this pattern, the line of reflection is the vertical line between columns 5 and 6. Because the vertical line is not perfectly in the middle of the pattern, part of the pattern (column 1) has nowhere to reflect onto and can be ignored; every other column has a reflected column within the pattern and must match exactly: column 2 matches column 9, column 3 matches 8, 4 matches 7, and 5 matches 6.

The second pattern reflects across a horizontal line instead:

1 #...##..# 1
2 #....#..# 2
3 ..##..### 3
4v#####.##.v4
5^#####.##.^5
6 ..##..### 6
7 #....#..# 7

This pattern reflects across the horizontal line between rows 4 and 5. Row 1 would reflect with a hypothetical row 8, but since that's not in the pattern, row 1 doesn't need to match anything. The remaining rows match: row 2 matches row 7, row 3 matches row 6, and row 4 matches row 5.

To summarize your pattern notes, add up the number of columns to the left of each vertical line of reflection; to that, also add 100 multiplied by the number of rows above each horizontal line of reflection. In the above example, the first pattern's vertical line has 5 columns to its left and the second pattern's horizontal line has 4 rows above it, a total of 405.

Find the line of reflection in each of the patterns in your notes. What number do you get after summarizing all of your notes?

]#


proc echoValley(valley: seq[string]) =
    for v in valley:
        echo(v)

proc findReflectionVertical(valley: seq[string]): int =
    let yb = len(valley)
    let xb = len(valley[0])
    var max = -1
    for x in 0 ..< xb - 1:
        var y = 0
        while y < yb:
            var (r, l) = (x, x + 1)
            while l >= 0 and r < xb and l != r and valley[y][l] == valley[y][r]:
                l -= 1
                r += 1
            if l < 0 or r == xb:
                y += 1
            else:
                break
        if y >= yb:
            # echo "vertical:", (x + 1, x + 2)
            max = max(max, x + 1)
    return max

proc findReflectionHorizontal(valley: seq[string]): int =
    let yb = len(valley)
    let xb = len(valley[0])
    var max = -1
    for y in 0 ..< yb - 1:
        var x = 0
        # echo "checking row:", y
        while x < xb:
            var (u, d) = (y, y + 1)
            while u >= 0 and d < yb and u != d and valley[u][x] == valley[d][x]:
                u -= 1
                d += 1
            if u < 0 or d == yb:
                # echo "col:", x
                x += 1
            else:
                # echo "nop!"
                break
        if x >= xb:
            # echo "horizontal:", (y + 1, y + 2)
            max = max(max, y + 1)
    return max

proc findReflection(valley: seq[string]): int =
    let v = findReflectionVertical(valley)
    let h = findReflectionHorizontal(valley)
    if v > h:
        # echo "return v:", v
        return v
    # echo "return h:", h
    return 100 * h

proc day13*(fileName: string): int =
    var sum = 0 
    var valley: seq[string]
    for line in lines(fileName):
        if len(line) > 0:
            valley.add(line)
        else:
            let r = findReflection(valley)
            sum += r
            # echo "r:", r 
            # echoValley(valley)
            # echo "-----"
            valley = @[]
    # echoValley(valley)
    let r = findReflection(valley)
    # echo "r:", r
    return sum + r
    