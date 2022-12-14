// --- Day 13: Distress Signal ---
// You climb the hill and again try contacting the Elves. However, you instead receive a signal you weren't expecting: a distress signal.

// Your handheld device must still not be working properly; the packets from the distress signal got decoded out of order. You'll need to re-order the list of received packets (your puzzle input) to decode the message.

// Your list consists of pairs of packets; pairs are separated by a blank line. You need to identify how many pairs of packets are in the right order.

// For example:

// [1,1,3,1,1]
// [1,1,5,1,1]

// [[1],[2,3,4]]
// [[1],4]

// [9]
// [[8,7,6]]

// [[4,4],4,4]
// [[4,4],4,4,4]

// [7,7,7,7]
// [7,7,7]

// []
// [3]

// [[[]]]
// [[]]

// [1,[2,[3,[4,[5,6,7]]]],8,9]
// [1,[2,[3,[4,[5,6,0]]]],8,9]
// Packet data consists of lists and integers. Each list starts with [, ends with ], and contains zero or more comma-separated values (either integers or other lists). Each packet is always a list and appears on its own line.

// When comparing two values, the first value is called left and the second value is called right. Then:

// If both values are integers, the lower integer should come first. If the left integer is lower than the right integer, the inputs are in the right order. If the left integer is higher than the right integer, the inputs are not in the right order. Otherwise, the inputs are the same integer; continue checking the next part of the input.
// If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
// If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].
// Using these rules, you can determine which of the pairs in the example are in the right order:

// == Pair 1 ==
// - Compare [1,1,3,1,1] vs [1,1,5,1,1]
//   - Compare 1 vs 1
//   - Compare 1 vs 1
//   - Compare 3 vs 5
//     - Left side is smaller, so inputs are in the right order

// == Pair 2 ==
// - Compare [[1],[2,3,4]] vs [[1],4]
//   - Compare [1] vs [1]
//     - Compare 1 vs 1
//   - Compare [2,3,4] vs 4
//     - Mixed types; convert right to [4] and retry comparison
//     - Compare [2,3,4] vs [4]
//       - Compare 2 vs 4
//         - Left side is smaller, so inputs are in the right order

// == Pair 3 ==
// - Compare [9] vs [[8,7,6]]
//   - Compare 9 vs [8,7,6]
//     - Mixed types; convert left to [9] and retry comparison
//     - Compare [9] vs [8,7,6]
//       - Compare 9 vs 8
//         - Right side is smaller, so inputs are not in the right order

// == Pair 4 ==
// - Compare [[4,4],4,4] vs [[4,4],4,4,4]
//   - Compare [4,4] vs [4,4]
//     - Compare 4 vs 4
//     - Compare 4 vs 4
//   - Compare 4 vs 4
//   - Compare 4 vs 4
//   - Left side ran out of items, so inputs are in the right order

// == Pair 5 ==
// - Compare [7,7,7,7] vs [7,7,7]
//   - Compare 7 vs 7
//   - Compare 7 vs 7
//   - Compare 7 vs 7
//   - Right side ran out of items, so inputs are not in the right order

// == Pair 6 ==
// - Compare [] vs [3]
//   - Left side ran out of items, so inputs are in the right order

// == Pair 7 ==
// - Compare [[[]]] vs [[]]
//   - Compare [[]] vs []
//     - Right side ran out of items, so inputs are not in the right order

// == Pair 8 ==
// - Compare [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
//   - Compare 1 vs 1
//   - Compare [2,[3,[4,[5,6,7]]]] vs [2,[3,[4,[5,6,0]]]]
//     - Compare 2 vs 2
//     - Compare [3,[4,[5,6,7]]] vs [3,[4,[5,6,0]]]
//       - Compare 3 vs 3
//       - Compare [4,[5,6,7]] vs [4,[5,6,0]]
//         - Compare 4 vs 4
//         - Compare [5,6,7] vs [5,6,0]
//           - Compare 5 vs 5
//           - Compare 6 vs 6
//           - Compare 7 vs 0
//             - Right side is smaller, so inputs are not in the right order
// What are the indices of the pairs that are already in the right order? (The first pair has index 1, the second pair has index 2, and so on.) In the above example, the pairs in the right order are 1, 2, 4, and 6; the sum of these indices is 13.

// Determine which pairs of packets are already in the right order. What is the sum of the indices of those pairs?

// --- Part Two ---
// Now, you just need to put all of the packets in the right order. Disregard the blank lines in your list of received packets.

// The distress signal protocol also requires that you include two additional divider packets:

// [[2]]
// [[6]]
// Using the same rules as before, organize all packets - the ones in your list of received packets as well as the two divider packets - into the correct order.

// For the example above, the result of putting the packets in the correct order is:

// []
// [[]]
// [[[]]]
// [1,1,3,1,1]
// [1,1,5,1,1]
// [[1],[2,3,4]]
// [1,[2,[3,[4,[5,6,0]]]],8,9]
// [1,[2,[3,[4,[5,6,7]]]],8,9]
// [[1],4]
// [[2]]
// [3]
// [[4,4],4,4]
// [[4,4],4,4,4]
// [[6]]
// [7,7,7]
// [7,7,7,7]
// [[8,7,6]]
// [9]
// Afterward, locate the divider packets. To find the decoder key for this distress signal, you need to determine the indices of the two divider packets and multiply them together. (The first packet is at index 1, the second packet is at index 2, and so on.) In this example, the divider packets are 10th and 14th, and so the decoder key is 140.

// Organize all of the packets into the correct order. What is the decoder key for the distress signal?

package main

import(
    "fmt"
    _ "embed"
    "strconv"
    "os"
    "sort"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)


func isList(exp string) bool {
    return len(exp) >= 2 && exp[0] == '[' && exp[len(exp)-1] == ']'
}

func elems(exp string) []string {
    // fmt.Println("exp:", exp)
    l := []string{}
    level := 0
    j := 0
    for i, val := range exp {
        if val == '[' {
            level += 1
        } else if val == ']' {
            level -= 1
        } else if val == ',' {
            if level == 0 {
                l = append(l, exp[j:i])
                j = i + 1
            }
        }
    }
    l = append(l, exp[j:])
    return l
}

// what about negative ?
func stringToInt(s string) int {
    val, _ := strconv.Atoi(s)
    return val
}

func stringCmp(left, right string) int {
    li := stringToInt(left)
    ri := stringToInt(right)
    // fmt.Println("print:", li, ri)
    if li < ri {
        return -1
    } else if li > ri {
        return 1
    }
    return 0
}

func cmp(left, right string) int {
    // both are numbers
    if !isList(left) && !isList(right) {
        if len(left) == 0 && len(right) != 0 {
            return -1
        } else if len(left) != 0 && len(right) == 0 {
            return 1
        } else if len(left) == len(right) && len(left) == 0 {
            return 0
        }
        // TODO: yo strings.Compare return wrong result for "4" and "10"
        // fmt.Println(left, "vs", right, "=>", strings.Compare(left, right), " | but:", stringCmp(left, right))
        // if len(left) == 0 && len(right) == 0 {
        //     return -1
        // }
        return stringCmp(left, right)
    }
    if (isList(left) && !isList(right)) {
       right = "[" + right + "]"
    } else if (!isList(left) && isList(right)) {
       left = "[" + left + "]"
    }
    // both are list
    // fmt.Printf("d:%v(%v)|%v(%v)\n", left, len(left), right, len(right))
    l := elems(left[1:len(left)-1])
    r := elems(right[1:len(right)-1])
    // fmt.Printf("l:%v => len: %v\n", l, len(l))
    // fmt.Printf("r:%v => len: %v\n", r, len(r))
    for i := 0; i < len(l) && i < len(r); i++ {
        // fmt.Printf("elem-l:%v\n", l[i])
        // fmt.Printf("elem-r:%v\n", r[i])
        cmpResult := cmp(l[i], r[i])
        if cmpResult != 0 {
            // fmt.Println("got-back-result:", cmpResult)
            return cmpResult
        }
    }
    if len(l) < len(r){
        return -1
    }
    if len(l) > len(r){
        return 1
    }
    return 0
}


func sum(l []int) int{
    count := 0
    for _,val := range l {
        count += val
    }
    return count
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func writePairs(path string, pairs []int) {

    // open file using READ & WRITE permission
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)

    if isError(err) {
        return
    }
    defer file.Close()

    // write into file
    _, err = file.WriteString(fmt.Sprintln(pairs))
    if isError(err) {
        return
    }

    // save changes
    err = file.Sync()
    if isError(err) {
        return
    }

    fmt.Println("==> done writing to file")
}

type onPairFn func(string, string, int)
type afterInputFn func() int

//go:embed day13.txt
var day13txt string

func readPairs(onPair onPairFn, afterInput afterInputFn) int {
    // rightOrderPairs := 0
    pairNo := 1
    lineNo := 1
    leftLine := 1
    rightLine := 2
    left := ""
    right := ""
    onLine := func(line string) error {
        // fmt.Println(lineNo, line)
        if len(line) > 0 {
            if lineNo == leftLine {
                left = line
                leftLine += 3
            } else {
                right = line
                rightLine += 3
            }
        } else {
            fmt.Printf(">> [Pair %v] Comparing\n%v\n and \n%v\n", pairNo, left, right)
            onPair(left, right, pairNo)
            pairNo += 1
            left = ""
            right = ""
        }
        lineNo += 1
        return nil
    }

    if err := util.ReadLinesEmbed(day13txt, onLine); err != nil {
        return -1
    }

    if len(left) != 0 && len(right) != 0 {
        fmt.Printf(">> [Last] Comparing %v and %v\n", left, right)
        onPair(left, right, pairNo)
        pairNo += 1
    }
    return afterInput()
}

func part1() (onPairFn, afterInputFn) {
    rightOrderPairs := 0
    return func(left, right string, pairNo int) {
        if cmp(left, right) < 0 {
            rightOrderPairs += pairNo
        }
    }, func() int {
        return rightOrderPairs
    }
}


func findPacket(packets []string, packet string) int {
    for i, val := range packets {
        if val == packet {
            return i
        }
    }
    return -1
}

func part2() (onPairFn, afterInputFn) {
    packets := []string{}
    return func(left, right string, pairNo int) {
        packets = append(packets, []string{left, right}...)
    }, func() int {
        p1 := "[[2]]"
        p2 := "[[6]]"
        packets = append(packets, []string{p1, p2}...)
        sort.SliceStable(packets, func(i, j int) bool {
            return cmp(packets[i], packets[j]) < 0
        })
        fmt.Println(packets, len(packets))
        p1i := findPacket(packets, p1) + 1
        p2i := findPacket(packets, p2) + 1
        fmt.Println("p1i:", p1i)
        fmt.Println("p2i:", p2i)
        return p1i * p2i
    }
}

func main() {
    // fmt.Println("part1:", readPairs(part1()))
    fmt.Println("part2:", readPairs(part2()))
}
