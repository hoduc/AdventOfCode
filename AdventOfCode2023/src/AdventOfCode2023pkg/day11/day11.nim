#[

--- Day 11: Cosmic Expansion ---

You continue following signs for "Hot Springs" and eventually come across an observatory. The Elf within turns out to be a researcher studying cosmic expansion using the giant telescope here.

He doesn't know anything about the missing machine parts; he's only visiting for this research project. However, he confirms that the hot springs are the next-closest area likely to have people; he'll even take you straight there once he's done with today's observation analysis.

Maybe you can help him with the analysis to speed things up?

The researcher has collected a bunch of data and compiled the data into a single giant image (your puzzle input). The image includes empty space (.) and galaxies (#). For example:

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....

The researcher is trying to figure out the sum of the lengths of the shortest path between every pair of galaxies. However, there's a catch: the universe expanded in the time it took the light from those galaxies to reach the observatory.

Due to something involving gravitational effects, only some space expands. In fact, the result is that any rows or columns that contain no galaxies should all actually be twice as big.

In the above example, three columns and two rows contain no galaxies:

   v  v  v
 ...#......
 .......#..
 #.........
>..........<
 ......#...
 .#........
 .........#
>..........<
 .......#..
 #...#.....
   ^  ^  ^

These rows and columns need to be twice as big; the result of cosmic expansion therefore looks like this:

....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......

Equipped with this expanded universe, the shortest path between every pair of galaxies can be found. It can help to assign every galaxy a unique number:

....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......

In these 9 galaxies, there are 36 pairs. Only count each pair once; order within the pair doesn't matter. For each pair, find any shortest path between the two galaxies using only steps that move up, down, left, or right exactly one . or # at a time. (The shortest path between two galaxies is allowed to pass through another galaxy.)

For example, here is one of the shortest paths between galaxies 5 and 9:

....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......

This path has length 9 because it takes a minimum of nine steps to get from galaxy 5 to galaxy 9 (the eight locations marked # plus the step onto galaxy 9 itself). Here are some other example shortest path lengths:

    Between galaxy 1 and galaxy 7: 15
    Between galaxy 3 and galaxy 6: 17
    Between galaxy 8 and galaxy 9: 5

In this example, after expanding the universe, the sum of the shortest path between all 36 pairs of galaxies is 374.

Expand the universe, then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?

]#

import std/tables
import std/strutils
import std/sequtils
import std/sugar
import std/sets

type Coord = object
    y, x: int

proc manhattanDistance(a, b: Coord): int =
    return abs(a.y - b.y) + abs(a.x - b.x)

proc day11*(fileName: string): int =
    # mapping from (y,x) to (incy, incx)
    var
        galaxies = initTable[Coord, Coord]()
        sr: seq[int] = @[]
        sc: seq[int] = @[]
        tc: seq[int] = @[]
        y = 0
        yb = 0
        xb = 0
    for line in lines(fileName):
        xb = len(line)
        if len(tc) == 0:
            tc = toSeq(0 ..< xb).map(x => 0)
        var src = 0
        for x, c in line:
            let coord = Coord(y: y, x: x)
            if c == '#':
                if coord notin galaxies:
                    galaxies[coord] = Coord(y: 0, x: 0)
            else:
                src += 1
                tc[x] += 1
        if src == xb:
            sr.add(y)
        y += 1
    yb = y
    
    for i, v in tc.pairs:
        if v == yb:
            sc.add(i)
    # echo galaxies
    # echo sr
    # echo sc

    # expand inc col
    for x in sc:
        for g, i in galaxies.mpairs:
            if g.x > x:
                i.x += 1
    # expand inc row
    for y in sr:
        for g, i in galaxies.mpairs:
            if g.y > y:
                i.y += 1
    # echo "after-expands"
    # echo galaxies

    var sumLength = 0
    let keys = galaxies.keys.toSeq()
    for i in 0 ..< len(keys) - 1:
        let
            ci = keys[i]
            ii = galaxies[ci]
            gi = Coord(y: ci.y + ii.y, x: ci.x + ii.x)
        for j in i + 1 ..< len(keys):
            let
                cj = keys[j]
                ij = galaxies[cj]
                gj = Coord(y: cj.y + ij.y, x: cj.x + ij.x)
            sumLength += manhattanDistance(gi, gj)

    return sumLength