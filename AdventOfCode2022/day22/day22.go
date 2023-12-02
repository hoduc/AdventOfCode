// --- Day 22: Monkey Map ---

// The monkeys take you on a surprisingly easy trail through the jungle. They're even going in roughly the right direction according to your handheld device's Grove Positioning System.

// As you walk, the monkeys explain that the grove is protected by a force field. To pass through the force field, you have to enter a password; doing so involves tracing a specific path on a strangely-shaped board.

// At least, you're pretty sure that's what you have to do; the elephants aren't exactly fluent in monkey.

// The monkeys give you notes that they took when they last saw the password entered (your puzzle input).

// For example:

//         ...#
//         .#..
//         #...
//         ....
// ...#.......#
// ........#...
// ..#....#....
// ..........#.
//         ...#....
//         .....#..
//         .#......
//         ......#.

// 10R5L5R10L4R5L5

// The first half of the monkeys' notes is a map of the board. It is comprised of a set of open tiles (on which you can move, drawn .) and solid walls (tiles which you cannot enter, drawn #).

// The second half is a description of the path you must follow. It consists of alternating numbers and letters:

//     A number indicates the number of tiles to move in the direction you are facing. If you run into a wall, you stop moving forward and continue with the next instruction.
//     A letter indicates whether to turn 90 degrees clockwise (R) or counterclockwise (L). Turning happens in-place; it does not change your current tile.

// So, a path like 10R5 means "go forward 10 tiles, then turn clockwise 90 degrees, then go forward 5 tiles".

// You begin the path in the leftmost open tile of the top row of tiles. Initially, you are facing to the right (from the perspective of how the map is drawn).

// If a movement instruction would take you off of the map, you wrap around to the other side of the board. In other words, if your next tile is off of the board, you should instead look in the direction opposite of your current facing as far as you can until you find the opposite edge of the board, then reappear there.

// For example, if you are at A and facing to the right, the tile in front of you is marked B; if you are at C and facing down, the tile in front of you is marked D:

//         ...#
//         .#..
//         #...
//         ....
// ...#.D.....#
// ........#...
// B.#....#...A
// .....C....#.
//         ...#....
//         .....#..
//         .#......
//         ......#.

// It is possible for the next tile (after wrapping around) to be a wall; this still counts as there being a wall in front of you, and so movement stops before you actually wrap to the other side of the board.

// By drawing the last facing you had with an arrow on each tile you visit, the full path taken by the above example looks like this:

//         >>v#
//         .#v.
//         #.v.
//         ..v.
// ...#...v..v#
// >>>v...>#.>>
// ..#v...#....
// ...>>>>v..#.
//         ...#....
//         .....#..
//         .#......
//         ......#.

// To finish providing the password to this strange input device, you need to determine numbers for your final row, column, and facing as your final position appears from the perspective of the original map. Rows start from 1 at the top and count downward; columns start from 1 at the left and count rightward. (In the above example, row 1, column 1 refers to the empty space with no tile on it in the top-left corner.) Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^). The final password is the sum of 1000 times the row, 4 times the column, and the facing.

// In the above example, the final row is 6, the final column is 8, and the final facing is 0. So, the final password is 1000 * 6 + 4 * 8 + 0: 6032.

// Follow the path given in the monkeys' notes. What is the final password?

package main

import (
	_ "embed"
	"fmt"
	"github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
	"strconv"
)

type Position struct {
	x int
	y int
}

type Range struct {
	start int
	end   int
}

func newPos(x, y int) Position {
	return Position{x: x, y: y}
}

type TYPE int

const (
	MOVE TYPE = 1
	TURN      = 2
)

type Instruction struct {
	t     TYPE
	value int
}

func newInstruction(t TYPE, value int) Instruction {
	return Instruction{t: t, value: value}
}

func parseInstructions(line string) ([]Instruction, error) {
	instructions := []Instruction{}
	i := 0
	for j, c := range line {
		if c == 'L' || c == 'R' {
			n, err := strconv.Atoi(line[i:j])
			if err != nil {
				return instructions, err
			}
			// fmt.Print(n, "-", string(c), "|")
			instructions = append(instructions, []Instruction{newInstruction(MOVE, n), newInstruction(TURN, int(c))}...)
			i = j + 1
		}
	}
	if i < len(line) {
		n, err := strconv.Atoi(line[i:])
		if err != nil {
			return instructions, err
		}
		instructions = append(instructions, []Instruction{newInstruction(MOVE, n)}...)
	}
	return instructions, nil
}

type DIR int

const (
	WEST  DIR = 0
	NORTH     = 1
	EAST      = 2
	SOUTH     = 3
)

func (d DIR) marker() string {
	switch d {
	case WEST:
		return "<"
	case NORTH:
		return "^"
	case EAST:
		return ">"
	}
	return "v"
}

func (d DIR) score() int {
	switch d {
	case WEST:
		return 2
	case NORTH:
		return 3
	case EAST:
		return 0
	}
	return 1
}

type Dir struct {
	wrap int
	dir  DIR
	move func(Position, Range) Position
}

func isWall(position Position, walls map[Position]bool) bool {
	_, ok := walls[position]
	return ok == true
}

func bounds(dir Dir, position Position, rowsBounds map[int]Range, colsBounds map[int]Range) Range {
	// fmt.Println("bounds:", dir.dir.marker())
	switch dir.dir {
	case EAST, WEST:
		rb, _ := rowsBounds[position.y]
		// fmt.Println("rb:", rb)
		return rb
	}
	cb, _ := colsBounds[position.x]
	// fmt.Println("cb:", cb)
	return cb
}

func applyInstructions(start Position, instructions []Instruction, walls map[Position]bool, rowsBounds map[int]Range, colsBounds map[int]Range, trails map[Position]DIR, grid map[Position]bool, xMin, xMax, yMin, yMax int) (Position, DIR) {
	// start facing right
	dir := EAST
	dirs := []Dir{
		Dir{dir: WEST, move: func(position Position, r Range) Position {
			if position.x-1 < r.start {
				return newPos(r.end, position.y)
			}
			return newPos(position.x-1, position.y)
		}},
		Dir{dir: NORTH, move: func(position Position, r Range) Position {
			if position.y-1 < r.start {
				return newPos(position.x, r.end)
			}
			return newPos(position.x, position.y-1)
		}},
		Dir{dir: EAST, move: func(position Position, r Range) Position {
			if position.x+1 > r.end {
				return newPos(r.start, position.y)
			}
			return newPos(position.x+1, position.y)
		}},
		Dir{dir: SOUTH, move: func(position Position, r Range) Position {
			if position.y+1 > r.end {
				return newPos(position.x, r.start)
			}
			return newPos(position.x, position.y+1)
		}},
	}
	d := dirs[dir]
	// fmt.Println("start:", start)
	// fmt.Println("hello!!!")
	for _, instruction := range instructions {
		if instruction.t == MOVE {
			r := bounds(d, start, rowsBounds, colsBounds)
			// fmt.Println(d.dir.marker(), "move by:", instruction.value, "|range:", r)
			for i := 0; i < instruction.value; i++ {
				next := d.move(start, r)
				// fmt.Println("next:", next)
				if isWall(next, walls) {
					// fmt.Println(next, " is wall!!!")
					break
				}
				trails[start] = d.dir
				start = next
			}
		} else { // change dir
			if instruction.value == 'L' {
				if dir-1 < 0 {
					dir = len(dirs) - 1
				} else {
					dir -= 1
				}
			} else {
				dir = (dir + 1) % len(dirs)
			}
			d = dirs[dir]
			trails[start] = d.dir
		}
		// render(xMin, xMax, yMin, yMax, grid, walls, trails)
		// fmt.Println("=============")
	}
	return start, d.dir
}

func render(xMin, xMax, yMin, yMax int, grid map[Position]bool, walls map[Position]bool, trails map[Position]DIR) {
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			pos := newPos(x, y)
			// visited ?
			if dir, ok := trails[pos]; ok {
				fmt.Print(dir.marker())
				continue
			}
			// wall ?
			if _, ok := walls[pos]; ok {
				fmt.Print("#")
				continue
			}
			// grid ?
			if _, ok := grid[pos]; ok {
				fmt.Print(".")
				continue
			}
			// empty
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

//go:embed day22.txt
var day22txt string

func readMonkeyMap() int {
	walls := make(map[Position]bool)
	grid := make(map[Position]bool)   // only contains traversablex
	rowsBounds := make(map[int]Range) // enter x
	colsBounds := make(map[int]Range) // enter y
	instructions := []Instruction{}
	dirScanMode := false
	lineNo := 0
	var start *Position
	xMin, yMin := 0, 0
	xMax, yMax := -1, -1
	onLine := func(line string) error {
		fmt.Println(lineNo%10, line, len(line))
		if len(line) > 0 {
			yMax = lineNo
			if dirScanMode {
				dirScanMode = false
				ins, err := parseInstructions(line)
				if err != nil {
					return err
				}
				instructions = ins
			} else {
				startX, endX := -1, -1
				for i, val := range line {
					if val != ' ' {
						if i > xMax {
							xMax = i
						}
						if startX == -1 {
							startX = i
						}
						endX = i
						if _, ok := colsBounds[i]; !ok {
							colsBounds[i] = Range{start: lineNo, end: -1}
						}
						if val == '.' {
							grid[newPos(i, lineNo)] = true
							if start == nil {
								start = &Position{x: i, y: lineNo}
							}
						} else if val == '#' {
							walls[newPos(i, lineNo)] = true
						}
					} else if r, ok := colsBounds[i]; ok && r.end == -1 {
						colsBounds[i] = Range{start: r.start, end: lineNo - 1}
						// fmt.Println("i:", i, r, "=>", colsBounds[i])
					}
				}
				rowsBounds[lineNo] = Range{start: startX, end: endX}
				// to the end of the lines update the previous
				for i := len(line); i <= xMax; i++ {
					if r, ok := colsBounds[i]; ok && r.end == -1 {
						colsBounds[i] = Range{start: r.start, end: lineNo - 1}
					}
				}
			}
		} else if !dirScanMode {
			fmt.Println("colsBounds:", colsBounds)
			for k, v := range colsBounds {
				if v.end == -1 {
					colsBounds[k] = Range{start: v.start, end: lineNo - 1}
				}
			}
			dirScanMode = true
		}
		lineNo += 1
		return nil
	}

	if err := util.ReadLinesEmbed(day22txt, onLine); err != nil {
		return -1
	}

	fmt.Println("walls:", len(walls))
	// fmt.Println(walls)
	fmt.Println("instructions:", len(instructions))
	// fmt.Println(instructions)
	fmt.Printf("xMin: %v | xMax: %v\n", xMin, xMax)
	fmt.Printf("yMin: %v | yMax: %v\n", yMin, yMax)
	fmt.Printf("startX: %v | startY: %v\n", start.x, start.y)
	fmt.Printf("rb: %v\ncb:%v\n", rowsBounds, colsBounds)
	fmt.Println("----")
	trails := make(map[Position]DIR)
	end, endDir := applyInstructions(*start, instructions, walls, rowsBounds, colsBounds, trails, grid, xMin, xMax, yMin, yMax)
	render(xMin, xMax, yMin, yMax, grid, walls, trails)
	fmt.Println("position:", end)
	return (end.y+1)*1000 + (end.x+1)*4 + endDir.score()
}

func main() {
	fmt.Println("part1:", readMonkeyMap())
}
