#[
    --- Day 5: If You Give A Seed A Fertilizer ---

You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island is the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with! Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no." His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat. Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry, maybe you can help us with our food production problem. The latest Island Island Almanac just arrived and we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number, but numbers are reused by each category - that is, soil 123 and fertilizer 123 aren't necessarily related to each other.

For example:

seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4

The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55, and 13.

The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into numbers in a destination category. That is, the section that starts with seed-to-soil map: describes how to convert a seed number (the source) to a soil number (the destination). This lets the gardener and his team know which soil to use with which seeds, which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination number one by one, the maps describe entire ranges of numbers that can be converted. Each line within a map contains three numbers: the destination range start, the source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48

The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length, but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. This corresponds to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53 corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51

With this map, you can look up the soil number required for each initial seed number:

    Seed number 79 corresponds to soil number 81.
    Seed number 14 corresponds to soil number 14.
    Seed number 55 corresponds to soil number 57.
    Seed number 13 corresponds to soil number 13.

The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that needs a seed. Using these maps, find the lowest location number that corresponds to any of the initial seeds. To do this, you'll need to convert each seed number through other categories until you can find its corresponding location number. In this example, the corresponding types are:

    Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
    Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
    Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
    Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.

So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed numbers?


--- Part Two ---

Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13

This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?


]#

import strutils
import std/sequtils
import std/sugar
# import std/algorithm


type Mapper = tuple
    src: int
    dst: int
    length: int

proc convert(x: int, mapper: seq[Mapper]): int =
    for m in mapper:
        if x >= m.src and x < m.src + m.length:
            return x - m.src + m.dst
    return x


proc convert(ran: tuple[l: int, r: int], mapper: seq[Mapper]): seq[tuple] =
    var converted_ranges: seq[tuple]
    let q = @[ran]
    while q:
        var added = false
        let
            l = q[0][0]
            r = q[0][1]
        for m in mapper:
            let
                f = m.src
                t = m.src + m.length
            # outside left and outside right
            if (l < f and r < f) or (l >= t and r >= t):
                continue
            # fit in range or equals
            if l >= f and r <= t:
                converted_ranges.add((l - f + m.dst, r - f + m.dst))
                added = true
                break
            #[
                    [---]
                [-----]
                [---]
            ]#
            if l < f and t >= f:
                converted_ranges.add((m.dst, r - f + m.dst))
                q.add((l, f - 1))
                added = true
                break
            #[
                    [---]
                      [---]
                        [---]
            ]#    
            if l <= t and r > t:
                converted_ranges.add((l - f + m.st, m.dst))
                q.add((t + 1, r))
                added = true
                break
        # does not exist in mapping
        if not added:
            converted_ranges.add((l, r))
        q.deletes(0)
    return converted_ranges

proc seedsMappers(fileName: string): (seq[int], seq[seq[Mapper]]) =
    var seeds: seq[int]
    var mappers: seq[seq[Mapper]]
    var mapper: seq[Mapper]
    var lineCount = 0
    for line in lines(fileName):
        # echo(line)
        if line.contains("seeds"):
            # echo(line.splitWhitespace())
            seeds = line.splitWhitespace()[1..^1].map(c => parseInt(c))
        elif line.contains(':'):
            # echo "skipping this==> ", line
            continue
        elif len(line) == 0:
            if lineCount == 1:
                continue
            # echo(mapper)
            mappers.add(mapper)
            mapper = @[]
        else:
            let l = line.splitWhitespace().map(c => parseInt(c))
            mapper.add((l[1], l[0], l[2]))
        lineCount += 1
    mappers.add(mapper)
    return (seeds, mappers)

proc minLoc(seeds: seq[int], mappers: seq[seq[Mapper]]): int =
    echo "seeds:", seeds
    # echo len(mappers)
    var minLocation = -1
    for seed in seeds:
        # echo "===", seed
        var x = seed
        for mapper in mappers:
            x = convert(x, mapper)
            # echo x
        # echo x
        minLocation = if (minLocation == -1): x else: min(minLocation, x)
    return minLocation

proc day5p1*(fileName: string): int =
    let (seeds, mappers) = seedsMappers(fileName)
    return minLoc(seeds, mappers)

proc day5p2*(fileName: string): int =
    let (seeds, mappers) = seedsMappers(fileName)
    # echo mappers.map(mapper => len(mapper))
    echo "pre-seeds:", seeds
    for mapper in mappers:
        echo mapper[0].src, "=>", mapper[^1].src, "|len:", len(mapper)
    var seedsRange: seq[tuple[l: int, r: int]]
    
    var i = 0
    while i < len(seeds):
        seedsRange.add((seeds[i], seeds[i+1]))
        i += 2
    
    # var transformedRanges = []
    # for seedRange in seedsRange:
    #     convert
    return 0
