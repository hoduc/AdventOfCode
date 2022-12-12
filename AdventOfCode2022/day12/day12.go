// --- Day 12: Hill Climbing Algorithm ---
// You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

// You ask the device for a heightmap of the surrounding area (your puzzle input). The heightmap shows the local area from above broken into a grid; the elevation of each square of the grid is given by a single lowercase letter, where a is the lowest elevation, b is the next-lowest, and so on up to the highest elevation, z.

// Also included on the heightmap are marks for your current position (S) and the location that should get the best signal (E). Your current position (S) has elevation a, and the location that should get the best signal (E) has elevation z.

// You'd like to reach E, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be at most one higher than the elevation of your current square; that is, if your current elevation is m, you could step to elevation n, but not to elevation o. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

// For example:

// Sabqponm
// abcryxxl
// accszExk
// acctuvwj
// abdefghi
// Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:

// v..v<<<<
// >v.vv<<^
// .>vv>E^^
// ..v>>>^^
// ..>>>>>^
// In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left (<), or right (>). The location that should get the best signal is still E, and . marks unvisited squares.

// This path reaches the goal in 31 steps, the fewest possible.

// What is the fewest steps required to move from your current position to the location that should get the best signal?


package main

import(
    "fmt"
    _ "embed"
    "container/heap"
    // "strconv"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type Location struct {
    x int
    y int
    value rune
}

func (loc Location) String() string {
    return fmt.Sprintf("(%v, %v, %v)", loc.x, loc.y, string(loc.value))
}


func push(q *util.PriorityQueue, value Location, priority int) {
    heap.Push(q, &util.Item{Value: value, Priority: priority, Index: -1})
}



// lifted directly from: https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func addTo(q *util.PriorityQueue, board []rune, R int, C int, visited map[Location]bool, dist map[Location]int, prev map[Location]Location, loc Location, addToX bool) {
    // fmt.Println("cur_loc:", loc)
    locDist := 0
    if val, ok := dist[loc]; ok {
        locDist = val
    }
    for i := -1; i <= 1; i += 2 {
        var newLoc Location
        newLoc.x = loc.x
        newLoc.y = loc.y
        if addToX {
            newLoc.x += i
        } else {
            newLoc.y += i
        }
        if newLoc.x < 0 || newLoc.x >= C || newLoc.y < 0 || newLoc.y >= R {
            continue
        }
        newLoc.value = board[newLoc.y*C + newLoc.x]
        if newLoc.value - loc.value > 1 {
            continue
        }
        newLocDist := locDist + 1
        oldNewLocDist, ok := dist[newLoc]
        if !ok || newLocDist < oldNewLocDist {
            dist[newLoc] = newLocDist
            prev[newLoc] = loc
        }
        push(q, newLoc, newLocDist)
    }
}

func drawPath(visited map[Location]bool, R, C int) {
    ps := []string{}
    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            ps = append(ps, "---")
        }
    }

    for loc, _:= range visited {
        ps[loc.y*C + loc.x] = "-X-"
    }

    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            fmt.Print(ps[i*C + j])
        }
        fmt.Println()
    }
}

//go:embed day12.txt
var day12txt string

func hillClimb() int {
    var start Location
    var end Location
    board := []rune{}
    q := &util.PriorityQueue{
        Items: []*util.Item{},
        LessFn: func(pi, pj int) bool {
            return pi < pj
        },
    }
    R := 0
    C := 0
    onLine := func(line string) error {
        fmt.Println(R, line)
        if C == 0 {
            C = len(line)
        }
        if len(line) > 0 {
            for i, c := range line {
                v := c
                if c == 'S' {
                    v = 'a'
                    start = Location{x: i, y: R, value: v}
                } else if c == 'E' {
                    v = 'z'
                    end = Location{x: i, y: R, value: v}
                }
                board = append(board, v)
            }
        }
        R += 1
        return nil
    }

    if err := util.ReadLinesEmbed(day12txt, onLine); err != nil {
        return -1
    }
    fmt.Println("start:", start)
    fmt.Println("end:", end)
    fmt.Println("R:", R)
    fmt.Println("C:", C)
    visited := make(map[Location]bool)
    dist := make(map[Location]int)
    prev := make(map[Location]Location)
    dist[start] = 0
    heap.Push(q, &util.Item{Value: start, Priority: dist[start], Index: -1})
    // fmt.Println(q.Items)
    for q.Len() > 0 {
        location := heap.Pop(q).(*util.Item).Value.(Location)
        if _, ok := visited[location]; ok {
            continue
        }
        if location == end {
            // fmt.Println(q.Items)
            drawPath(visited, R, C)
            // fmt.Println("start:", start)
            // fmt.Println("end:", end)
            // fmt.Println("prev:", prev, len(prev))
            // fmt.Println("dist:", dist, len(dist))
            return dist[end]
        }
        visited[location] = true
        fmt.Println("visited:", location)
        // left right
        // fmt.Println("left-right")
        addTo(q, board, R, C, visited, dist, prev, location, true)
        // up down
        // fmt.Println("up-down")
        addTo(q, board, R, C, visited, dist, prev, location, false)
        // fmt.Println(dist)
        // fmt.Println(q.Len())
        // fmt.Println("visited:", len(visited))
    }
    return -1
}

func main() {
    fmt.Println("part1:", hillClimb())
}
