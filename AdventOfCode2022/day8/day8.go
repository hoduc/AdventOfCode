// --- Day 8: Treetop Tree House ---
// The expedition comes across a peculiar patch of tall trees all planted carefully in a grid. The Elves explain that a previous expedition planted these trees as a reforestation effort. Now, they're curious if this would be a good location for a tree house.

// First, determine whether there is enough tree cover here to keep a tree house hidden. To do this, you need to count the number of trees that are visible from outside the grid when looking directly along a row or column.

// The Elves have already launched a quadcopter to generate a map with the height of each tree (your puzzle input). For example:

// 30373
// 25512
// 65332
// 33549
// 35390
// Each tree is represented as a single digit whose value is its height, where 0 is the shortest and 9 is the tallest.

// A tree is visible if all of the other trees between it and an edge of the grid are shorter than it. Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.

// All of the trees around the edge of the grid are visible - since they are already on the edge, there are no trees to block the view. In this example, that only leaves the interior nine trees to consider:

// The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
// The top-middle 5 is visible from the top and right.
// The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
// The left-middle 5 is visible, but only from the right.
// The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
// The right-middle 3 is visible from the right.
// In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
// With 16 trees visible on the edge and another 5 visible in the interior, a total of 21 trees are visible in this arrangement.

// Consider your map; how many trees are visible from outside the grid?

// --- Part Two ---
// Content with the amount of tree cover available, the Elves just need to know the best spot to build their tree house: they would like to be able to see a lot of trees.

// To measure the viewing distance from a given tree, look up, down, left, and right from that tree; stop if you reach an edge or at the first tree that is the same height or taller than the tree under consideration. (If a tree is right on the edge, at least one of its viewing distances will be zero.)

// The Elves don't care about distant trees taller than those found by the rules above; the proposed tree house has large eaves to keep it dry, so they wouldn't be able to see higher than the tree house anyway.

// In the example above, consider the middle 5 in the second row:

// 30373
// 25512
// 65332
// 33549
// 35390
// Looking up, its view is not blocked; it can see 1 tree (of height 3).
// Looking left, its view is blocked immediately; it can see only 1 tree (of height 5, right next to it).
// Looking right, its view is not blocked; it can see 2 trees.
// Looking down, its view is blocked eventually; it can see 2 trees (one of height 3, then the tree of height 5 that blocks its view).
// A tree's scenic score is found by multiplying together its viewing distance in each of the four directions. For this tree, this is 4 (found by multiplying 1 * 1 * 2 * 2).

// However, you can do even better: consider the tree of height 5 in the middle of the fourth row:

// 30373
// 25512
// 65332
// 33549
// 35390
// Looking up, its view is blocked at 2 trees (by another tree with a height of 5).
// Looking left, its view is not blocked; it can see 2 trees.
// Looking down, its view is also not blocked; it can see 1 tree.
// Looking right, its view is blocked at 2 trees (by a massive tree of height 9).
// This tree's scenic score is 8 (2 * 2 * 1 * 2); this is the ideal spot for the tree house.

// Consider each tree on your map. What is the highest scenic score possible for any tree?

package main

import(
    "fmt"
    _ "embed"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)



//go:embed day8.txt
var day8txt string

type onPreTreeFn func(int, int) int
type onTreeFn func([]int, int, int, int, int, int) int


func treeTop(onPreTree onPreTreeFn, onTree onTreeFn) int {
    trees := []int{}
    R := 0
    C := 0

    onLine := func(line string) error {
        // fmt.Println(R, line)
        if C == 0 {
            C = len(line)
        }
        if len(line) > 0 {
            for _, ch := range line {
                chi := int(ch - '0')
                trees = append(trees, chi)
            }
        }
        R += 1
        return nil
    }

    if err := util.ReadLinesEmbed(day8txt, onLine); err != nil {
        return -1
    }

    res := onPreTree(R, C)
    for i := 1; i < R - 1; i++ {
        for j := 1; j < C - 1; j++ {
            res = onTree(trees, i, j, R, C, res)
        }
    }

    return res
}


func onPreTreePart1(R, C int) int {
    return C*2 + (R-2)*2
}

func onTreePart1(trees []int, i, j, R, C, res int) int{
    h := trees[i*C + j]

    // fmt.Printf("---(%v, %v) %v\n", i, j, h)
    // left
    k := j - 1
    for k >= 0 && trees[i*C + k] < h {
        k -= 1
    }
    if k < 0 {
        return res + 1
    }
    // right
    k = j + 1
    for k < C && trees[i*C + k] < h {
        k += 1
    }

    if k >= C {
        return res + 1
    }

    // up
    k = i - 1
    for k >= 0 && trees[k*C + j] < h {
        k -= 1
    }
    if k < 0 {
        return res + 1
    }

    // down
    k = i + 1
    for k < R && trees[k*C + j] < h {
        k += 1
    }
    if k >= R {
        return res + 1
    }
    // fmt.Printf("[x]\n")
    // fmt.Println("---")
    return res
}

func onPreTreePart2(R, C int) int {
    return 0
}

func onTreePart2(trees []int, i, j, R, C, res int) int{
    h := trees[i*C + j]

    // fmt.Printf("---(%v, %v) %v\n", i, j, h)
    // left
    k := j - 1
    for k >= 0 && trees[i*C + k] < h {
        k -= 1
    }
    left := j - k
    if k < 0 {
        left -= 1
    }

    // right
    k = j + 1
    for k < C && trees[i*C + k] < h {
        k += 1
    }
    right := k - j
    if k >= C {
        right -= 1
    }


    // up
    k = i - 1
    for k >= 0 && trees[k*C + j] < h {
        k -= 1
    }

    up := i - k
    if k < 0 {
        up -= 1
    }

    // down
    k = i + 1
    for k < R && trees[k*C + j] < h {
        k += 1
    }
    down := k - i
    if k >= R {
        down -= 1
    }
    // fmt.Printf("[x]\n")
    // fmt.Println("---")
    // fmt.Printf("%v(%v, %v): [%v %v %v %v] => [%v]\n", h, i, j, left, right, up, down, left*right*up*down)
    return util.Max(res, left*right*up*down).(int)
}



func main() {
    fmt.Println("part1:", treeTop(onPreTreePart1, onTreePart1))
    fmt.Println("part2:", treeTop(onPreTreePart2, onTreePart2))
}
