// --- Day 23: Unstable Diffusion ---

// You enter a large crater of gray dirt where the grove is supposed to be. All around you, plants you imagine were expected to be full of fruit are instead withered and broken. A large group of Elves has formed in the middle of the grove.

// "...but this volcano has been dormant for months. Without ash, the fruit can't grow!"

// You look up to see a massive, snow-capped mountain towering above you.

// "It's not like there are other active volcanoes here; we've looked everywhere."

// "But our scanners show active magma flows; clearly it's going somewhere."

// They finally notice you at the edge of the grove, your pack almost overflowing from the random star fruit you've been collecting. Behind you, elephants and monkeys explore the grove, looking concerned. Then, the Elves recognize the ash cloud slowly spreading above your recent detour.

// "Why do you--" "How is--" "Did you just--"

// Before any of them can form a complete question, another Elf speaks up: "Okay, new plan. We have almost enough fruit already, and ash from the plume should spread here eventually. If we quickly plant new seedlings now, we can still make it to the extraction point. Spread out!"

// The Elves each reach into their pack and pull out a tiny plant. The plants rely on important nutrients from the ash, so they can't be planted too close together.

// There isn't enough time to let the Elves figure out where to plant the seedlings themselves; you quickly scan the grove (your puzzle input) and note their positions.

// For example:

// ....#..
// ..###.#
// #...#.#
// .#...##
// #.###..
// ##.#.##
// .#..#..

// The scan shows Elves # and empty ground .; outside your scan, more empty ground extends a long way in every direction. The scan is oriented so that north is up; orthogonal directions are written N (north), S (south), W (west), and E (east), while diagonal directions are written NE, NW, SE, SW.

// The Elves follow a time-consuming process to figure out where they should each go; you can speed up this process considerably. The process consists of some number of rounds during which Elves alternate between considering where to move and actually moving.

// During the first half of each round, each Elf considers the eight positions adjacent to themself. If no other Elves are in one of those eight positions, the Elf does not do anything during this round. Otherwise, the Elf looks in each of four directions in the following order and proposes moving one step in the first valid direction:

//     If there is no Elf in the N, NE, or NW adjacent positions, the Elf proposes moving north one step.
//     If there is no Elf in the S, SE, or SW adjacent positions, the Elf proposes moving south one step.
//     If there is no Elf in the W, NW, or SW adjacent positions, the Elf proposes moving west one step.
//     If there is no Elf in the E, NE, or SE adjacent positions, the Elf proposes moving east one step.

// After each Elf has had a chance to propose a move, the second half of the round can begin. Simultaneously, each Elf moves to their proposed destination tile if they were the only Elf to propose moving to that position. If two or more Elves propose moving to the same position, none of those Elves move.

// Finally, at the end of the round, the first direction the Elves considered is moved to the end of the list of directions. For example, during the second round, the Elves would try proposing a move to the south first, then west, then east, then north. On the third round, the Elves would first consider west, then east, then north, then south.

// As a smaller example, consider just these five Elves:

// .....
// ..##.
// ..#..
// .....
// ..##.
// .....

// The northernmost two Elves and southernmost two Elves all propose moving north, while the middle Elf cannot move north and proposes moving south. The middle Elf proposes the same destination as the southwest Elf, so neither of them move, but the other three do:

// ..##.
// .....
// ..#..
// ...#.
// ..#..
// .....

// Next, the northernmost two Elves and the southernmost Elf all propose moving south. Of the remaining middle two Elves, the west one cannot move south and proposes moving west, while the east one cannot move south or west and proposes moving east. All five Elves succeed in moving to their proposed positions:

// .....
// ..##.
// .#...
// ....#
// .....
// ..#..

// Finally, the southernmost two Elves choose not to move at all. Of the remaining three Elves, the west one proposes moving west, the east one proposes moving east, and the middle one proposes moving north; all three succeed in moving:

// ..#..
// ....#
// #....
// ....#
// .....
// ..#..

// At this point, no Elves need to move, and so the process ends.

// The larger example above proceeds as follows:

// == Initial State ==
// ..............
// ..............
// .......#......
// .....###.#....
// ...#...#.#....
// ....#...##....
// ...#.###......
// ...##.#.##....
// ....#..#......
// ..............
// ..............
// ..............

// == End of Round 1 ==
// ..............
// .......#......
// .....#...#....
// ...#..#.#.....
// .......#..#...
// ....#.#.##....
// ..#..#.#......
// ..#.#.#.##....
// ..............
// ....#..#......
// ..............
// ..............

// == End of Round 2 ==
// ..............
// .......#......
// ....#.....#...
// ...#..#.#.....
// .......#...#..
// ...#..#.#.....
// .#...#.#.#....
// ..............
// ..#.#.#.##....
// ....#..#......
// ..............
// ..............

// == End of Round 3 ==
// ..............
// .......#......
// .....#....#...
// ..#..#...#....
// .......#...#..
// ...#..#.#.....
// .#..#.....#...
// .......##.....
// ..##.#....#...
// ...#..........
// .......#......
// ..............

// == End of Round 4 ==
// ..............
// .......#......
// ......#....#..
// ..#...##......
// ...#.....#.#..
// .........#....
// .#...###..#...
// ..#......#....
// ....##....#...
// ....#.........
// .......#......
// ..............

// == End of Round 5 ==
// .......#......
// ..............
// ..#..#.....#..
// .........#....
// ......##...#..
// .#.#.####.....
// ...........#..
// ....##..#.....
// ..#...........
// ..........#...
// ....#..#......
// ..............

// After a few more rounds...

// == End of Round 10 ==
// .......#......
// ...........#..
// ..#.#..#......
// ......#.......
// ...#.....#..#.
// .#......##....
// .....##.......
// ..#........#..
// ....#.#..#....
// ..............
// ....#..#..#...
// ..............

// To make sure they're on the right track, the Elves like to check after round 10 that they're making good progress toward covering enough ground. To do this, count the number of empty ground tiles contained by the smallest rectangle that contains every Elf. (The edges of the rectangle should be aligned to the N/S/E/W directions; the Elves do not have the patience to calculate arbitrary rectangles.) In the above example, that rectangle is:

// ......#.....
// ..........#.
// .#.#..#.....
// .....#......
// ..#.....#..#
// #......##...
// ....##......
// .#........#.
// ...#.#..#...
// ............
// ...#..#..#..

// In this region, the number of empty ground tiles is 110.

// Simulate the Elves' process and find the smallest rectangle that contains the Elves after 10 rounds. How many empty ground tiles does that rectangle contain?

// --- Part Two ---

// It seems you're on the right track. Finish simulating the process and figure out where the Elves need to go. How many rounds did you save them?

// In the example above, the first round where no Elf moved was round 20:

// .......#......
// ....#......#..
// ..#.....#.....
// ......#.......
// ...#....#.#..#
// #.............
// ....#.....#...
// ..#.....#.....
// ....#.#....#..
// .........#....
// ....#......#..
// .......#......

// Figure out where the Elves need to go. What is the number of the first round where no Elf moves?

package main

import (
	_ "embed"
	"fmt"
	"github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type Position struct {
	x int
	y int
}

func (p *Position) nw() Position {
	return Position{x: p.x - 1, y: p.y - 1}
}

func (p *Position) n() Position {
	return Position{x: p.x, y: p.y - 1}
}

func (p *Position) ne() Position {
	return Position{x: p.x + 1, y: p.y - 1}
}

func (p *Position) w() Position {
	return Position{x: p.x - 1, y: p.y}
}

func (p *Position) e() Position {
	return Position{x: p.x + 1, y: p.y}
}

func (p *Position) sw() Position {
	return Position{x: p.x - 1, y: p.y + 1}
}

func (p *Position) s() Position {
	return Position{x: p.x, y: p.y + 1}
}

func (p *Position) se() Position {
	return Position{x: p.x + 1, y: p.y + 1}
}

func newPos(x, y int) Position {
	return Position{x: x, y: y}
}

func hasElf(positions map[Position]bool, position Position) bool {
	_, ok := positions[position]
	return ok == true
}

func noElves(positions map[Position]bool, poss ...Position) bool {
	for _, pos := range poss {
		if hasElf(positions, pos) {
			return false
		}
	}
	return true
}

func addToPropositions(propositions map[Position][]Position, proposerPosition Position, proposedPosition Position) {
	if _, ok := propositions[proposedPosition]; !ok {
		propositions[proposedPosition] = []Position{}
	}
	propositions[proposedPosition] = append(propositions[proposedPosition], proposerPosition)
}

func updateAxisElfCount(axisElfCount map[int]int, axis, inc int) {
	if _, ok := axisElfCount[axis]; !ok {
		axisElfCount[axis] = 0
	}
	axisElfCount[axis] = axisElfCount[axis] + inc
	if axisElfCount[axis] <= 0 {
		delete(axisElfCount, axis)
	}
}

type checkElvesFn func(Position, map[Position]bool, map[Position][]Position) bool

func checkElves(position Position, positions map[Position]bool, propositions map[Position][]Position, poss ...Position) bool {
	if noElves(positions, poss...) {
		addToPropositions(propositions, position, poss[0])
		return true
	}
	return false
}

func checkElvesNorth(position Position, positions map[Position]bool, propositions map[Position][]Position) bool {
	n, ne, nw := position.n(), position.ne(), position.nw()
	return checkElves(position, positions, propositions, n, ne, nw)
}

func checkElvesSouth(position Position, positions map[Position]bool, propositions map[Position][]Position) bool {
	s, se, sw := position.s(), position.se(), position.sw()
	return checkElves(position, positions, propositions, s, se, sw)
}

func checkElvesWest(position Position, positions map[Position]bool, propositions map[Position][]Position) bool {
	w, nw, sw := position.w(), position.nw(), position.sw()
	return checkElves(position, positions, propositions, w, nw, sw)
}

func checkElvesEast(position Position, positions map[Position]bool, propositions map[Position][]Position) bool {
	e, ne, se := position.e(), position.ne(), position.se()
	return checkElves(position, positions, propositions, e, ne, se)
}

func proposePosition(round int, checkElves []checkElvesFn, position Position, positions map[Position]bool, propositions map[Position][]Position) int {
	if noElves(positions, position.n(), position.ne(), position.nw(), position.s(), position.se(), position.sw(), position.w(), position.e()) {
		return 0
	}
	for i := round; i < round+4; i++ {
		if checkElves[i%4](position, positions, propositions) {
			// did move
			return 1
		}
	}
	// no elves all around, no move
	return 0
}

func render(xMin, xMax, yMin, yMax int, positions map[Position]bool) {
	for y := yMin - 1; y <= yMax+1; y++ {
		for x := xMin - 1; x <= xMax+1; x++ {
			if hasElf(positions, newPos(x, y)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type keepSimulatingFn func(int, int) bool
type reportFn func(int, int, int, int, int, map[int]int) int

//go:embed day23.txt
var day23txt string

func plantSeedlings(keepSimulating keepSimulatingFn, report reportFn) int {
	positions := make(map[Position]bool)
	// xElfCounts := make(map[int]int)
	yElfCounts := make(map[int]int)
	xMin, xMax := -1, -1
	yMin, yMax := -1, -1
	lineNo := 0
	onLine := func(line string) error {
		fmt.Println(lineNo, line)
		if len(line) > 0 {
			for i, val := range line {
				// fmt.Println(i, val)
				if val == '#' {
					position := newPos(i, lineNo)
					positions[position] = true
					updateAxisElfCount(yElfCounts, lineNo, 1)
					if xMin == -1 || xMin > i {
						xMin = i
					}
					if xMax == -1 || xMax < i {
						xMax = i
					}
					if yMin == -1 || yMin > lineNo {
						yMin = lineNo
					}
					if yMax == -1 || yMax < lineNo {
						yMax = lineNo
					}
				}
			}
		}

		lineNo += 1
		return nil
	}

	if err := util.ReadLinesEmbed(day23txt, onLine); err != nil {
		return -1
	}
	fmt.Println("xMin:", xMin)
	fmt.Println("xMax:", xMax)
	fmt.Println("yMin:", yMin)
	fmt.Println("yMax:", yMax)
	render(xMin, xMax, yMin, yMax, positions)

	checkElves := []checkElvesFn{
		checkElvesNorth,
		checkElvesSouth,
		checkElvesWest,
		checkElvesEast,
	}
	// fmt.Println("before")
	// fmt.Println("positions:", positions)
	// fmt.Println("xPositions:", xElfCounts)
	// fmt.Println("yPositions:", yElfCounts)
	// fmt.Println("----")
	fmt.Println("elves:", len(positions))
	round := 0
	moved := 0
	for keepSimulating(moved, round) {
		propositions := make(map[Position][]Position)
		// first half
		moved = 0
		for position, _ := range positions {
			moved += proposePosition(round, checkElves, position, positions, propositions)
		}
		// second half
		for proposedPosition, proposers := range propositions {
			// fmt.Println("proposed:", proposedPosition, proposers)
			if len(proposers) > 1 {
				moved -= len(proposers)
				continue
			}
			// update proposer position
			proposer := proposers[0]
			delete(positions, proposers[0])
			positions[proposedPosition] = true
			updateAxisElfCount(yElfCounts, proposer.y, -1)
			updateAxisElfCount(yElfCounts, proposedPosition.y, 1)
			// update min max of grid
			proposedPositionX, proposedPositionY := proposedPosition.x, proposedPosition.y
			if xMin > proposedPositionX {
				xMin = proposedPositionX
			}
			if xMax < proposedPositionX {
				xMax = proposedPositionX
			}
			if yMin > proposedPositionY {
				yMin = proposedPositionY
			}
			if yMax < proposedPositionY {
				yMax = proposedPositionY
			}
		}
		round += 1
	}
	render(xMin, xMax, yMin, yMax, positions)
	return report(round, xMin, xMax, yMin, yMax, yElfCounts)
}

func part1() (keepSimulatingFn, reportFn) {
	return func(moved, round int) bool {
			return round == 0 || round < 10 && moved > 0
		}, func(round, xMin, xMax, yMin, yMax int, yElfCounts map[int]int) int {
			emptyGround := 0
			width := xMax - xMin + 1
			for y := yMin; y <= yMax; y++ {
				if count, ok := yElfCounts[y]; ok {
					emptyGround += width - count
				} else {
					emptyGround += width
				}
			}

			return emptyGround
		}
}

func part2() (keepSimulatingFn, reportFn) {
	return func(moved, round int) bool {
			return round == 0 || moved > 0
		}, func(round, xMin, xMax, yMin, yMax int, yElfCounts map[int]int) int {
			return round
		}
}

func main() {
	fmt.Println("part1:", plantSeedlings(part1()))
	fmt.Println("part2:", plantSeedlings(part2()))
}
