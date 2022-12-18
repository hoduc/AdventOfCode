// --- Day 17: Pyroclastic Flow ---
// Your handheld device has located an alternative exit from the cave for you and the elephants. The ground is rumbling almost continuously now, but the strange valves bought you some time. It's definitely getting warmer in here, though.

// The tunnels eventually open into a very tall, narrow chamber. Large, oddly-shaped rocks are falling into the chamber from above, presumably due to all the rumbling. If you can't work out where the rocks will fall next, you might be crushed!

// The five types of rocks have the following peculiar shapes, where # is rock and . is empty space:

// ####

// .#.
// ###
// .#.

// ..#
// ..#
// ###

// #
// #
// #
// #

// ##
// ##
// The rocks fall in the order shown above: first the - shape, then the + shape, and so on. Once the end of the list is reached, the same order repeats: the - shape falls first, sixth, 11th, 16th, etc.

// The rocks don't spin, but they do get pushed around by jets of hot gas coming out of the walls themselves. A quick scan reveals the effect the jets of hot gas will have on the rocks as they fall (your puzzle input).

// For example, suppose this was the jet pattern in your cave:

// >>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
// In jet patterns, < means a push to the left, while > means a push to the right. The pattern above means that the jets will push a falling rock right, then right, then right, then left, then left, then right, and so on. If the end of the list is reached, it repeats.

// The tall, vertical chamber is exactly seven units wide. Each rock appears so that its left edge is two units away from the left wall and its bottom edge is three units above the highest rock in the room (or the floor, if there isn't one).

// After a rock appears, it alternates between being pushed by a jet of hot gas one unit (in the direction indicated by the next symbol in the jet pattern) and then falling one unit down. If any movement would cause any part of the rock to move into the walls, floor, or a stopped rock, the movement instead does not occur. If a downward movement would have caused a falling rock to move into the floor or an already-fallen rock, the falling rock stops where it is (having landed on something) and a new rock immediately begins falling.

// Drawing falling rocks with @ and stopped rocks with #, the jet pattern in the example above manifests as follows:

// The first rock begins falling:
// |..@@@@.|
// |.......|
// |.......|
// |.......|
// +-------+

// Jet of gas pushes rock right:
// |...@@@@|
// |.......|
// |.......|
// |.......|
// +-------+

// Rock falls 1 unit:
// |...@@@@|
// |.......|
// |.......|
// +-------+

// Jet of gas pushes rock right, but nothing happens:
// |...@@@@|
// |.......|
// |.......|
// +-------+

// Rock falls 1 unit:
// |...@@@@|
// |.......|
// +-------+

// Jet of gas pushes rock right, but nothing happens:
// |...@@@@|
// |.......|
// +-------+

// Rock falls 1 unit:
// |...@@@@|
// +-------+

// Jet of gas pushes rock left:
// |..@@@@.|
// +-------+

// Rock falls 1 unit, causing it to come to rest:
// |..####.|
// +-------+

// A new rock begins falling:
// |...@...|
// |..@@@..|
// |...@...|
// |.......|
// |.......|
// |.......|
// |..####.|
// +-------+

// Jet of gas pushes rock left:
// |..@....|
// |.@@@...|
// |..@....|
// |.......|
// |.......|
// |.......|
// |..####.|
// +-------+

// Rock falls 1 unit:
// |..@....|
// |.@@@...|
// |..@....|
// |.......|
// |.......|
// |..####.|
// +-------+

// Jet of gas pushes rock right:
// |...@...|
// |..@@@..|
// |...@...|
// |.......|
// |.......|
// |..####.|
// +-------+

// Rock falls 1 unit:
// |...@...|
// |..@@@..|
// |...@...|
// |.......|
// |..####.|
// +-------+

// Jet of gas pushes rock left:
// |..@....|
// |.@@@...|
// |..@....|
// |.......|
// |..####.|
// +-------+

// Rock falls 1 unit:
// |..@....|
// |.@@@...|
// |..@....|
// |..####.|
// +-------+

// Jet of gas pushes rock right:
// |...@...|
// |..@@@..|
// |...@...|
// |..####.|
// +-------+

// Rock falls 1 unit, causing it to come to rest:
// |...#...|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// A new rock begins falling:
// |....@..|
// |....@..|
// |..@@@..|
// |.......|
// |.......|
// |.......|
// |...#...|
// |..###..|
// |...#...|
// |..####.|
// +-------+
// The moment each of the next few rocks begins falling, you would see this:

// |..@....|
// |..@....|
// |..@....|
// |..@....|
// |.......|
// |.......|
// |.......|
// |..#....|
// |..#....|
// |####...|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |..@@...|
// |..@@...|
// |.......|
// |.......|
// |.......|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |..@@@@.|
// |.......|
// |.......|
// |.......|
// |....##.|
// |....##.|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |...@...|
// |..@@@..|
// |...@...|
// |.......|
// |.......|
// |.......|
// |.####..|
// |....##.|
// |....##.|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |....@..|
// |....@..|
// |..@@@..|
// |.......|
// |.......|
// |.......|
// |..#....|
// |.###...|
// |..#....|
// |.####..|
// |....##.|
// |....##.|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |..@....|
// |..@....|
// |..@....|
// |..@....|
// |.......|
// |.......|
// |.......|
// |.....#.|
// |.....#.|
// |..####.|
// |.###...|
// |..#....|
// |.####..|
// |....##.|
// |....##.|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |..@@...|
// |..@@...|
// |.......|
// |.......|
// |.......|
// |....#..|
// |....#..|
// |....##.|
// |....##.|
// |..####.|
// |.###...|
// |..#....|
// |.####..|
// |....##.|
// |....##.|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+

// |..@@@@.|
// |.......|
// |.......|
// |.......|
// |....#..|
// |....#..|
// |....##.|
// |##..##.|
// |######.|
// |.###...|
// |..#....|
// |.####..|
// |....##.|
// |....##.|
// |....#..|
// |..#.#..|
// |..#.#..|
// |#####..|
// |..###..|
// |...#...|
// |..####.|
// +-------+
// To prove to the elephants your simulation is accurate, they want to know how tall the tower will get after 2022 rocks have stopped (but before the 2023rd rock begins falling). In this example, the tower of rocks will be 3068 units tall.

// How many units tall will the tower of rocks be after 2022 rocks have stopped falling?

package main

import(
    "fmt"
    _ "embed"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type Coord struct {
    x int
    y int
}

type Rock struct{
    pieces []Coord
}

const (
    WALL_LEFT = 0
    WALL_RIGHT = 6
    FLOOR_Y = 0
)


func xRock(xRocks [][][]int, xRockIndex int) [][]int {
    return xRocks[xRockIndex %  len(xRocks)]
}

func newRock(xRocks [][][]int, xRockIndex int, lineY int) Rock {
    pieces := []Coord{}
    nxr := xRock(xRocks, xRockIndex)
    // fmt.Println("lineY:", lineY)
    // fmt.Println("nxr:", nxr)
    // fmt.Println(len(nxr))
    y := lineY - 3 - (len(nxr) - 1)
    if lineY > 0 {
        y -= 1
    }
    // fmt.Println("y-start:", y)
    for i, r := range nxr {
        for _, x := range r {
            pieces = append(pieces, Coord{x: x + 2, y: y + i})
        }
    }
    // fmt.Println("newRockPieces:", pieces)
    return Rock{pieces: pieces}
}

func addRockPiecesToStopPieces(rock Rock, stoppedPieces map[Coord]bool, floor int) (bool, int){
    // fmt.Println("current-floor:", floor)
    for _, piece := range rock.pieces {
        stoppedPieces[piece] = true
        if piece.y <= floor {
            floor = piece.y - 1
        }
    }
    // fmt.Println("after-floor:", floor)
    return false, floor
}

func fall(rock Rock, stoppedPieces map[Coord]bool, floor int) (bool, int) {
    // fmt.Println("before:", rock)
    for _, piece := range rock.pieces {
        newPiece := Coord{x: piece.x, y: piece.y + 1}
        if newPiece.y > FLOOR_Y {
            // record stopped piece
            return addRockPiecesToStopPieces(rock, stoppedPieces, floor)
        }
        if _, ok := stoppedPieces[newPiece]; ok {
            return addRockPiecesToStopPieces(rock, stoppedPieces, floor)
        }
    }

    for i := 0; i < len(rock.pieces); i++ {
        piece := rock.pieces[i]
        rock.pieces[i] = Coord{x: piece.x , y: piece.y + 1}
    }
    return true, floor
}

func nextJetDir(jets string, index int) byte{
    return jets[index % len(jets)]
}

func jetPush(dir byte, rock Rock, stoppedPieces map[Coord]bool) bool {
    inc := 1
    if dir == '<' {
        inc = -1
    }
    // fmt.Println(string(dir))
    // fmt.Println("dir:", string(dir), "inc:", inc)
    // fmt.Println("before:", rock)

    for _, piece := range rock.pieces {
        newPiece := Coord{x: piece.x + inc, y: piece.y}
        // wall
        if newPiece.x < WALL_LEFT || newPiece.x > WALL_RIGHT {
            // fmt.Println("nothing-happen!!!")
            return false
        }
        // colliding with other piece
        if _, ok := stoppedPieces[newPiece]; ok {
            return false
        }
    }
    for i := 0; i < len(rock.pieces); i++ {
        piece := rock.pieces[i]
        rock.pieces[i] = Coord{x: piece.x + inc, y: piece.y}
    }

    return true
}

func render(rock Rock, stoppedPieces map[Coord]bool, floor int) {
    // get the top y
    // fmt.Println("floor:", floor)
    topY := floor - 3
    // fmt.Println("top-Y-before:", topY)
    rockPieces := make(map[Coord]bool)
    for _, piece := range rock.pieces {
        // fmt.Println("piece:", piece)
        if topY > piece.y {
            topY = piece.y
        }
        rockPieces[Coord{x: piece.x, y: piece.y}] = true
    }
    fmt.Println(rock.pieces)
    // fmt.Println("top-Y-after:", topY)
    fmt.Print("|")
    for x := WALL_LEFT; x <= WALL_RIGHT; x++ {
        fmt.Print(x)
    }
    fmt.Println("|")

    for y := topY; y <= FLOOR_Y; y++ {
        fmt.Print("|")
        for x := WALL_LEFT ; x <= WALL_RIGHT; x++ {
            coord := Coord{x: x, y: y}
            if _, ok := stoppedPieces[coord]; ok {
                fmt.Print("#")
            } else if _, ok := rockPieces[coord]; ok {
                fmt.Print("@")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Print("|")
        if y == floor {
            fmt.Print("   *")
        }
        fmt.Println()
    }

    fmt.Print("+")
    for x := WALL_LEFT; x <= WALL_RIGHT; x++ {
        fmt.Print("-")
    }

    fmt.Println("+")
    fmt.Println("=====================")
}

//go:embed day17.txt
var day17txt string

func tetris(restRockCount int) int {
    // x : left to right
    // y : bottom(0) to top(?)
    XROCKS := [][][]int{
        // ####
        // {2, 3, 4, 5},
        {{0, 1, 2, 3}},
        // .#.
        // ###
        // .#.
        //{3, 2, 3, 4, 3},
        {{1}, {0, 1, 2}, {1}},
        // ..#
        // ..#
        // ###
        {{2}, {2}, {0, 1, 2}},
        // #
        // #
        // #
        // #
        {{0}, {0}, {0}, {0}},
        // ##
        // ##
        {{0, 1}, {0, 1}},
    }
    jets := ""
    onLine := func(line string) error {
        if len(line) > 0 {
            jets = line
        }
        return nil
    }

    if err := util.ReadLinesEmbed(day17txt, onLine); err != nil {
        return -1
    }

    floor := 0
    stoppedPieces := make(map[Coord]bool)
    ri, ji := 0, 0
    rock := newRock(XROCKS, ri, floor)
    // fmt.Println("init:", rock)
    fmt.Println("jets:", jets)
    for i, restRock := 0, 0 ; restRock < restRockCount; i++ {
        // render(rock, stoppedPieces, floor)
        if i % 2 > 0 { // down
            // fmt.Println("fall")
            moved, newFloor := fall(rock, stoppedPieces, floor)
            if newFloor < floor {
                floor = newFloor
            }
            // fmt.Println("moved:", moved, "newFloor:", newFloor)
            if !moved { // get to rest
                restRock += 1
                ri += 1
                // spawn new rock
                rock = newRock(XROCKS, ri, floor)
                fmt.Println("restRock:", restRock)
                // render(rock, stoppedPieces, floor)
            }


        } else { // jet push
            // fmt.Println("jetPush:")
            jetPush(nextJetDir(jets, ji), rock, stoppedPieces)
            // render(rock, stoppedPieces, floor)
            ji += 1
        }
        // fmt.Println(rock)
    }
    return util.Abs(floor).(int)
}

func main() {
    fmt.Println("day1:", tetris(2022))
}
