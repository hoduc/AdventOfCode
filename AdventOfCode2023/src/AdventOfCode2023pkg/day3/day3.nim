#[
--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?


--- Part Two ---

The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?

]#

import strutils
import std/tables
import unittest

type NonEmptyCell = tuple
    n: int
    r: int
    st: int
    en: int

proc inBound(v: int, b: int): bool =
    return v >= 0 and v < b

proc readBoardsNumbersGears(fileName: string): (seq[string], seq[NonEmptyCell], TableRef[tuple[r: int, c: int], seq[int]]) =
    var boards: seq[string]
    var numbers: seq[NonEmptyCell]
    let gears = newTable[tuple[r: int, c: int], seq[int]]()
    var row = 0
    for line in lines(fileName):
        var start = -1
        for i, c in line.pairs:
            if c == '*':
                gears[(row, i)] = @[]
            if c.isDigit:
                if start == -1:
                    start = i
            else: #  boundary
                if start != -1:
                    numbers.add((parseInt(line[start .. i - 1]), row, start, i - 1))
                start = -1
        if start != -1:
            numbers.add((parseInt(line[start .. len(line) - 1]), row, start, len(line) - 1))
        boards.add(line)
        row += 1
    return (boards, numbers, gears)

proc day3*(fileName: string): (int, int) =
    let (boards, numbers, gears) = readBoardsNumbersGears(fileName)
    let
        rb = len(boards)
        cb = len(boards[0])
    var
        partSum = 0
        gearSum = 0
    for number in numbers:
        for r in number.r - 1 .. number.r + 1:
            for c in number.st - 1 .. number.en + 1:
                if inBound(r, rb) and inBound(c, cb) and not boards[r][c].isDigit and boards[r][c] != '.':
                    partSum += number.n
                    if (r, c) in gears:
                        gears[(r,c)].add(number.n)
    for _, parts in pairs(gears):
        if len(parts) != 2:
            continue
        gearSum += parts[0] * parts[1]
    return (partSum, gearSum)

    
# test "day3_sample.txt":
#     check day3("day3_sample.txt") == (4361, 467835)
# test "day3_sample2.txt":
#     check day3("day3_sample2.txt") == (4, 0)
# test "day3_sample3.txt":
#     check day3("day3_sample3.txt") == (9, 20)
# test "day3.txt":
#     check day3("day3.txt") == (535235, 79844424)