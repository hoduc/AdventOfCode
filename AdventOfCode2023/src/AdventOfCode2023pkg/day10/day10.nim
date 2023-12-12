# #[

# --- Day 10: Pipe Maze ---

# You use the hang glider to ride the hot air from Desert Island all the way up to the floating metal island. This island is surprisingly cold and there definitely aren't any thermals to glide on, so you leave your hang glider behind.

# You wander around for a while, but you don't find any people or animals. However, you do occasionally find signposts labeled "Hot Springs" pointing in a seemingly consistent direction; maybe you can find someone at the hot springs and ask them where the desert-machine parts are made.

# The landscape here is alien; even the flowers and trees are made of metal. As you stop to admire some metal grass, you notice something metallic scurry away in your peripheral vision and jump into a big pipe! It didn't look like any animal you've ever seen; if you want a better look, you'll need to get ahead of it.

# Scanning the area, you discover that the entire field you're standing on is densely packed with pipes; it was hard to tell at first because they're the same metallic silver color as the "ground". You make a quick sketch of all of the surface pipes you can see (your puzzle input).

# The pipes are arranged in a two-dimensional grid of tiles:

#     | is a vertical pipe connecting north and south.
#     - is a horizontal pipe connecting east and west.
#     L is a 90-degree bend connecting north and east.
#     J is a 90-degree bend connecting north and west.
#     7 is a 90-degree bend connecting south and west.
#     F is a 90-degree bend connecting south and east.
#     . is ground; there is no pipe in this tile.
#     S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

# Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large, continuous loop.

# For example, here is a square loop of pipe:

# .....
# .F-7.
# .|.|.
# .L-J.
# .....

# If the animal had entered this loop in the northwest corner, the sketch would instead look like this:

# .....
# .S-7.
# .|.|.
# .L-J.
# .....

# In the above diagram, the S tile is still a 90-degree F bend: you can tell because of how the adjacent pipes connect to it.

# Unfortunately, there are also many pipes that aren't connected to the loop! This sketch shows the same loop as above:

# -L|F7
# 7S-7|
# L|7||
# -L-J|
# L|-JF

# In the above diagram, you can still figure out which pipes form the main loop: they're the ones connected to S, pipes those pipes connect to, pipes those pipes connect to, and so on. Every pipe in the main loop connects to its two neighbors (including S, which will have exactly two pipes connecting to it, and which is assumed to connect back to those two pipes).

# Here is a sketch that contains a slightly more complex main loop:

# ..F7.
# .FJ|.
# SJ.L7
# |F--J
# LJ...

# Here's the same example sketch with the extra, non-main-loop pipe tiles also shown:

# 7-F7-
# .FJ|7
# SJLL7
# |F--J
# LJ.LJ

# If you want to get out ahead of the animal, you should find the tile in the loop that is farthest from the starting position. Because the animal is in the pipe, it doesn't make sense to measure this by direct distance. Instead, you need to find the tile that would take the longest number of steps along the loop to reach from the starting point - regardless of which way around the loop the animal went.

# In the first example with the square loop:

# .....
# .S-7.
# .|.|.
# .L-J.
# .....

# You can count the distance each tile in the loop is from the starting point like this:

# .....
# .012.
# .1.3.
# .234.
# .....

# In this example, the farthest point from the start is 4 steps away.

# Here's the more complex loop again:

# ..F7.
# .FJ|.
# SJ.L7
# |F--J
# LJ...

# Here are the distances for each tile on that loop:

# ..45.
# .236.
# 01.78
# 14567
# 23...

# Find the single giant loop starting at S. How many steps along the loop does it take to get from the starting position to the point farthest from the starting position?

# ]#

import std/strutils
import std/sets
import std/heapqueue
import unittest

type Cell = tuple
    y: int
    x: int

type Job = object
    c: Cell
    priority: int

proc `<`(a, b: Job): bool = a.priority < b.priority 

type DIR = enum
    UP, DOWN, LEFT, RIGHT, NONE

proc echoMaze(maze: seq[string]) =
    for row in maze:
        echo row

proc northMatched(cy, cx, yb, xb, d: int, maze: seq[string], visited: HashSet[Cell], q: var HeapQueue[Job],matchedSymbols: varargs[char]) =
    if cy - 1 >= 0 and cy - 1 <= yb and (cy-1, cx) notin visited and matchedSymbols.contains(maze[cy-1][cx]):
        q.push(Job(c: (cy - 1, cx), priority: d + 1))

proc southMatched(cy, cx, yb, xb, d: int, maze: seq[string], visited: HashSet[Cell], q: var HeapQueue[Job], matchedSymbols: varargs[char]) =
    if cy + 1 >= 0 and cy + 1 <= yb and (cy+1, cx) notin visited and matchedSymbols.contains(maze[cy+1][cx]):
        q.push(Job(c: (cy + 1, cx), priority: d + 1))

proc westMatched(cy, cx, yb, xb, d: int, maze: seq[string], visited: HashSet[Cell], q: var HeapQueue[Job], matchedSymbols: varargs[char]) =
    if cx - 1 >= 0 and cx - 1 <= xb and (cy, cx-1) notin visited and matchedSymbols.contains(maze[cy][cx-1]):
        q.push(Job(c: (cy, cx - 1), priority: d + 1))

proc eastMatched(cy, cx, yb, xb, d: int, maze: seq[string], visited: HashSet[Cell], q: var HeapQueue[Job], matchedSymbols: varargs[char]) =
    if cx + 1 >= 0 and cx + 1 <= xb and (cy, cx+1) notin visited and matchedSymbols.contains(maze[cy][cx+1]):
        q.push(Job(c: (cy, cx + 1), priority: d + 1))

proc walk(start: Cell, maze: var seq[string]): int =
    # consider (up, down, left, right)
    let (yb, xb) = (len(maze) - 1, len(maze[0]) - 1)
    var visited = initHashSet[Cell]()
    var q = initHeapQueue[Job]()
    q.push(Job(priority: 0, c: start))
    var count = 0
    while len(q) > 0:
        echo(q)
        let
            job = q.pop()
            cell = job.c
            d = job.priority
            (cy, cx) = (cell.y, cell.x)
            cyx = maze[cy][cx]
        echo cyx
        if count > 0 and cell == start:
            echo "found loop!!!"
            break
        # maze[cell.y][cell.x] = 'x'
        visited.incl(cell)
        if cyx == '|':
            # north
            northMatched(cy, cx, yb, xb, d, maze, visited, q, '|', 'F', 'J', '7')
            # south
            southMatched(cy, cx, yb, xb, d, maze, visited, q, '|', 'L', 'J')
        elif cyx == '-':
            # west
            westMatched(cy, cx, yb, xb, d, maze, visited, q, '-', 'L', 'F')
            # east
            eastMatched(cy, cx, yb, xb, d, maze, visited, q, '-', 'J', '7')
        elif cyx == 'L':
            # north
            northMatched(cy, cx, yb, xb, d, maze, visited, q, '|', 'F', '7')
            # east
            eastMatched(cy, cx, yb, xb, d, maze, visited, q, '-', '7', 'J')
        elif cyx == 'J':
            # north
            northMatched(cy, cx, yb, xb, d, maze, visited, q, '|', 'F', '7')
            # west
            westMatched(cy, cx, yb, xb, d, maze, visited, q, '-', 'L', 'F')
        elif cyx == '7':
            # west
            westMatched(cy, cx, yb, xb, d, maze, visited, q, '-', 'L', 'F')
            # south
            southMatched(cy, cx, yb, xb, d, maze, visited, q, '|', 'L', 'J')
        elif cyx == 'F':
            # east
            eastMatched(cy, cx, yb, xb, d, maze, visited, q, '-', '7', 'J')
            # south
            southMatched(cy, cx, yb, xb, d, maze, visited, q, '|', 'L', 'J')
        count += 1
    # echo q[0]
    return 0

#[
    Since it will have to start and end at the same position
    it is implied that there are 2 connect pipes to it
    Thus need to check 2 opening
]#
proc startSymbol(west, north, east, south: char): char =
    var (up, down, left, right) = (false, false, false, false)

    if north in "|7F":
        up = true
    if east in "-7J":
        right = true
    if south in "|LJ":
        down = true
    if west in "-LF":
        left = true

    if up and down:
        return '|'
    if left and right:
        return '-'
    if up and right:
        return 'L'
    if up and left:
        return 'J'
    if down and left:
        return '7'
    return 'F'

proc nextCellDir(current: Cell, d: DIR, maze: seq[string]): (Cell, Dir) =
    var
        dir = d
        (y, x) = (current.y, current.x)

    if dir == UP:
        y -= 1
        if maze[y][x] == '7':
            dir = LEFT
        elif maze[y][x] == 'F':
            dir = RIGHT
        # | no change dir
    elif dir == RIGHT:
        x += 1
        if maze[y][x] == '7':
            dir = DOWN
        elif maze[y][x] == 'J':
            dir = UP
        # - no change dir
    elif dir == DOWN:
        y += 1
        if maze[y][x] == 'L':
            dir = RIGHT
        elif maze[y][x] == 'J':
            dir = LEFT
        # | no change dir
    else:
        x -= 1
        if maze[y][x] == 'L':
            dir = UP
        elif maze[y][x] == 'F':
            dir = DOWN
        # - no change dir
    return ((y, x), dir)

proc walk(current: var Cell, dir: var DIR, maze: var seq[string]): int =
    let goal = current
    var
        step = 0
        (current, dir) = nextCellDir(current, dir, maze)
    while current != goal:
        # echo current, dir
        step += 1
        (current, dir) = nextCellDir(current, dir, maze)
    return step div 2 + step mod 2

proc day10*(fileName: string): int =
    var start: Cell
    var maze: seq[string]
    var y = 0
    for line in lines(fileName):
        for i, c in line.pairs:
            if c == 'S':
                start = (y, i)
        maze.add(line)
        y += 1
    # echoMaze(maze)
    # echo "========"
    maze[start.y][start.x] = 'F'

    var
        w: Cell = (start.y, start.x - 1)
        n: Cell = (start.y - 1, start.x)
        e: Cell = (start.y, start.x + 1)
        s: Cell = (start.y + 1, start.x)
        (west, north, east, south) = ('.', '.', '.', '.')

    if w.x >= 0 and w.x < len(maze[0]):
        west = maze[w.y][w.x]
    if n.y >= 0 and n.y < len(maze):
        north = maze[n.y][n.x]
    if e.x >= 0 and e.x < len(maze[0]):
        east = maze[e.y][e.x]
    if s.y >= 0 and s.y < len(maze):
        south = maze[s.y][s.x]
    
    maze[start.y][start.x] = startSymbol(west, north, east, south)
    
    var dir = LEFT
    let symbol = maze[start.y][start.x]
    if symbol in "|JL":
        dir = UP
    elif symbol in "-7":
        dir = DOWN
    elif symbol == 'F':
        dir = RIGHT
    # echo symbol, start, dir
    return walk(start, dir, maze)


test "logictest":
    check startSymbol('.', '-', '|', '.') == 'F'
    check startSymbol('7', 'L', '-', '|') == 'F'
    check startSymbol('|', '|', 'L', 'L') == '|'
    check startSymbol('.', '|', '-', '.') == 'L'
    check startSymbol('.', '7', '7', '.') == 'L'
    check startSymbol('L', '|', '.', '.') == 'J'
    check startSymbol('F', '|', '.', '.') == 'J'
    check startSymbol('L', '.', '7', '.') == '-'
    check startSymbol('-', '.', 'J', '.') == '-'
    