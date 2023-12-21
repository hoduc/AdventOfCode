#[
--- Day 21: Step Counter ---

You manage to catch the airship right as it's dropping someone else off on their all-expenses-paid trip to Desert Island! It even helpfully drops you off near the gardener and his massive farm.

"You got the sand flowing again! Great work! Now we just need to wait until we have enough sand to filter the water for Snow Island and we'll have snow again in no time."

While you wait, one of the Elves that works with the gardener heard how good you are at solving problems and would like your help. He needs to get his steps in for the day, and so he'd like to know which garden plots he can reach with exactly his remaining 64 steps.

He gives you an up-to-date map (your puzzle input) of his starting position (S), garden plots (.), and rocks (#). For example:

...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........

The Elf starts at the starting position (S) which also counts as a garden plot. Then, he can take one step north, south, east, or west, but only onto tiles that are garden plots. This would allow him to reach any of the tiles marked O:

...........
.....###.#.
.###.##..#.
..#.#...#..
....#O#....
.##.OS####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........

Then, he takes a second step. Since at this point he could be at either tile marked O, his second step would allow him to reach any garden plot that is one step north, south, east, or west of any tile that he could have reached after the first step:

...........
.....###.#.
.###.##..#.
..#.#O..#..
....#.#....
.##O.O####.
.##.O#...#.
.......##..
.##.#.####.
.##..##.##.
...........

After two steps, he could be at any of the tiles marked O above, including the starting position (either by going north-then-south or by going west-then-east).

A single third step leads to even more possibilities:

...........
.....###.#.
.###.##..#.
..#.#.O.#..
...O#O#....
.##.OS####.
.##O.#...#.
....O..##..
.##.#.####.
.##..##.##.
...........

He will continue like this until his steps for the day have been exhausted. After a total of 6 steps, he could reach any of the garden plots marked O:

...........
.....###.#.
.###.##.O#.
.O#O#O.O#..
O.O.#.#.O..
.##O.O####.
.##.O#O..#.
.O.O.O.##..
.##.#.####.
.##O.##.##.
...........

In this example, if the Elf's goal was to get exactly 6 more steps today, he could use them to reach any of 16 garden plots.

However, the Elf actually needs to get 64 steps today, and the map he's handed you is much larger than the example map.

Starting from the garden plot marked S on your map, how many garden plots could the Elf reach in exactly 64 steps?

]#
import std/sets

proc echoGarden(garden: seq[string], yb, xb: int, visited: HashSet[tuple[y,x:int]]) =
    for y in 0 ..< yb:
        for x in 0 ..< xb:
            if (y, x) in visited:
                stdout.write('O')
            else:
                stdout.write(garden[y][x])
        stdout.write('\n')

proc day21*(fileName: string, step: int): int =
    var
        garden: seq[string]
        yb = 0
        xb = 0
        start = (0,0)
    for line in lines(fileName):
        xb = len(line)
        for x, c in line.pairs:
            if c == 'S':
                start = (yb, x)
        garden.add(line)
        yb += 1
    
    # echo start
    # echo garden

    # do bfs
    var 
        q = @[start]
        st = 0
        visited = initHashSet[tuple[y, x: int]]()
    while st < step:
        # echo len(q), " | ", q, len(toHashSet(q))
        # echo len(visited), visited
        echo "step:", st
        var next: seq[tuple[y,x:int]]
        # echoGarden(garden, yb, xb, toHashSet(q))
        for (y, x) in q:
            # if (y,x) in visited:
            #     continue
            visited.incl((y, x))
            for (ny, nx) in @[(y-1,x), (y+1, x), (y, x-1), (y, x+1)]:
                if ny < 0 or ny >= yb or nx < 0 or nx >= xb or (ny,nx) in visited or (garden[ny][nx] != '.' and garden[ny][nx] == '#'):
                    continue
                visited.incl((ny, nx))
                next.add((ny, nx))
        visited = initHashSet[tuple[y, x: int]]()
        q = next
        st += 1
    # echo len(visited), visited
    # echo "step:", st
    # echoGarden(garden, yb, xb, toHashSet(q))
    return len(toHashSet(q))