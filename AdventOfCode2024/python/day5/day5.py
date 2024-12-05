# https://adventofcode.com/2024/day/5
"""
--- Day 5: Print Queue ---

Satisfied with their search on Ceres, the squadron of scholars suggests subsequently scanning the stationery stacks of sub-basement 17.

The North Pole printing department is busier than ever this close to Christmas, and while The Historians continue their search of this historically significant facility, an Elf operating a very familiar printer beckons you over.

The Elf must recognize you, because they waste no time explaining that the new sleigh launch safety manual updates won't print correctly. Failure to update the safety manuals would be dire indeed, so you offer your services.

Safety protocols clearly indicate that new pages for the safety manuals must be printed in a very specific order. The notation X|Y means that if both page number X and page number Y are to be produced as part of an update, page number X must be printed at some point before page number Y.

The Elf has for you both the page ordering rules and the pages to produce in each update (your puzzle input), but can't figure out whether each update has the pages in the right order.

For example:

47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47

The first section specifies the page ordering rules, one per line. The first rule, 47|53, means that if an update includes both page number 47 and page number 53, then page number 47 must be printed at some point before page number 53. (47 doesn't necessarily need to be immediately before 53; other pages are allowed to be between them.)

The second section specifies the page numbers of each update. Because most safety manuals are different, the pages needed in the updates are different too. The first update, 75,47,61,53,29, means that the update consists of page numbers 75, 47, 61, 53, and 29.

To get the printers going as soon as possible, start by identifying which updates are already in the right order.

In the above example, the first update (75,47,61,53,29) is in the right order:

    75 is correctly first because there are rules that put each other page after it: 75|47, 75|61, 75|53, and 75|29.
    47 is correctly second because 75 must be before it (75|47) and every other page must be after it according to 47|61, 47|53, and 47|29.
    61 is correctly in the middle because 75 and 47 are before it (75|61 and 47|61) and 53 and 29 are after it (61|53 and 61|29).
    53 is correctly fourth because it is before page number 29 (53|29).
    29 is the only page left and so is correctly last.

Because the first update does not include some page numbers, the ordering rules involving those missing page numbers are ignored.

The second and third updates are also in the correct order according to the rules. Like the first update, they also do not include every page number, and so only some of the ordering rules apply - within each update, the ordering rules that involve missing page numbers are not used.

The fourth update, 75,97,47,61,53, is not in the correct order: it would print 75 before 97, which violates the rule 97|75.

The fifth update, 61,13,29, is also not in the correct order, since it breaks the rule 29|13.

The last update, 97,13,75,29,47, is not in the correct order due to breaking several rules.

For some reason, the Elves also need to know the middle page number of each update being printed. Because you are currently only printing the correctly-ordered updates, you will need to find the middle page number of each correctly-ordered update. In the above example, the correctly-ordered updates are:

75,47,61,53,29
97,61,53,29,13
75,29,13

These have middle page numbers of 61, 53, and 29 respectively. Adding these page numbers together gives 143.

Of course, you'll need to be careful: the actual list of page ordering rules is bigger and more complicated than the above example.

Determine which updates are already in the correct order. What do you get if you add up the middle page number from those correctly-ordered updates?

"""

from collections import OrderedDict, defaultdict
from functools import cmp_to_key

## LETS PARSE THE INPUT INTO MAP FIRST
def part(file_name):
    middle_page_sum = 0
    update_mode = False
    rules = defaultdict(set)
    correct_updates = []
    def cmp(a, b):
        if b in rules[a]:
            return -1
        elif a in rules[b]:
            return 1
        return 0
    
    with open(file_name) as f:
        for line in f:
            # print("line=", line)
            line = line.strip()
            # parse the ordering rule
            if line:
                if not update_mode:
                    line_split = [int(e) for e in line.split("|")]
                    rules[line_split[0]].add(line_split[-1])
                else:
                    updates = [int(e) for e in line.split(",")]
                    print("===", updates)
                    # determine if the update is in the right order according to rule
                    # each element should be before other in the correct order
                    # LETS DO THE BRUTE FORCE SOLUTION FIRST
                    # BE RIGHT BACK GETTING WATER
                    # ANOTHER THOUGHT JUST NOW IS THAT IF STH COMES BEFORE STH THAT MEANS IT IS LESS THAN THAT
                    # SO IS IT SIMILAR TO SORTING ?
                    # coz the brute force will not finish it seems like i run just now
                    # DOES THAT MEANS THAT IF SORT THE UPDATE BASE ON THE RULE ABOVE AND IF IT COMES OUT TO BE THE SAME ARRAY
                    # THEN THAT MEANS IT CORRECT ORDER ELSE IT IS NOT. LETS TRY THAT
                    sort_updates = sorted(updates, key=cmp_to_key(cmp))
                    print("-->", sort_updates)
                    if sort_updates == updates:
                        print("===correct")
                        middle_page_sum += updates[len(updates) // 2]
                        continue
                    print("===NOPE")
                    
            else:
                update_mode = True
    return middle_page_sum

assert part("day5_example.txt") == 143

def part1():
    return part("day5.txt")


print(part1()) # 6242

# LETS GIT COMMIT THIS