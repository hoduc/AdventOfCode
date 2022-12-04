// --- Day 3: Rucksack Reorganization ---
// One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

// Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

// The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

// The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

// For example, suppose you have the following list of contents from six rucksacks:

// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw
// The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
// The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
// The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
// The fourth rucksack's compartments only share item type v.
// The fifth rucksack's compartments only share item type t.
// The sixth rucksack's compartments only share item type s.
// To help prioritize item rearrangement, every item type can be converted to a priority:

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
// In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

// Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

package main
import(
    "fmt"
    "unicode"
    _ "embed"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

func itemPriority(item rune) int {
    if unicode.IsUpper(item) {
        return int(item - 'A') + 27
    }
    return int(item - 'a') + 1
}

func itemsCount(items string) map[rune]bool{
    // fmt.Println("items:", items)
    f := map[rune]bool {}
    for _, e := range items {
        f[e] = true
    }
    return f
}

func itemIntersect(left, right map[rune]bool) rune{
    if len(left) > len(right) {
        left, right = right, left
    }
    for item, _ := range left {
        if right[item] {
            return item
        }
    }
    return '0'
}

//go:embed day3.txt
var day3txt string

func part1() int {
    sum_priority := 0
    onLine := func(line string) error {
        if len(line) > 0 {
            sum_priority += itemPriority(itemIntersect(itemsCount(line[:len(line)/2]), itemsCount(line[len(line)/2:])))
        }
        return nil
    }
    if err := util.ReadLinesEmbed(day3txt, onLine); err != nil {
        return -1
    }
    return sum_priority
}

func main() {
    // s := "vJrwpWtwJgWrhcsFMMfFFhFp"
    // p := itemPriority(itemIntersect(itemsCount(s[:len(s)/2]), itemsCount(s[len(s)/2:])))
    // fmt.Println(p)
    fmt.Println("part1:", part1())
}
