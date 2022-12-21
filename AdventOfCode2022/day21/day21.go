// --- Day 21: Monkey Math ---
// The monkeys are back! You're worried they're going to try to steal your stuff again, but it seems like they're just holding their ground and making various monkey noises at you.

// Eventually, one of the elephants realizes you don't speak monkey and comes over to interpret. As it turns out, they overheard you talking about trying to find the grove; they can show you a shortcut if you answer their riddle.

// Each monkey is given a job: either to yell a specific number or to yell the result of a math operation. All of the number-yelling monkeys know their number from the start; however, the math operation monkeys need to wait for two other monkeys to yell a number, and those two other monkeys might also be waiting on other monkeys.

// Your job is to work out the number the monkey named root will yell before the monkeys figure it out themselves.

// For example:

// root: pppw + sjmn
// dbpl: 5
// cczh: sllz + lgvd
// zczc: 2
// ptdq: humn - dvpt
// dvpt: 3
// lfqf: 4
// humn: 5
// ljgn: 2
// sjmn: drzm * dbpl
// sllz: 4
// pppw: cczh / lfqf
// lgvd: ljgn * ptdq
// drzm: hmdt - zczc
// hmdt: 32
// Each line contains the name of a monkey, a colon, and then the job of that monkey:

// A lone number means the monkey's job is simply to yell that number.
// A job like aaaa + bbbb means the monkey waits for monkeys aaaa and bbbb to yell each of their numbers; the monkey then yells the sum of those two numbers.
// aaaa - bbbb means the monkey yells aaaa's number minus bbbb's number.
// Job aaaa * bbbb will yell aaaa's number multiplied by bbbb's number.
// Job aaaa / bbbb will yell aaaa's number divided by bbbb's number.
// So, in the above example, monkey drzm has to wait for monkeys hmdt and zczc to yell their numbers. Fortunately, both hmdt and zczc have jobs that involve simply yelling a single number, so they do this immediately: 32 and 2. Monkey drzm can then yell its number by finding 32 minus 2: 30.

// Then, monkey sjmn has one of its numbers (30, from monkey drzm), and already has its other number, 5, from dbpl. This allows it to yell its own number by finding 30 multiplied by 5: 150.

// This process continues until root yells a number: 152.

// However, your actual situation involves considerably more monkeys. What number will the monkey named root yell?

package main

import(
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

func isNumber(item []any) bool {
    return len(item) == 1
}

func eval(op string, left, right int) int {
    switch op {
    case "+":
        return left + right
    case "-":
        return left - right
    case "*":
        return left * right
    case "/":
        return left / right
    }
    return -1
}


func parseExp(variable, exp string) ([]any, []string){
    splits := strings.Split(exp, " ")
    return []any{splits[1], variable}, []string{splits[0], splits[2]}
}

func evaluate(constants map[string]int, expressions map[string]string) int {
    if rootVal, ok := constants["root"]; ok {
        return rootVal
    }

    rootExp, ok := expressions["root"]
    if !ok {
        return -1
    }

    operator, variables := parseExp("root", rootExp)
    operators := [][]any{operator}
    fmt.Println("operators:", operators)
    fmt.Println("variables:", variables)

    count := 0
    for len(variables) > 0{
        // check if we can evaluate operators
        for len(operators) > 2 && isNumber(operators[0]) && isNumber(operators[1]) {
            fmt.Println("hello:", operators[2])
            op, variable := operators[2][0].(string), operators[2][1].(string)
            leftOperand, rightOperand := operators[1][0].(int), operators[0][0].(int)
            result := eval(op, leftOperand, rightOperand)
            // update variable result
            constants[variable] = result
            operators = operators[3:]
            operators = append([][]any{[]any{result}}, operators...)
        }
        variable := variables[0]
        variables = variables[1:]
        fmt.Println("variable:", variable)
        // check in constant
        if val, ok := constants[variable]; ok {
            operators = append([][]any{[]any{val}}, operators...)
        } else if exp, ok := expressions[variable]; ok { // not constant
            operator, variable := parseExp(variable, exp)
            fmt.Println("operator:", operator)
            fmt.Println("variable:", variable)
            variables = append(variable, variables...)
            operators = append([][]any{operator}, operators...)
        }
        count += 1
        fmt.Println(operators)
        fmt.Println(variables)
        fmt.Println("----")
    }

    for len(operators) > 2 && isNumber(operators[0]) && isNumber(operators[1]) {
        fmt.Println("hello:", operators[2])
        op, variable := operators[2][0].(string), operators[2][1].(string)
        leftOperand, rightOperand := operators[1][0].(int), operators[0][0].(int)
        result := eval(op, leftOperand, rightOperand)
        // update variable result
        constants[variable] = result
        operators = operators[3:]
        operators = append([][]any{[]any{result}}, operators...)
    }
    fmt.Println("after===")
    fmt.Println(operators)
    fmt.Println(variables)
    return operators[0][0].(int)

}


func parseInput(line string, variables map[string]int, expressions map[string]string) {
    splits := strings.Split(line, ":")
    name := strings.TrimSpace(splits[0])
    val := strings.TrimSpace(splits[1])
    fmt.Println(name, " => ", val)
    if i, err := strconv.Atoi(val); err == nil {
        variables[name] = i
    } else {
        expressions[name] = val
    }
}

//go:embed day21.txt
var day21txt string

func yell() int {
    variables := make(map[string]int)
    expressions := make(map[string]string)
    onLine := func(line string) error {
        if len(line) > 0 {
            parseInput(line, variables, expressions)
        }
        return nil
    }

    if err := util.ReadLinesEmbed(day21txt, onLine); err != nil {
        return -1
    }
    fmt.Println("variables:", variables)
    fmt.Println("expressions:", expressions)
    return evaluate(variables, expressions)
}

func main() {
    fmt.Println("part1:", yell())
}
