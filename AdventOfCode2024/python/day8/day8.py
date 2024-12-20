# https://adventofcode.com/2024/day/8
"""
--- Day 8: Resonant Collinearity ---

You find yourselves on the roof of a top-secret Easter Bunny installation.

While The Historians do their thing, you take a look at the familiar huge antenna. Much to your surprise, it seems to have been reconfigured to emit a signal that makes people 0.1% more likely to buy Easter Bunny brand Imitation Mediocre Chocolate as a Christmas gift! Unthinkable!

Scanning across the city, you find that there are actually many such antennas. Each antenna is tuned to a specific frequency indicated by a single lowercase letter, uppercase letter, or digit. You create a map (your puzzle input) of these antennas. For example:

............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............

The signal only applies its nefarious effect at specific antinodes based on the resonant frequencies of the antennas. In particular, an antinode occurs at any point that is perfectly in line with two antennas of the same frequency - but only when one of the antennas is twice as far away as the other. This means that for any pair of antennas with the same frequency, there are two antinodes, one on either side of them.

So, for these two antennas with frequency a, they create the two antinodes marked with #:

..........
...#......
..........
....a.....
..........
.....a....
..........
......#...
..........
..........

Adding a third antenna with the same frequency creates several more antinodes. It would ideally add four antinodes, but two are off the right side of the map, so instead it adds only two:

..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......#...
..........
..........

Antennas with different frequencies don't create antinodes; A and a count as different frequencies. However, antinodes can occur at locations that contain antennas. In this diagram, the lone antenna with frequency capital A creates no antinodes but has a lowercase-a-frequency antinode at its location:

..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......A...
..........
..........

The first example has antennas with two different frequencies, so the antinodes they create look like this, plus an antinode overlapping the topmost A-frequency antenna:

......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.

Because the topmost A-frequency antenna overlaps with a 0-frequency antinode, there are 14 total unique locations that contain an antinode within the bounds of the map.

Calculate the impact of the signal. How many unique locations within the bounds of the map contain an antinode?

--- Part Two ---

Watching over your shoulder as you work, one of The Historians asks if you took the effects of resonant harmonics into your calculations.

Whoops!

After updating your model, it turns out that an antinode occurs at any grid position exactly in line with at least two antennas of the same frequency, regardless of distance. This means that some of the new antinodes will occur at the position of each antenna (unless that antenna is the only one of its frequency).

So, these three T-frequency antennas now create many antinodes:

T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........

In fact, the three T-frequency antennas are all exactly in line with two antennas, so they are all also antinodes! This brings the total number of antinodes in the above example to 9.

The original example now has 34 antinodes, including the antinodes that appear on every antenna:

##....#....#
.#.#....0...
..#.#0....#.
..##...0....
....0....#..
.#...#A....#
...#..#.....
#....#.#....
..#.....A...
....#....A..
.#........#.
...#......##

Calculate the impact of the signal using this updated model. How many unique locations within the bounds of the map contain an antinode?


"""


# RUSTY BIT ON GEOMETRY NOW
# LINE THROUGH TWO POINTS
# y = mx + b
# m = (y2-y1)/(x2-x1)
# m < 0, slope go up => [(x1 - |x1-x2|, y1 + |y1 - y2|), (x2 + |x1-x2|, y2 - |y1 - y2|)]
# m > 0, slope go down => [(x1 - |x1-x2|, y1 - |y1 - y2|), (x2 + |x1-x2|, y2 + |y1 - y2|)]
# condition is sorted of reverse for the case y pointing downward
# on the condition that x1 <= x2
# also the special case will be that it is vertical or horizontal line

def antinodes_enumerate(y, x, yb, xb, dy, dx, antinodes_location, part2):
    yp = y + dy
    xp = x + dx
    while (yp >= 0 and yp < yb and xp >= 0 and xp < xb):
        antinodes_location.add((yp, xp))
        if not part2:
            break
        yp += dy
        xp += dx


# need to extend beyond just two points
def antinodes(y1, x1, y2, x2, yb, xb, antinodes_location, part2):
    dy = abs(y1 - y2)
    dx = abs(x1 - x2)
    # TODO: add the cap of the while loop
    if x1 == x2: # vertical line
        # dy = abs(y1 - y2)
        if y1 < y2:
            antinodes_enumerate(y1, x1, yb, xb, -dy, dx, antinodes_location, part2)
            antinodes_enumerate(y2, x2, yb, xb, dy, dx, antinodes_location, part2)
            return
        antinodes(y2, x2, y1, x1, yb, xb, antinodes_location, part2)
        return
    if y1 == y2: # horizontal line
        if x1 < x2:
            antinodes_enumerate(y1, x1, yb, xb, dy, -dx, antinodes_location, part2)
            antinodes_enumerate(y2, x2, yb, xb, dy, dx, antinodes_location, part2)
            return
        antinodes(y2, x2, y1, x1, yb, xb, antinodes_location, part2)
        return
    # diagonal line
    if x1 < x2:
        m = (y2-y1) // (x2-x1)
        if m < 0:
            antinodes_enumerate(y1, x1, yb, xb, dy, -dx, antinodes_location, part2)
            antinodes_enumerate(y2, x2, yb, xb, -dy, dx, antinodes_location, part2)
            return
        antinodes_enumerate(y1, x1, yb, xb, -dy, -dx, antinodes_location, part2)
        antinodes_enumerate(y2, x2, yb, xb, dy, dx, antinodes_location, part2)
        return
    antinodes(y2, x2, y1, x1, yb, xb, antinodes_location, part2)

def assert_antinodes(expected, actual):
    assert expected == actual, "Expected {} But Got {}".format(expected, actual)

# # diagonal up
# assert_antinodes([(9, 0), (-3, 6)], antinodes(5, 2, 1, 4, 999, 999))
# # diagonal down
# assert_antinodes([(-2, -2), (10, 7)], antinodes(2, 1, 6, 4))
# # vertical
# assert_antinodes([(-1, 4), (8, 4)], antinodes(2, 4, 5, 4))
# # horizontal
# assert_antinodes([(4, -1), (4, 8)], antinodes(4, 2, 4, 5))


# OKAY LETS PARSE THE FILE?

def part(file_name, part2 = False):
    antinodes_location = set([])
    antennas_by_type = {}
    board = []
    # bounds of the board
    yb = 0
    xb = -1
    with open(file_name) as f:
        for line in f:
            line = line.strip()
            for x in range(len(line)):
                # hmm forgot to strip line but is it okay ?
                if line[x].isalnum():
                    #print(line[x], (yb, x))
                    if line[x] not in antennas_by_type:
                        antennas_by_type[line[x]] = []
                    antennas_by_type[line[x]].append((yb,x))
            board.append(line)
            xb = len(line)
            yb += 1
    
    # OH WAIT ANTENNAS IS BY TYPE
    for _, antennas in antennas_by_type.items():
        for i in range(len(antennas)):
            aiy, aix = antennas[i]
            print((aiy, aix), "===")
            for j in range(i+1, len(antennas)):
                ajy, ajx = antennas[j]
                print((ajy, ajx), "---")
                antinodes(aiy, aix, ajy, ajx, yb, xb, antinodes_location, part2)
                if part2:
                    antinodes_location.add((aiy, aix))
                    antinodes_location.add((ajy, ajx))
            print("===")

    # print(yb, xb)
    # print(antennas)
    print(len(antinodes_location), antinodes_location)
    for y in range(yb):
        for x in range(xb):
            if (y,x) in antinodes_location:
                print("#", end='')
            else:
                print(board[y][x], end='')
        print()
    return len(antinodes_location)

# assert part("day8_example.txt") == 14

# assert part("day8_example2.txt") == 4

# assert part("day8_example3.txt") == 6. # I miscount on paper. This is 6 which is correct

# assert part("day8_example4.txt") == 2 # Hmm this is correct. wonder where is the fault is

# LETS construct another example simple one and see if it is correct but also watch the location


def part1():
    return part("day8.txt")

# print(part1()) # 364. ANSWER TOO HIGH . OKAY new ANSWER 357 (forgot to strip the line is the issue??) . ANYWAY LETS TRY SUBMIT. YESSSSSSSSS

def part2():
    return part("day8.txt", True)

# assert part("day8_example.txt", True) == 34
print(part2()) # 1266



