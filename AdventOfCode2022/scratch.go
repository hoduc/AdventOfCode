package main
import(
    "fmt"
    "strings"
)

func inputToCrate(input string) string {
    return input[1:len(input)-1]
}

func lineToStacks(line string, stacks []string) {
    fmt.Println("len-b4:", len(line))
    splits := strings.Split(line, " ")
    fmt.Println("splits:", splits, len(splits))
    i := 0
    for _, e := range splits {
        // fmt.Println("{" + e + "}")
        if len(e) > 0 {
            crate := inputToCrate(e)
            fmt.Println("crate:", crate)
            if len(stacks) == 0 || i >= len(stacks) {
                stacks = append(stacks, "")
                fmt.Println("hello:", stacks)
            }
            stacks[i] += crate
            fmt.Println("hi:", stacks[i])
            fmt.Println("hi2:", stacks)
            i += 1
        }
    }
}

func lineToRow(line string) []int {
    row := []int{}
    for _, e := range line {
        row = append(row, int(e - '0'))
    }
    return row
}
func main() {
    // stacks := []string{}
    // lineToStacks("    [D]    ", stacks)
    // stacks := make([]string, 3)
    // stacks[0] += "abc"
    // stacks[1] += "def"
    // fmt.Println(stacks)
    fmt.Println(lineToRow("123456"))
}
