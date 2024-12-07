# https://adventofcode.com/2024/day/7
"""
--- Day 7: Bridge Repair ---

The Historians take you to a familiar rope bridge over a river in the middle of a jungle. The Chief isn't on this side of the bridge, though; maybe he's on the other side?

When you go to cross the bridge, you notice a group of engineers trying to repair it. (Apparently, it breaks pretty frequently.) You won't be able to cross until it's fixed.

You ask how long it'll take; the engineers tell you that it only needs final calibrations, but some young elephants were playing nearby and stole all the operators from their calibration equations! They could finish the calibrations if only someone could determine which test values could possibly be produced by placing any combination of operators into their calibration equations (your puzzle input).

For example:

190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20

Each line represents a single equation. The test value appears before the colon on each line; it is your job to determine whether the remaining numbers can be combined with operators to produce the test value.

Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: add (+) and multiply (*).

Only three of the above equations can be made true by inserting operators:

    190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but choosing * would give the test value (10 * 19 = 190).
    3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two cause the right side to match the test value: 81 + 40 * 27 and 81 * 40 + 27 both equal 3267 (when evaluated left-to-right)!
    292: 11 6 16 20 can be solved in exactly one way: 11 + 6 * 16 + 20.

The engineers just need the total calibration result, which is the sum of the test values from just the equations that could possibly be true. In the above example, the sum of the test values for the three equations listed above is 3749.

Determine which equations could possibly be true. What is their total calibration result?

--- Part Two ---

The engineers seem concerned; the total calibration result you gave them is nowhere close to being within safety tolerances. Just then, you spot your mistake: some well-hidden elephants are holding a third type of operator.

The concatenation operator (||) combines the digits from its left and right inputs into a single number. For example, 12 || 345 would become 12345. All operators are still evaluated left-to-right.

Now, apart from the three equations that could be made true using only addition and multiplication, the above example has three more equations that can be made true by inserting operators:

    156: 15 6 can be made true through a single concatenation: 15 || 6 = 156.
    7290: 6 8 6 15 can be made true using 6 * 8 || 6 * 15.
    192: 17 8 14 can be made true using 17 || 8 + 14.

Adding up all six test values (the three that could be made before using only + and * plus the new three that can now be made by also using ||) produces the new total calibration result of 11387.

Using your new knowledge of elephant hiding spots, determine which equations could possibly be true. What is their total calibration result?


"""

# HMM JUST COMBINATORICS testing brute force? Lets see we can get part1 done
# OKAY NOW HAS TO DO STRING OPERATIONS

import time

def part1_next_val_fn(i, n, numbers, test_value):
    mul = str(int(n) * int(numbers[i+1]))
    add = str(int(n) + int(numbers[i+1]))
    next_vals = []
    if len(mul) <= len(test_value):
        next_vals += [(i+1, mul)]
    if len(add) <= len(test_value):
        next_vals += [(i+1, add)]
    return next_vals

def part2_next_val_fn(i, n, numbers, test_value):
    concat = n + numbers[i+1]
    next_vals = part1_next_val_fn(i, n, numbers, test_value)
    if len(concat) <= len(test_value):
        next_vals += [(i+1, concat)]
    return next_vals

def part(file_name, next_val_fn = part1_next_val_fn):
    start_time = time.perf_counter()
    total_calibration = 0
    with open(file_name) as f:
        for line in f:
            line_split = line.strip().split()
            test_value = line_split[0][:-1]
            numbers = [e for e in line_split[1:]]
            q = [(0, numbers[0])]
            while q:                
                # MAYBE have some cache so we stop early
                # BUT MAYBE DONT HAVE TO DO IN part1?
                # print(q)
                # OKAY we have to use all numbers maybe
                if q[0][0] == len(numbers) - 1 and q[0][-1] == test_value:
                    break
                i, n = q.pop(0)
                if i < len(numbers) - 1:
                    q += next_val_fn(i, n, numbers, test_value)
            # hmm maybe the test_value is duplicate?
            if q and q[0][-1] == test_value:
                total_calibration += int(test_value)
    end_time = time.perf_counter()
    sec = end_time - start_time
    mins = sec // 60
    sec = sec % 60
    hours = mins // 60
    mins = mins % 60
    print(f'{file_name} took {int(hours):02}:{int(mins):02}:{sec:.2f}')
    return total_calibration

def part1():
    return part("day7.txt")

def part2():
    return part("day7.txt", part2_next_val_fn)


assert part("day7_example2.txt") == 0 # should be 0. Valid values [16, 39, 45, 126]
assert part("day7_example.txt") == 3749
assert part("day7_example.txt", part2_next_val_fn) == 11387

print(part1()) # 1260333054159

# FIRST LETS MEASURE HOW LONG THE FUNCTION TAKE AND SEE IF OUR OPTIMIZATION MAKE IT RUNNABLE FASTER
print(part2()) # 162059698077915 OK IT GONNA NOT FINISH. OKAY IT DOES FINISH. LETS TRY OPTIMIZE IT FURTHER TOO.
# NEW NUMBER: 162042343638683
# WAITING FOR IT TO RUN AND SEE HOW LONG IT TAKES FIRST
# CURRENT: 00:01:25.08 . 1 minutes sth
# O1: 00:00:36.78 . okay 30 seconds faster. but the answer is wrong somehow ....
# THIS IS BECAUSE WE HAVE TO USE ALL THE NUMBER INSTEAD OF BACKING OUT WHEN WE FOUND A MATCH