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

// --- Part Two ---
// You realize you misread the scan. There isn't an endless void at the bottom of the scan - there's floor, and you're standing on it!

// You don't have time to scan the floor, so assume the floor is an infinite horizontal line with a y coordinate equal to two plus the highest y coordinate of any point in your scan.

// In the example above, the highest y coordinate of any point is 9, and so the floor is at y=11. (This is as if your scan contained one extra rock path like -infinity,11 -> infinity,11.) With the added floor, the example above now looks like this:

//         ...........+........
//         ....................
//         ....................
//         ....................
//         .........#...##.....
//         .........#...#......
//         .......###...#......
//         .............#......
//         .............#......
//         .....#########......
//         ....................
// <-- etc #################### etc -->
// To find somewhere safe to stand, you'll need to simulate falling sand until a unit of sand comes to rest at 500,0, blocking the source entirely and stopping the flow of sand into the cave. In the example above, the situation finally looks like this after 93 units of sand come to rest:

// ............o............
// ...........ooo...........
// ..........ooooo..........
// .........ooooooo.........
// ........oo#ooo##o........
// .......ooo#ooo#ooo.......
// ......oo###ooo#oooo......
// .....oooo.oooo#ooooo.....
// ....oooooooooo#oooooo....
// ...ooo#########ooooooo...
// ..ooooo.......ooooooooo..
// #########################
// Using your scan, simulate the falling sand until the source of the sand becomes blocked. How many units of sand come to rest?

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

func renderSand(obstacles, restUnits map[Location]bool, startX, width, startY, height , floorY int) {
    for y := startY; y < startY + height; y++ {
        for x := startX; x < startX + width; x++ {
            loc := Location{x: x, y: y}
            if _, ok := obstacles[loc]; ok {
                if _, ok := restUnits[loc]; ok {
                    fmt.Print("o")
                } else {
                    fmt.Print("#")
                }
            } else {
                if floorY != -1 && y == floorY {
                    fmt.Print("#")
                } else {
                    fmt.Print(".")
                }
            }
        }
        fmt.Println()
    }
}

type shouldStopFn func(Location, int) bool
type onSandRestFn func(Location)
type isBlockedFn func(map[Location]bool, Location, int) bool

//go:embed day14.txt
var day14txt string

func sandScan(onSandRest onSandRestFn, shouldStop shouldStopFn, isBlocked isBlockedFn) int {
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
    // fmt.Println(obstacles, len(obstacles))
    maxY := 0
    for location, _ := range obstacles {
        if location.y > maxY {
            maxY = location.y
        }
    }

    restUnits := make(map[Location]bool)
    sand := Location{x: 500, y: 0}
    fmt.Println("maxY:", maxY)
    for !shouldStop(sand, maxY){
        // fmt.Println(sand)
        nextLocations := []Location{
            Location{x: sand.x, y: sand.y + 1},
            Location{x: sand.x - 1, y: sand.y + 1},
            Location{x: sand.x + 1, y: sand.y + 1},
        }
        blocked := 0
        var nextLocation Location
        for _, val := range nextLocations {
            if isBlocked(obstacles, val, maxY) {
                blocked += 1
            } else {
                nextLocation = val
                break
            }
        }

        // all 3 ways blocked
        if blocked == 3 {
            // become rest
            // add to the obstacles map
            onSandRest(sand)
            if _, ok := obstacles[sand]; !ok {
                obstacles[sand] = true
                restUnits[sand] = true
                // fmt.Println("obstacles:", len(obstacles))
                // fmt.Println("restUnits:", len(restUnits))
                // renderSand(obstacles, restUnits, 485, 30, 0, 15, 11)
                // fmt.Println("----")
            }

            sand = Location{x: 500, y: 0}
        } else {
            // one of the three ways not blocked
            sand = nextLocation
        }
    }
    // fmt.Println("obstacles:", len(obstacles))
    // fmt.Println("restUnits:", len(restUnits))
    return len(restUnits)
}

func isLocObstacle(loc Location, obstacles map[Location]bool) bool {
    if _, ok := obstacles[loc]; ok {
        return true
    }
    return false
}

func part1() (onSandRestFn, shouldStopFn, isBlockedFn) {
    return func(sand Location) {
    }, func(sand Location, maxY int) bool {
        return sand.y >= maxY
    }, func(obstacles map[Location]bool, nextLoc Location, maxY int) bool {
        return isLocObstacle(nextLoc, obstacles)
    }
}

func part2() (onSandRestFn, shouldStopFn, isBlockedFn) {
    stopped := false
    return func(sand Location) {
        if sand.x == 500 && sand.y == 0 {
            // fmt.Println("Hello-you-should-stop-here")
            stopped = true
        }
    }, func(sand Location, maxY int) bool {
        return stopped == true
    }, func(obstacles map[Location]bool, nextLoc Location, maxY int) bool{
        // maybe it is floor
        if nextLoc.y == maxY + 2 {
            return true
        }
        return isLocObstacle(nextLoc, obstacles)
    }
}

func main() {
    // fmt.Println("part1:", sandScan(part1()))
    fmt.Println("part2:", sandScan(part2()))
}
