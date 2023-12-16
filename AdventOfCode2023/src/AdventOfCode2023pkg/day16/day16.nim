#[
--- Day 16: The Floor Will Be Lava ---

With the beam of light completely focused somewhere, the reindeer leads you deeper still into the Lava Production Facility. At some point, you realize that the steel facility walls have been replaced with cave, and the doorways are just cave, and the floor is cave, and you're pretty sure this is actually just a giant cave.

Finally, as you approach what must be the heart of the mountain, you see a bright light in a cavern up ahead. There, you discover that the beam of light you so carefully focused is emerging from the cavern wall closest to the facility and pouring all of its energy into a contraption on the opposite side.

Upon closer inspection, the contraption appears to be a flat, two-dimensional square grid containing empty space (.), mirrors (/ and \), and splitters (| and -).

The contraption is aligned so that most of the beam bounces around the grid, but each tile on the grid converts some of the beam's light into heat to melt the rock in the cavern.

You note the layout of the contraption (your puzzle input). For example:

.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....

The beam enters in the top-left corner from the left and heading to the right. Then, its behavior depends on what it encounters as it moves:

    If the beam encounters empty space (.), it continues in the same direction.
    If the beam encounters a mirror (/ or \), the beam is reflected 90 degrees depending on the angle of the mirror. For instance, a rightward-moving beam that encounters a / mirror would continue upward in the mirror's column, while a rightward-moving beam that encounters a \ mirror would continue downward from the mirror's column.
    If the beam encounters the pointy end of a splitter (| or -), the beam passes through the splitter as if the splitter were empty space. For instance, a rightward-moving beam that encounters a - splitter would continue in the same direction.
    If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.

Beams do not interact with other beams; a tile can have many beams passing through it at the same time. A tile is energized if that tile has at least one beam pass through it, reflect in it, or split in it.

In the above example, here is how the beam of light bounces around the contraption:

>|<<<\....
|v-.\^....
.v...|->>>
.v...v^.|.
.v...v^...
.v...v^..\
.v../2\\..
<->-/vv|..
.|<<<2-|.\
.v//.|.v..

Beams are only shown on empty tiles; arrows indicate the direction of the beams. If a tile contains beams moving in multiple directions, the number of distinct directions is shown instead. Here is the same diagram but instead only showing whether a tile is energized (#) or not (.):

######....
.#...#....
.#...#####
.#...##...
.#...##...
.#...##...
.#..####..
########..
.#######..
.#...#.#..

Ultimately, in this example, 46 tiles become energized.

The light isn't energizing enough tiles to produce lava; to debug the contraption, you need to start by analyzing the current situation. With the beam starting in the top-left heading right, how many tiles end up being energized?

--- Part Two ---

As you try to work out what might be wrong, the reindeer tugs on your shirt and leads you to a nearby control panel. There, a collection of buttons lets you align the contraption so that the beam enters from any edge tile and heading away from that edge. (You can choose either of two directions for the beam if it starts on a corner; for instance, if the beam starts in the bottom-right corner, it can start heading either left or upward.)

So, the beam could start on any tile in the top row (heading downward), any tile in the bottom row (heading upward), any tile in the leftmost column (heading right), or any tile in the rightmost column (heading left). To produce lava, you need to find the configuration that energizes as many tiles as possible.

In the above example, this can be achieved by starting the beam in the fourth tile from the left in the top row:

.|<2<\....
|v-v\^....
.v.v.|->>>
.v.v.v^.|.
.v.v.v^...
.v.v.v^..\
.v.v/2\\..
<-2-/vv|..
.|<<<2-|.\
.v//.|.v..

Using this configuration, 51 tiles are energized:

.#####....
.#.#.#....
.#.#.#####
.#.#.##...
.#.#.##...
.#.#.##...
.#.#####..
########..
.#######..
.#...#.#..

Find the initial beam configuration that energizes the largest number of tiles; how many tiles are energized in that configuration?

]#

import std/sets
import std/algorithm
import unittest

# model both position and direction

type Coord = object
    y, x: int

proc makeCoord(y, x: int): Coord =
    return Coord(y: y, x: x)

proc makeCoord(p: tuple[y,x: int]): Coord =
    return makeCoord(p.y, p.x)

let
    NORTH = makeCoord(-1, 0)
    SOUTH = makeCoord(1, 0)
    WEST  = makeCoord(0, -1)
    EAST  = makeCoord(0, 1)

type Entity = object
    position: Coord
    direction: Coord

proc makeEntity(py, px, dy, dx: int): Entity =
    return Entity(position: makeCoord(py, px), direction: makeCoord(dy, dx))

proc makeEntity(position: tuple[y,x: int], direction: tuple[y,x: int]): Entity =
    return makeEntity(position.y, position.x, direction.y, direction.x)

proc makeEntity(py, px: int, direction: Coord): Entity =
    return makeEntity(py, px, direction.y, direction.x)

proc makeEntity(position: Coord, dy, dx: int): Entity =
    return makeEntity(position.y, position.x, dy, dx)

proc makeEntity(position: Coord, direction: Coord): Entity =
    return makeEntity(position.y, position.x, direction.y, direction.x)

# move entity in the same direction that it is heading
proc moveEntity(entity: Entity): Entity =
    return makeEntity((entity.position.y + entity.direction.y, entity.position.x + entity.direction.x), (entity.direction.y, entity.direction.x))

proc hit(beam: Entity, bl: int): seq[Entity] =
    if bl == ord('.') or
        (bl == ord('|') and (beam.direction == NORTH or beam.direction == SOUTH)) or
        (bl == ord('-') and (beam.direction == WEST or beam.direction == EAST))
        :
        return @[moveEntity(beam)]
    elif bl == ord('|') and (beam.direction == EAST or beam.direction == WEST):
        return @[makeEntity(beam.position.y - 1, beam.position.x, NORTH), makeEntity(beam.position.y + 1, beam.position.x, SOUTH)]
    elif bl == ord('-') and (beam.direction == NORTH or beam.direction == SOUTH):
        return @[makeEntity(beam.position.y, beam.position.x - 1, WEST), makeEntity(beam.position.y, beam.position.x + 1, EAST)]
    elif bl == ord('/'):
        if beam.direction == WEST:
            return @[makeEntity(beam.position.y + 1, beam.position.x, SOUTH)]
        if beam.direction == EAST:
            return @[makeEntity(beam.position.y - 1, beam.position.x, NORTH)]
        if beam.direction == NORTH:
            return @[makeEntity(beam.position.y, beam.position.x + 1, EAST)]
        return @[makeEntity(beam.position.y, beam.position.x - 1, WEST)]
    elif bl == 92: # backslash confused it ('\')
        if beam.direction == WEST:
            return @[makeEntity(beam.position.y - 1, beam.position.x, NORTH)]
        if beam.direction == EAST:
            return @[makeEntity(beam.position.y + 1, beam.position.x, SOUTH)]
        if beam.direction == NORTH:
            return @[makeEntity(beam.position.y, beam.position.x - 1, WEST)]
        return @[makeEntity(beam.position.y, beam.position.x + 1, EAST)]
    return @[makeEntity(0, 0, 0, 0)]


proc echoContraption(contraption: seq[string], visited: HashSet[Coord]) =
    echo "~~~~~~~~~~~~~~~~~~~~~~~"
    for y in 0 ..< len(contraption):
        for x in 0 ..< len(contraption[0]):
            if makeCoord(y, x) in visited:
                stdout.write('x')
            else:
                stdout.write(contraption[y][x])
        stdout.write('\n')

proc readContraption(fileName: string): (seq[string], int, int) =
    var
        contraption: seq[string]
        yb = 0
        xb = 0
    for line in lines(fileName):
        xb = len(line)
        contraption.add(line)
        yb += 1
    return (contraption, yb, xb)

proc explore(startPosition: Entity, contraption: seq[string], yb, xb: int): int =
    var
        q = @[startPosition]
        visited = initHashSet[Entity]()
        visitedPosition = initHashSet[Coord]()
    while len(q) > 0:
        # echo q
        # echo visited
        let beam = q[0]
        visitedPosition.incl(beam.position)
        q.delete(0)
        visited.incl(beam)
        let
            y = beam.position.y
            x = beam.position.x
            bl = contraption[y][x]
            beams: seq[Entity] = hit(beam, ord(bl))

        for b in beams.reversed:
            let (by, bx) = (b.position.y, b.position.x)
            if bx < 0 or bx >= xb or by < 0 or by >= yb or b in visited:
                continue
            q.insert(b, 0)
    # echo contraption
    return len(visitedPosition)

proc day161*(fileName: string): int =
    var (contraption, yb, xb) = readContraption(fileName)
    return explore(makeEntity(0, 0, EAST), contraption, yb, xb)

proc day162*(fileName: string): int =
    var (contraption, yb, xb) = readContraption(fileName)
    var maxConfiguration = -1
    # top most left, right
    for y in 0 ..< yb:
        maxConfiguration = max(maxConfiguration, explore(makeEntity(y, 0, EAST), contraption, yb, xb))
        maxConfiguration = max(maxConfiguration, explore(makeEntity(y, xb - 1, WEST), contraption, yb, xb))
    for x in 0 ..< xb:
        maxConfiguration = max(maxConfiguration, explore(makeEntity(0, x, SOUTH), contraption, yb, xb))
        maxConfiguration = max(maxConfiguration, explore(makeEntity(yb - 1, x, NORTH), contraption, yb, xb))
    return maxConfiguration

test "logic":
    var entity = makeEntity(0, 0, SOUTH)
    check hit(entity, ord('.')) == @[makeEntity(1, 0, SOUTH)]
    check hit(entity, ord('|')) == @[makeEntity(1, 0, SOUTH)]
    check hit(entity, ord('-')) == @[makeEntity(0, -1, WEST), makeEntity(0, 1, EAST)]
    check hit(entity, ord('/')) == @[makeEntity(0, -1, WEST)]
    check hit(entity, 92) == @[makeEntity(0, 1, EAST)]

    entity = makeEntity(0, 0, NORTH)
    check hit(entity, ord('|')) == @[makeEntity(-1, 0, NORTH)]
    check hit(entity, ord('-')) == @[makeEntity(0, -1, WEST), makeEntity(0, 1, EAST)]
    check hit(entity, ord('/')) == @[makeEntity(0, 1, EAST)]
    check hit(entity, 92) == @[makeEntity(0, -1, WEST)]

    entity = makeEntity(0, 0, WEST)
    check hit(entity, ord('-')) == @[makeEntity(0, -1, WEST)]
    check hit(entity, ord('|')) == @[makeEntity(-1, 0, NORTH), makeEntity(1, 0, SOUTH)]
    check hit(entity, ord('/')) == @[makeEntity(1, 0, SOUTH)]
    check hit(entity, 92) == @[makeEntity(-1, 0, NORTH)]

    entity = makeEntity(0, 0, EAST)
    check hit(entity, ord('-')) == @[makeEntity(0, 1, EAST)]
    check hit(entity, ord('|')) == @[makeEntity(-1, 0, NORTH), makeEntity(1, 0, SOUTH)]
    check hit(entity, ord('/')) == @[makeEntity(-1, 0, NORTH)]
    check hit(entity, 92) == @[makeEntity(1, 0, SOUTH)]
    