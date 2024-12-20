# https://adventofcode.com/2024/day/2
"""
--- Day 2: Red-Nosed Reports ---

Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

    The levels are either all increasing or all decreasing.
    Any two adjacent levels differ by at least one and at most three.

In the example above, the reports can be found safe or unsafe by checking those rules:

    7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
    1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
    9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
    1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
    8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
    1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?

--- Part Two ---

The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

    7 6 4 2 1: Safe without removing any level.
    1 2 7 8 9: Unsafe regardless of which level is removed.
    9 7 6 2 1: Unsafe regardless of which level is removed.
    1 3 2 4 5: Safe by removing the second level, 3.
    8 6 4 4 1: Safe by removing the third level, 4.
    1 3 6 7 9: Safe without removing any level.

Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?


"""

# assuming l1 occurs before l2

# false if decreasing
def are_levels_increasing(l1, l2):
    return l2 > l1

def are_levels_differed_safe(l1, l2, is_increasing = True):
    return are_levels_increasing(l1, l2) == is_increasing and abs(l1 - l2) >= 1 and abs(l1 - l2) <= 3

# increasing with before increasing
assert are_levels_differed_safe(1, 3, True) == True
# increasing with before decreasing
assert are_levels_differed_safe(1, 3, False) == False
# plateau with before increasing
assert are_levels_differed_safe(1, 1, True) == False
# plateau with before decreasing
assert are_levels_differed_safe(1, 1, False) == False
# decreasing with before increasing
assert are_levels_differed_safe(1, 0, True) == False
# decreasing with before decreasing
assert are_levels_differed_safe(1, 0, False) == True
# decreasing with before decreasing but outside range of [1, 3]
assert are_levels_differed_safe(5, 1, False) == False
assert are_levels_differed_safe(1, -5, False) == False

def is_report_safe(levels):
    # FORGOT THE FIRST TWO
    is_increasing = are_levels_increasing(levels[0], levels[1])
    if not are_levels_differed_safe(levels[0], levels[1], is_increasing):
        return False
    for i in range(1, len(levels) - 1):
        if not are_levels_differed_safe(levels[i], levels[i+1], is_increasing):
            return False
    return True

assert is_report_safe([7, 6, 4, 2, 1]) == True
assert is_report_safe([1, 2, 7, 8 ,9]) == False
assert is_report_safe([9, 7, 6, 2, 1]) == False
assert is_report_safe([1, 3, 2, 4, 5]) == False
assert is_report_safe([8, 6, 4, 4, 1]) == False
assert is_report_safe([1, 3, 6, 7, 9]) == True


def part(is_report_safe_fn = is_report_safe):
    safe = 0
    with open('day2.txt') as f:
        for line in f:
            report = [int(l) for l in line.strip().split()]
            report_safe = is_report_safe_fn(report)
            # print(report, report_safe)
            if report_safe:
                safe += 1
    return safe



def part1():
    return part()


def is_report_safe_dampener(report):
    report_safe = is_report_safe(report)
    remove_index = 0
    while not report_safe and remove_index < len(report):
        report_safe = is_report_safe(report[:remove_index] + report[remove_index + 1:])
        remove_index += 1
    return report_safe



assert is_report_safe_dampener([7, 6, 4, 2, 1]) == True
assert is_report_safe_dampener([1, 2, 7, 8, 9]) == False
assert is_report_safe_dampener([9, 7, 6, 2, 1]) == False
assert is_report_safe_dampener([1, 3, 2, 4, 5]) == True
assert is_report_safe_dampener([8, 6, 4, 4, 1]) == True
assert is_report_safe_dampener([1, 3, 6, 7, 9]) == True


def part2():
    return part(is_report_safe_dampener)

print(part1())
print(part2())
