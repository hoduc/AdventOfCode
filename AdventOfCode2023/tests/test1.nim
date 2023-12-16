# This is just an example to get you started. You may wish to put all of your
# tests into a single file, or separate them into multiple `test1`, `test2`
# etc. files (better names are recommended, just make sure the name starts with
# the letter 't').
#
# To run these tests, simply execute `nimble test`.

import unittest

import AdventOfCode2023pkg/submodule
# import AdventOfCode2023pkg/day1/day1
# import AdventOfCode2023pkg/day2/day2
# import AdventOfCode2023pkg/day3/day3
# import AdventOfCode2023pkg/day4/day4
# import AdventOfCode2023pkg/day5/day5
# import AdventOfCode2023pkg/day6/day6
# import AdventOfCode2023pkg/day7/day7
# import AdventOfCode2023pkg/day8/day8
# import AdventOfCode2023pkg/day9/day9
# import AdventOfCode2023pkg/day10/day10
import AdventOfCode2023pkg/day11/day11
# import AdventOfCode2023pkg/day13/day13
# import AdventOfCode2023pkg/day14/day14
import AdventOfCode2023pkg/day16/day16

test "correct welcome":
  check getWelcomeMessage() == "Hello, World!"


# TODO: Where and how to resolve all these paths for input file?
# test "day1":
#     check day1("day1_sample.txt") == 142
#     check day1("day1_1.txt") == 55123
#     check day1("day1_2_sample.txt", false) == 281
#     check day1("day1_2_sample2.txt", false) == 162
#     check day1("day1_2_sample3.txt", false) == 21
#     check day1("day1_2_sample4.txt", false) == 33
#     check day1("day1_2_sample5.txt", false) == 18
#     check day1("day1_2.txt", false) == 55260
  
# test "day2":
#     check day2("day2_sample.txt", day2Part1) == 8
#     check day2("day2_1.txt", day2Part1) == 2795
#     check day2("day2_sample.txt", day2Part2) == 2286
#     check day2("day2_1.txt", day2Part2) == 75561

# test "day3":
#     check day3("day3_sample.txt") == (4361, 467835)
#     check day3("day3_sample2.txt") == (4, 0)
#     check day3("day3_sample3.txt") == (9, 20)
#     check day3("day3.txt") == (535235, 79844424)

# test "day4":
#     check day4("day4_sample.txt") == (13, 30)
#     check day4("day4.txt") == (18619, 8063216)

# test "day5":
#     check day5p1("day5_sample.txt") == 35
#     check day5p1("day5.txt") == 462648396
    
#     check day5p2("day5_sample.txt") == 46
#     check day5p2("day5.txt") == 2520479

# test "day6":
#     check day6p1("day6_sample.txt") == 288
#     check day6p1("day6.txt") == 1084752
#     check day6p2("day6_sample.txt") == 71503
#     check day6p2("day6.txt") == 28228952

# test "day7":
#   check day71("day7_sample.txt") == 6440
#   check day71("day7.txt") == 246163188
#   check day72("day7_sample.txt") == 5905
#   check day72("day7.txt") == 245794069

# test "day8":
#     check day81("day8_sample.txt") == 2
#     check day81("day8_sample2.txt") == 6
#     check day81("day8.txt") == 20093
#     check day82("day8_sample3.txt") == 6
#     check day82("day8.txt") ==  22103062509257
  
# test "day9":
#   check day91("day9_sample.txt") == 114
#   check day91("day9.txt") == 1641934234
#   check day92("day9_sample.txt") == 2
#   check day92("day9.txt") == 2

# test "day10":
  # check day101("day10_sample.txt") == 4
  # check day101("day10_sample2.txt") == 8
  # check day101("day10.txt") == 6786
  # check day102("day10_sample3.txt") == 4
  # check day102("day10_sample4.txt") == 4
  # check day102("day10_sample5.txt") == 8
  # check day102("day10_sample6.txt") == 10
  # check day102("day10.txt") == 8

test "day11":
  check day111("day11_sample.txt") == 374
  check day111("day11_sample2.txt") == 3
  check day111("day11_sample3.txt") == 7
  check day111("day11.txt") == 10422930
  check day112("day11_sample.txt", 10) == 1030
  check day112("day11_sample.txt", 100) == 8410
  check day112("day11.txt", 1000000) == 699909023130

# test "day13":
#   check day13("day13_sample.txt") == 405
#   check day13("day13.txt") == 34889

# test "day14":
#   check day14("day14_sample.txt") == 136 
#   check day14("day14.txt") == 136


# ======= measure time
# write a template for this
# import std/monotimes
# import std/times
# import std/os
# var startTime = getMonoTime()
# sleep(100)
# var endTime = getMonoTime()
# var elapsedTime: Duration = endTime - startTime
# echo " elapsed:", elapsedTime.inMilliseconds()

# test "day16":
#   check day161("day16_sample.txt") == 46
#   check day161("day16.txt") == 7210
#   check day162("day16_sample.txt") == 51
#   # TODO: fix runtime. This took : 68422 milliseconds
#   check day162("day16.txt") == 7673