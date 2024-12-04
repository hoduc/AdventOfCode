# https://adventofcode.com/2024/day/3
"""
--- Day 3: Mull It Over ---

"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

x`mul(2,4)`%&mul[3,7]!@^do_not_`mul(5,5)`+mul(32,64]then(`mul(11,8)mul(8,5)`)

Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?

--- Part Two ---

As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

    The do() instruction enables future mul instructions.
    The don't() instruction disables future mul instructions.

Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.

For example:

xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))

This corrupted memory is similar to the example from before, but this time the mul(5,5) and mul(11,8) instructions are disabled because there is a don't() instruction before them. The other mul instructions function normally, including the one at the end that gets re-enabled by a do() instruction.

This time, the sum of the results is 48 (2*4 + 8*5).

Handle the new instructions; what do you get if you add up all of the results of just the enabled multiplications?


"""
DONT = "don't()"
DO = "do()"
MUL = "mul("

def parse(s, part2 = False, initial_should_do = True):
    # print(s)
    sum_product = 0
    should_do = initial_should_do
    stack = []
    i = 0
    while i < len(s):
        # print(should_do, s[i:])
        if part2 and i + len(DONT) < len(s) and s[i: i + len(DONT)] == DONT:
            should_do = False
            # print("DONT!!!!")
            i += len(DONT)
            continue
        elif part2 and i + len(DO) < len(s) and s[i: i + len(DO)] == DO:
            should_do = True
            i += len(DO)
            continue
        elif i + len(MUL) < len(s) and s[i: i + len(MUL)] == MUL:
            i += len(MUL)
            while i < len(s) and s[i] != ')':
                if s[i].isdigit():
                    if not stack:
                        stack.append("")
                    stack[-1] += s[i]
                elif s[i] == ",":
                    stack.append("")
                else:
                    # parsing error skip
                    break
                i += 1
            
            executed = False
            if s[i] == ')' and len(stack) == 2 and len(stack[0]) <= 3 and len(stack[-1]) <= 3 and should_do:
                # print(stack, int(stack[0]) * int(stack[-1]))
                sum_product += int(stack[0]) * int(stack[-1])
                executed = True
            # print(stack, s[i] == ')', should_do, executed, sum_product)
            stack = []
            continue
        i += 1
    return sum_product, should_do

assert parse("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")[0] == 161
assert parse("mul(4*, mul(6,9!, ?(12,34)")[0] == 0
assert parse("mul ( 2 , 4 )")[0] == 0
assert parse("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", True)[0] == 48

def part(part2 = False):
    result = 0
    should_do = True
    with open("day3.txt") as f:
        for line in f:
            if not part2:
                should_do = True
                # THIS IS THE FUMBLE RIGHT HERE
                # the dont() can persist between different lines of input
            line_result, should_do = parse(line.strip(), part2, should_do)
            result += line_result
    return result

def part1():
    return part()

def part2():
    return part(True)

# YEP
print(part1()) # 173419328
print(part2()) # 91155532 (WRONG) ATTEMP 2: 90669332 (CORRECT)








