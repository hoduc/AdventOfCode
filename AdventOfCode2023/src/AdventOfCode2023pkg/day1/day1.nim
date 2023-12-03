#[
    --- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

--- Part Two ---

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?


]#

import strutils
import unittest

proc day1*(fileName: string, numberOnly: bool = true): int =
    var calibrationTotal = 0
    for line in lines(fileName):
        var calibration = 0
        var digit = 0
        var i = 0
        # echo "==="
        while i < len(line):
            var inc = 1
            if isDigit(line[i]):
                digit = int(line[i]) - int('0')
            elif not numberOnly:
                for j, s in @["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"].pairs:
                    # 1 -> -1
                    # 2 -> -1
                    # 3 -> -1
                    # 5 -> -1
                    # 7 -> -1
                    # 8 -> -1
                    if i + len(s) - 1 < len(line) and line[i .. i + len(s) - 1] == s:
                        inc = if (j + 1 in [1,2,3,5,7,8]): len(s) - 1 else: len(s)
                        digit = j + 1
                        break
            # echo "digit:", digit    
            if calibration == 0 and digit > 0:
                calibration += digit * 10
            # echo calibration
            i += inc
        calibration += digit
        # echo line, " | => :", calibration
        calibrationTotal += calibration
    return calibrationTotal

test "day1_sample.txt":
    check day1("day1_sample.txt") == 142
# test "day1_1.txt":
#     check day1("day1_1.txt") == 55123
# test "day1_2_sample.txt":
#     check day1("day1_2_sample.txt", false) == 281
# test "day1_2_sample2.txt":
#     check day1("day1_2_sample2.txt", false) == 162
# test "day1_2_sample3.txt":
#     check day1("day1_2_sample3.txt", false) == 21
# test "day1_2_sample4.txt":
#     check day1("day1_2_sample4.txt", false) == 33
# test "day1_2_sample5.txt":
#     check day1("day1_2_sample5.txt", false) == 18
# test "day1_2.txt":
#     check day1("day1_2.txt", false) == 55260
