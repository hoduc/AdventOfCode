#[
--- Day 19: Aplenty ---

The Elves of Gear Island are thankful for your help and send you on your way. They even have a hang glider that someone stole from Desert Island; since you're already going that direction, it would help them a lot if you would use it to get down there and return it to them.

As you reach the bottom of the relentless avalanche of machine parts, you discover that they're already forming a formidable heap. Don't worry, though - a group of Elves is already here organizing the parts, and they have a system.

To start, each part is rated in each of four categories:

    x: Extremely cool looking
    m: Musical (it makes a noise when you hit it)
    a: Aerodynamic
    s: Shiny

Then, each part is sent through a series of workflows that will ultimately accept or reject the part. Each workflow has a name and contains a list of rules; each rule specifies a condition and where to send the part if the condition is true. The first rule that matches the part being considered is applied immediately, and the part moves on to the destination described by the rule. (The last rule in each workflow has no condition and always applies if reached.)

Consider the workflow ex{x>10:one,m<20:two,a>30:R,A}. This workflow is named ex and contains four rules. If workflow ex were considering a specific part, it would perform the following steps in order:

    Rule "x>10:one": If the part's x is more than 10, send the part to the workflow named one.
    Rule "m<20:two": Otherwise, if the part's m is less than 20, send the part to the workflow named two.
    Rule "a>30:R": Otherwise, if the part's a is more than 30, the part is immediately rejected (R).
    Rule "A": Otherwise, because no other rules matched the part, the part is immediately accepted (A).

If a part is sent to another workflow, it immediately switches to the start of that workflow instead and never returns. If a part is accepted (sent to A) or rejected (sent to R), the part immediately stops any further processing.

The system works, but it's not keeping up with the torrent of weird metal shapes. The Elves ask if you can help sort a few parts and give you the list of workflows and some part ratings (your puzzle input). For example:

px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}

The workflows are listed first, followed by a blank line, then the ratings of the parts the Elves would like you to sort. All parts begin in the workflow named in. In this example, the five listed parts go through the following workflows:

    {x=787,m=2655,a=1222,s=2876}: in -> qqz -> qs -> lnx -> A
    {x=1679,m=44,a=2067,s=496}: in -> px -> rfg -> gd -> R
    {x=2036,m=264,a=79,s=2244}: in -> qqz -> hdj -> pv -> A
    {x=2461,m=1339,a=466,s=291}: in -> px -> qkq -> crn -> R
    {x=2127,m=1623,a=2188,s=1013}: in -> px -> rfg -> A

Ultimately, three parts are accepted. Adding up the x, m, a, and s rating for each of the accepted parts gives 7540 for the part with x=787, 4623 for the part with x=2036, and 6951 for the part with x=2127. Adding all of the ratings for all of the accepted parts gives the sum total of 19114.

Sort through all of the parts you've been given; what do you get if you add together all of the rating numbers for all of the parts that ultimately get accepted?


]#
# could try object variant but for now obect is fine

import std/sequtils
import std/math
import std/strutils
import std/tables

type Op = ref object of RootObj
    l, op, o: string
    r: int
method eval(self: Op, ir: int): (bool, string) {.base} = (true, self.o)
method `$`(self: Op): string {.base} = $self.l & " " & $self.op & " " & $self.r & "->" & $self.o

type TwoOp = ref object of Op
method eval(self: TwoOp, ir: int): (bool, string) =
    # echo "compare:", ir, " & ", self.r, " | op:", self.op
    if self.op == ">":
        return if ir > self.r: (true, self.o) else: (false, "")
    elif self.op == "<":
        return if ir < self.r: (true, self.o) else: (false, "")


proc day19*(fileName: string): int =
    var workflows = initTable[string, seq[Op]]()
    var ratings: seq[Table[string, int]] = @[]
    var shouldParseRatings = false
    for line in lines(fileName):
        if len(line) == 0:
            shouldParseRatings = true
            continue
        if not shouldParseRatings:
            var name = ""
            var i = 0
            while line[i] != '{':
                name &= line[i]
                i += 1
            if name notin workflows:
                workflows[name] = @[]
            for cond in line[i+1..^2].split(","):
                # echo cond
                var cs = cond.split(":")
                # echo "cs:", cs, len(cs)
                if len(cs) == 1:
                    workflows[name].add(Op(l: "", op: "", r: -1, o: cs[0]))
                else:
                    workflows[name].add(TwoOp(l: $cs[0][0], op: $cs[0][1], r: parseInt(cs[0][2..^1]), o: cs[1]))
        else:
            ratings.add(initTable[string, int]())
            for statement in line[1..^2].split(","):
                # echo statement
                let statementSplit = statement.split("=")
                let (l, r) = (statementSplit[0], parseInt(statementSplit[1]))
                ratings[^1][l] = r
    # echo workflows.keys.toSeq
    # echo ratings

    var ratingSum = 0
    # var count = 0
    for rating in ratings:
        # if count >= 2:
        #     break
        # echo "====="
        # echo rating
        var name = "in"
        var innerCount = 0
        while name != "A" and name != "R":
            # echo name
            let works = workflows[name]
            for op in works:
                # echo op, "->"
                if op.l in rating:
                    let (er, n) = op.eval(rating[op.l])
                    # echo "res:", (er, n)
                    if er:
                        name = n
                        break
                else:
                    let (_, n) = op.eval(-1)
                    # echo ("res:", n)
                    name = n
                    break
                innerCount += 1
        if name == "A":
            # echo "->", sum(rating.values.toSeq)
            ratingSum += sum(rating.values.toSeq)
        # else:
        #     echo "->", 0
        # count += 1
    return ratingSum