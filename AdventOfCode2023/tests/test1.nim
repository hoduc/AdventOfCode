# This is just an example to get you started. You may wish to put all of your
# tests into a single file, or separate them into multiple `test1`, `test2`
# etc. files (better names are recommended, just make sure the name starts with
# the letter 't').
#
# To run these tests, simply execute `nimble test`.

import unittest

import AdventOfCode2023pkg/submodule
import AdventOfCode2023pkg/day1/day1
import AdventOfCode2023pkg/day2/day2
import AdventOfCode2023pkg/day3/day3
import AdventOfCode2023pkg/day4/day4

test "correct welcome":
  check getWelcomeMessage() == "Hello, World!"


# TODO: Where and how to resolve all these paths for input file?
test "day1":
    check day1("day1_sample.txt") == 142
    check day1("day1_1.txt") == 55123
    check day1("day1_2_sample.txt", false) == 281
    check day1("day1_2_sample2.txt", false) == 162
    check day1("day1_2_sample3.txt", false) == 21
    check day1("day1_2_sample4.txt", false) == 33
    check day1("day1_2_sample5.txt", false) == 18
    check day1("day1_2.txt", false) == 55260
  
test "day2":
    check day2("day2_sample.txt", day2Part1) == 8
    check day2("day2_1.txt", day2Part1) == 2795
    check day2("day2_sample.txt", day2Part2) == 2286
    check day2("day2_1.txt", day2Part2) == 75561

test "day3":
    check day3("day3_sample.txt") == (4361, 467835)
    check day3("day3_sample2.txt") == (4, 0)
    check day3("day3_sample3.txt") == (9, 20)
    check day3("day3.txt") == (535235, 79844424)

test "day4":
    check day4("day4_sample.txt") == 13
    check day4("day4.txt") == 18619