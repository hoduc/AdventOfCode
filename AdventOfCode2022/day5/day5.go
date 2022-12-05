// --- Day 5: Supply Stacks ---
// The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

// The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

// The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

// They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2
// In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

// Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
// In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

//         [Z]
//         [N]
//     [C] [D]
//     [M] [P]
//  1   2   3
// Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

//         [Z]
//         [N]
// [M]     [D]
// [C]     [P]
//  1   2   3
// Finally, one crate is moved from stack 1 to stack 2:

//         [Z]
//         [N]
//         [D]
// [C] [M] [P]
//  1   2   3
// The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

// After the rearrangement procedure completes, what crate ends up on top of each stack?

package main

import(
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

func inputToStacksOrMove(line string, stacks []string) error {
    if line[1] == byte('1'){
        return nil
    } else if line[0] == byte('m') {
        return move(line, stacks)
    } else {

        for i := 0;  i < len(line); i += 4 {
            fmt.Println(i, "hello:", string(line[i]))
            if line[i] == byte('[') {
                fmt.Println("YSSSS:", i / 4)
                stacks[i / 4] += string(line[i+1:i+2])
            }
        }
    }
    fmt.Println("==>", stacks)
    return nil
}

func move(line string, stacks []string) error {
    fmt.Println("line:", line)
    splits := strings.Split(line, " ")
    fmt.Println("splits:", splits[1], splits[3], splits[5])
    amt, err := strconv.Atoi(splits[1])
    if err != nil {
        fmt.Println("1")
        return err
    }
    from, err := strconv.Atoi(splits[3])
    if err != nil {
        return err
    }
    to, err := strconv.Atoi(splits[5])
    if err != nil {
        return err
    }
    fmt.Println(amt, from, to)
    from -= 1
    to -= 1
    fmt.Println("from:", stacks[from])
    fmt.Println("to:", stacks[to])
    for i := 0; i < amt && len(stacks[from]) > 0; i++ {
        top := stacks[from][0]
        fmt.Println("top:", string(top))
        stacks[from] = stacks[from][1:]
        stacks[to] = string(top) + stacks[to]
    }
    fmt.Println("=>", stacks)
    return nil
}


//go:embed day5.txt
var day5txt string

func part1() string {
    var stacks []string
    onLine := func(line string) error {
        // fmt.Println("this-line:", line)
        if len(line) > 0 {
            if len(stacks) == 0 {
                cap := len(strings.Split(line, " ")) / 3
                fmt.Println("cap:", cap)
                stacks = make([]string, cap)
            }
            if err := inputToStacksOrMove(line, stacks); err != nil {
                return err

            }
        }
        return nil
    }

    if err := util.ReadLinesEmbed(day5txt, onLine); err != nil {
        return ""
    }
    fmt.Println(stacks)
    topStacks := ""
    for _, s := range stacks {
        topStacks += string(s[0])
    }
    return topStacks
}


func main() {
    fmt.Println("part1:", part1())
}
