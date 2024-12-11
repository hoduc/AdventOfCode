# https://adventofcode.com/2024/day/10
"""
--- Day 10: Hoof It ---

You all arrive at a Lava Production Facility on a floating island in the sky. As the others begin to search the massive industrial complex, you feel a small nose boop your leg and look down to discover a reindeer wearing a hard hat.

The reindeer is holding a book titled "Lava Island Hiking Guide". However, when you open the book, you discover that most of it seems to have been scorched by lava! As you're about to ask how you can help, the reindeer brings you a blank topographic map of the surrounding area (your puzzle input) and looks up at you excitedly.

Perhaps you can help fill in the missing hiking trails?

The topographic map indicates the height at each position using a scale from 0 (lowest) to 9 (highest). For example:

0123
1234
8765
9876

Based on un-scorched scraps of the book, you determine that a good hiking trail is as long as possible and has an even, gradual, uphill slope. For all practical purposes, this means that a hiking trail is any path that starts at height 0, ends at height 9, and always increases by a height of exactly 1 at each step. Hiking trails never include diagonal steps - only up, down, left, or right (from the perspective of the map).

You look up from the map and notice that the reindeer has helpfully begun to construct a small pile of pencils, markers, rulers, compasses, stickers, and other equipment you might need to update the map with hiking trails.

A trailhead is any position that starts one or more hiking trails - here, these positions will always have height 0. Assembling more fragments of pages, you establish that a trailhead's score is the number of 9-height positions reachable from that trailhead via a hiking trail. In the above example, the single trailhead in the top left corner has a score of 1 because it can reach a single 9 (the one in the bottom left).

This trailhead has a score of 2:

...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9

(The positions marked . are impassable tiles to simplify these examples; they do not appear on your actual topographic map.)

This trailhead has a score of 4 because every 9 is reachable via a hiking trail except the one immediately to the left of the trailhead:

..90..9
...1.98
...2..7
6543456
765.987
876....
987....

This topographic map contains two trailheads; the trailhead at the top has a score of 1, while the trailhead at the bottom has a score of 2:

10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01

Here's a larger example:

89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732

This larger example has 9 trailheads. Considering the trailheads in reading order, they have scores of 5, 6, 5, 3, 1, 3, 5, 3, and 5. Adding these scores together, the sum of the scores of all trailheads is 36.

The reindeer gleefully carries over a protractor and adds it to the pile. What is the sum of the scores of all trailheads on your topographic map?

"""

# LOOK LIKE JUST A DFS
# BREAK TO READ THE PROBLEM AND SOCIAL MEDIA FOR 5 MINS

def part(file_name):
    trail_starts = []
    m = []
    yb = 0
    with open(file_name) as f:
        for line in f:
            line = line.strip()
            for x in range(len(line)):
                if line[x] == '0':
                    trail_starts.append((yb, x))
            m.append(line)
            xb = len(line)
            yb += 1
    print(yb, xb)
    print(trail_starts)    

    def in_bound(y, x):
        return y >= 0 and y < yb and x >= 0 and x < xb
    
    def print_map(path = set([])):
        for y in range(yb):
            for x in range(xb):
                if (y,x) in path:
                    print('x', end='')
                else:
                    print(m[y][x], end='')
            print()
        print("-----")


    def dfs(trail):
        night_reachable = set([]) # night haha
        # print_map()
        # score = 0
        
        q = [(trail[0], trail[1],set([]))]

        def add_to_q(py, px, ph, visited):
            if in_bound(py, px) and m[py][px].isdigit() and int(m[py][px]) == ph + 1:
                q.append((py, px, set(visited)))

        while q:
            py, px, visited = q.pop(0)
            ph = int(m[py][px])
            # print(py, px, ph)
            # OH HMM IT JUST ASK TO COUNT THE POSITION OF 9 reachable
            # OHH WAIT A MINS IT WANT REACHABLE PER DFS
            visited.add((py, px))
            if ph == 9:
                # print_map(visited)
                # TODO: prolly part 2 will need to cache but lateer
                night_reachable.add((py,px))
                continue
            # remember to cache the path
            # but for part1 might be okay unless they pull some thing
            # up down left right
            add_to_q(py-1, px, ph, visited)
            add_to_q(py+1, px, ph, visited)
            add_to_q(py, px-1, ph, visited)
            add_to_q(py, px+1, ph, visited)
            # print(q)
        return len(night_reachable)
    
    return sum([dfs(t) for t in trail_starts])

assert part("day10_sample.txt") == 2
assert part("day10_sample2.txt") == 4
assert part("day10_sample3.txt") == 3
assert part("day10_sample_larger.txt") == 36
print(part("day10.txt"))