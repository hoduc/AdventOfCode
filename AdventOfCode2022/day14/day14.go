// --- Day 14: Regolith Reservoir ---
// The distress signal leads you to a giant waterfall! Actually, hang on - the signal seems like it's coming from the waterfall itself, and that doesn't make any sense. However, you do notice a little path that leads behind the waterfall.

// Correction: the distress signal leads you behind a giant waterfall! There seems to be a large cave system here, and the signal definitely leads further inside.

// As you begin to make your way deeper underground, you feel the ground rumble for a moment. Sand begins pouring into the cave! If you don't quickly figure out where the sand is going, you could quickly become trapped!

// Fortunately, your familiarity with analyzing the path of falling material will come in handy here. You scan a two-dimensional vertical slice of the cave above you (your puzzle input) and discover that it is mostly air with structures made of rock.

// Your scan traces the path of each solid rock structure and reports the x,y coordinates that form the shape of the path, where x represents distance to the right and y represents distance down. Each path appears as a single line of text in your scan. After the first point of each path, each point indicates the end of a straight horizontal or vertical line to be drawn from the previous point. For example:

// 498,4 -> 498,6 -> 496,6
// 503,4 -> 502,4 -> 502,9 -> 494,9
// This scan means that there are two paths of rock; the first path consists of two straight lines, and the second path consists of three straight lines. (Specifically, the first path consists of a line of rock from 498,4 through 498,6 and another line of rock from 498,6 through 496,6.)

// The sand is pouring into the cave from point 500,0.

// Drawing rock as #, air as ., and the source of the sand as +, this becomes:


//   4     5  5
//   9     0  0
//   4     0  3
// 0 ......+...
// 1 ..........
// 2 ..........
// 3 ..........
// 4 ....#...##
// 5 ....#...#.
// 6 ..###...#.
// 7 ........#.
// 8 ........#.
// 9 #########.
// Sand is produced one unit at a time, and the next unit of sand is not produced until the previous unit of sand comes to rest. A unit of sand is large enough to fill one tile of air in your scan.

// A unit of sand always falls down one step if possible. If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move diagonally one step down and to the left. If that tile is blocked, the unit of sand attempts to instead move diagonally one step down and to the right. Sand keeps moving as long as it is able to do so, at each step trying to move down, then down-left, then down-right. If all three possible destinations are blocked, the unit of sand comes to rest and no longer moves, at which point the next unit of sand is created back at the source.

// So, drawing sand that has come to rest as o, the first unit of sand simply falls straight down and then stops:

// ......+...
// ..........
// ..........
// ..........
// ....#...##
// ....#...#.
// ..###...#.
// ........#.
// ......o.#.
// #########.
// The second unit of sand then falls straight down, lands on the first one, and then comes to rest to its left:

// ......+...
// ..........
// ..........
// ..........
// ....#...##
// ....#...#.
// ..###...#.
// ........#.
// .....oo.#.
// #########.
// After a total of five units of sand have come to rest, they form this pattern:

// ......+...
// ..........
// ..........
// ..........
// ....#...##
// ....#...#.
// ..###...#.
// ......o.#.
// ....oooo#.
// #########.
// After a total of 22 units of sand:

// ......+...
// ..........
// ......o...
// .....ooo..
// ....#ooo##
// ....#ooo#.
// ..###ooo#.
// ....oooo#.
// ...ooooo#.
// #########.
// Finally, only two more units of sand can possibly come to rest:

// ......+...
// ..........
// ......o...
// .....ooo..
// ....#ooo##
// ...o#ooo#.
// ..###ooo#.
// ....oooo#.
// .o.ooooo#.
// #########.
// Once all 24 units of sand shown above have come to rest, all further sand flows out the bottom, falling into the endless void. Just for fun, the path any new sand takes before falling forever is shown here with ~:

// .......+...
// .......~...
// ......~o...
// .....~ooo..
// ....~#ooo##
// ...~o#ooo#.
// ..~###ooo#.
// ..~..oooo#.
// .~o.ooooo#.
// ~#########.
// ~..........
// ~..........
// ~..........
// Using your scan, simulate the falling sand. How many units of sand come to rest before sand starts flowing into the abyss below?
package main

import(
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type Location struct {
    x int
    y int
}

func updateObstacles(obstacles map[Location]bool, locations []Location) {
    for i := 1; i < len(locations); i++ {
        start := locations[i-1]
        end := locations[i]
        if start.y > end.y || start.x > end.x {
            start, end = end, start
        }
        addToX := start.y == end.y
        obstacles[start] = true
        for start != end {
            if addToX {
                start.x += 1
            } else {
                start.y += 1
            }
            obstacles[start] = true
        }

    }
}

//go:embed day14.txt
var day14txt string

func read() int {
    obstacles := make(map[Location]bool)
    onLine := func(line string) error {
        if len(line) > 0 {
            splits := strings.Split(line, "->")
            ends := []Location{}
            for _, val := range splits {
                xy := strings.Split(strings.TrimSpace(val), ",")
                x, err := strconv.Atoi(xy[0])
                if err != nil {
                    return err
                }
                y, err := strconv.Atoi(xy[1])
                if err != nil {
                    return err
                }
                ends = append(ends, Location{x: x, y: y})
            }
            updateObstacles(obstacles, ends)
        }
        return nil
    }

    if err := util.ReadLinesEmbed(day14txt, onLine); err != nil {
        return -1
    }
    fmt.Println(obstacles, len(obstacles))
    maxY := 0
    for location, _ := range obstacles {
        if location.y > maxY {
            maxY = location.y
        }
    }

    restUnits := 0
    sand := Location{x: 500, y: 0}
    fmt.Println("maxY:", maxY)
    for sand.y < maxY {
        // fmt.Println(sand)
        blocked := 0
        // down blocked ?
        if _, ok := obstacles[Location{x: sand.x, y: sand.y + 1}]; ok {
            // fmt.Println("down-blocked!")
            blocked += 1
            // down-left blocked ?
            if _, ok := obstacles[Location{x: sand.x - 1, y: sand.y + 1}]; ok {
                // fmt.Println("down-left-blocked")
                blocked += 1
                if _, ok := obstacles[Location{x: sand.x + 1, y: sand.y + 1}]; ok {
                    // fmt.Println("down-right-blocked")
                    blocked += 1
                } else {
                    sand.x += 1
                    sand.y += 1
                }
            } else {
                sand.x -= 1
                sand.y += 1
            }
        } else {
            // fmt.Println("no-blocked")
            sand.y += 1
        }

        // try all three way
        if blocked == 3 {
            // become rest
            // add to the obstacles map
            obstacles[sand] = true
            sand = Location{x: 500, y:0}
            restUnits += 1
            fmt.Println("restUnits:", restUnits)
            continue
        }
    }
    return restUnits
}

func main() {
    fmt.Println("part1:", read())
}
