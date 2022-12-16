// --- Day 15: Beacon Exclusion Zone ---
// You feel the ground rumble again as the distress signal leads you to a large network of subterranean tunnels. You don't have time to search them all, but you don't need to: your pack contains a set of deployable sensors that you imagine were originally built to locate lost Elves.

// The sensors aren't very powerful, but that's okay; your handheld device indicates that you're close enough to the source of the distress signal to use them. You pull the emergency sensor system out of your pack, hit the big button on top, and the sensors zoom off down the tunnels.

// Once a sensor finds a spot it thinks will give it a good reading, it attaches itself to a hard surface and begins monitoring for the nearest signal source beacon. Sensors and beacons always exist at integer coordinates. Each sensor knows its own position and can determine the position of a beacon precisely; however, sensors can only lock on to the one beacon closest to the sensor as measured by the Manhattan distance. (There is never a tie where two beacons are the same distance to a sensor.)

// It doesn't take long for the sensors to report back their positions and closest beacons (your puzzle input). For example:

// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
// Sensor at x=9, y=16: closest beacon is at x=10, y=16
// Sensor at x=13, y=2: closest beacon is at x=15, y=3
// Sensor at x=12, y=14: closest beacon is at x=10, y=16
// Sensor at x=10, y=20: closest beacon is at x=10, y=16
// Sensor at x=14, y=17: closest beacon is at x=10, y=16
// Sensor at x=8, y=7: closest beacon is at x=2, y=10
// Sensor at x=2, y=0: closest beacon is at x=2, y=10
// Sensor at x=0, y=11: closest beacon is at x=2, y=10
// Sensor at x=20, y=14: closest beacon is at x=25, y=17
// Sensor at x=17, y=20: closest beacon is at x=21, y=22
// Sensor at x=16, y=7: closest beacon is at x=15, y=3
// Sensor at x=14, y=3: closest beacon is at x=15, y=3
// Sensor at x=20, y=1: closest beacon is at x=15, y=3
// So, consider the sensor at 2,18; the closest beacon to it is at -2,15. For the sensor at 9,16, the closest beacon to it is at 10,16.

// Drawing sensors as S and beacons as B, the above arrangement of sensors and beacons looks like this:

//                1    1    2    2
//      0    5    0    5    0    5
//  0 ....S.......................
//  1 ......................S.....
//  2 ...............S............
//  3 ................SB..........
//  4 ............................
//  5 ............................
//  6 ............................
//  7 ..........S.......S.........
//  8 ............................
//  9 ............................
// 10 ....B.......................
// 11 ..S.........................
// 12 ............................
// 13 ............................
// 14 ..............S.......S.....
// 15 B...........................
// 16 ...........SB...............
// 17 ................S..........B
// 18 ....S.......................
// 19 ............................
// 20 ............S......S........
// 21 ............................
// 22 .......................B....
// This isn't necessarily a comprehensive map of all beacons in the area, though. Because each sensor only identifies its closest beacon, if a sensor detects a beacon, you know there are no other beacons that close or closer to that sensor. There could still be beacons that just happen to not be the closest beacon to any sensor. Consider the sensor at 8,7:

//                1    1    2    2
//      0    5    0    5    0    5
// -2 ..........#.................
// -1 .........###................
//  0 ....S...#####...............
//  1 .......#######........S.....
//  2 ......#########S............
//  3 .....###########SB..........
//  4 ....#############...........
//  5 ...###############..........
//  6 ..#################.........
//  7 .#########S#######S#........
//  8 ..#################.........
//  9 ...###############..........
// 10 ....B############...........
// 11 ..S..###########............
// 12 ......#########.............
// 13 .......#######..............
// 14 ........#####.S.......S.....
// 15 B........###................
// 16 ..........#SB...............
// 17 ................S..........B
// 18 ....S.......................
// 19 ............................
// 20 ............S......S........
// 21 ............................
// 22 .......................B....
// This sensor's closest beacon is at 2,10, and so you know there are no beacons that close or closer (in any positions marked #).

// None of the detected beacons seem to be producing the distress signal, so you'll need to work out where the distress beacon is by working out where it isn't. For now, keep things simple by counting the positions where a beacon cannot possibly be along just a single row.

// So, suppose you have an arrangement of beacons and sensors like in the example above and, just in the row where y=10, you'd like to count the number of positions a beacon cannot possibly exist. The coverage from all sensors near that row looks like this:

//                  1    1    2    2
//        0    5    0    5    0    5
//  9 ...#########################...
// 10 ..####B######################..
// 11 .###S#############.###########.
// In this example, in the row where y=10, there are 26 positions where a beacon cannot be present.

// Consult the report from the sensors you just deployed. In the row where y=2000000, how many positions cannot contain a beacon?

package main

import(
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "sort"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type Location struct {
    x int
    y int
}

type Interval struct {
    start int
    end int
}

func dist(a, b Location) int {
    return util.Abs(a.x - b.x).(int) + util.Abs(a.y - b.y).(int)
}

func updateCoverage(rows map[int][]Interval, sensor, beacon Location) {
    d := dist(sensor, beacon)
    startY, endY := sensor.y - d, sensor.y + d
    // fmt.Println("d=", d)
    // fmt.Println("startY=", startY, "endY=", endY)
    for row, pad := startY, 0; row <= endY; row++ {
        start, end := sensor.x - pad, sensor.x + pad
        if _, ok := rows[row]; !ok {
            rows[row] = []Interval{}
        }
        rows[row] = append(rows[row], Interval{start: start, end: end})
        // fmt.Print(row, "|", pad, endY - row,  "=>", start, end, " ; ")
        if endY - row <= d {
            pad -= 1
        } else {
            pad += 1
        }
    }
    // fmt.Println()
}

func parseCoord(line string) (int, error) {
    i, err := strconv.Atoi(strings.Split(strings.TrimSpace(line), "=")[1])
    if err != nil {
        return -1, err
    }
    return i, err
}

// Sensor at x=***, y=***
// closest beacon is at x=**, y=***
func parseLocation(line string) (Location, error) {
    splits := strings.Split(strings.TrimSpace(strings.Split(strings.TrimSpace(line), "at")[1]), ",")
    x, xerr := parseCoord(splits[0])
    if xerr != nil {
        return Location{}, xerr
    }
    y, yerr := parseCoord(splits[1])
    if yerr != nil {
        return Location{}, yerr
    }
    return Location{x: x, y: y}, nil
}

// Sensor at x=***, y=***: closest beacon is at x=**, y=***
func parseSensorBeacon(line string) (Location, Location, error) {
    splits := strings.Split(line, ":")
    beacon, berr := parseLocation(splits[0])
    if berr != nil {
        return Location{}, Location{}, berr
    }
    sensor, serr := parseLocation(splits[1])
    if serr != nil {
        return Location{}, Location{}, serr
    }
    return beacon, sensor, nil
}

func mergeIntervals(a Interval, b Interval) Interval {
    if a == b {
        return a
    }
    if a.start <= b.start {
        return Interval{start: a.start, end: util.Max(a.end, b.end).(int)}
    }
    return mergeIntervals(b, a)
}

func notOverlapped(a Interval, b Interval) bool {
    return a.start < b.start && a.end < b.start || b.start < a.start && b.end < a.start
}

func sortIntervalsAtRow(rows map[int][]Interval, row int) []Interval{
    // get the []interval
    intervals := rows[row]
    // fmt.Println("Before:", intervals)
    // sort them
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })
    // fmt.Println("AfteR:", intervals)

    // merge intervals
    newIntervals := []Interval{intervals[0]}
    intervals = intervals[1:]
    for len(intervals) > 0 {
        first, second := newIntervals[len(newIntervals)-1], intervals[0]
        // fmt.Println(first, second)
        // fmt.Println(first.start, first.end)
        // fmt.Println(second.start, second.end)
        if notOverlapped(first, second) {
            // fmt.Println("1")
            newIntervals = append(newIntervals, intervals[0])
        } else {
            // fmt.Println("2")
            newIntervals = append(newIntervals[:len(newIntervals)-1], mergeIntervals(first, second))
        }
        intervals = intervals[1:]
        // fmt.Println(newIntervals)
    }
    rows[row] = newIntervals
    return rows[row]
}

func updateLocationsRows(rows map[int]int, loc Location) {
    if _, ok := rows[loc.y]; !ok {
        rows[loc.y] = 1
    }
}

func getLocationsRows(rows map[int]int, y int) int {
    if _, ok := rows[y]; !ok {
        return 0
    }
    return rows[y]
}


//go:embed day15.txt
var day15txt string

//go:embed day15_ext.txt
var day15txtExt string

type senseFn func (map[int]int, map[int]int, map[int][]Interval) int

func beaconSensor(sense senseFn) int {
    lineNo := 0
    sensorsRows := make(map[int]int)
    beaconsRows := make(map[int]int)
    rows := make(map[int][]Interval)
    onLine := func(line string) error {
        // fmt.Println(lineNo, line)
        if len(line) > 0 {
            sensor, beacon, err := parseSensorBeacon(line)
            if err != nil {
                return err
            }
            fmt.Printf(">> Sensor [%v], Beacon [%v]\n", sensor, beacon)
            updateCoverage(rows, sensor, beacon)
            updateLocationsRows(sensorsRows, sensor)
            updateLocationsRows(beaconsRows, beacon)

        }
        lineNo += 1
        return nil
    }

    if err := util.ReadLinesEmbed(day15txt, onLine); err != nil {
        return -1
    }

    return sense(sensorsRows, beaconsRows, rows)
}


func part1(sensorsRows map[int]int, beaconsRows map[int]int, rows map[int][]Interval) int {
    inputRow := -1
    readInputRow := func(line string) error {
        if len(line) > 0 {
            // fmt.Println("line:", line)
            val, err := strconv.Atoi(line)
            if err != nil {
                return err
            }
            inputRow = val
        }
        return nil
    }


    if err := util.ReadLinesEmbedLineNumber(day15txtExt, 0, readInputRow); err != nil {
        return -1
    }

    if inputRow == -1 {
        fmt.Println("Failed to parse input Row")
        return -1
    }


    // countIntervalsRow := countIntervalsLengthAtRow(rows, inputRow)
    sortedIntervalsAtRow := sortIntervalsAtRow(rows, inputRow)
    countIntervalsRow := 0
    for _, val := range sortedIntervalsAtRow {
        fmt.Println(val, val.end - val.start)
        countIntervalsRow += util.Abs(val.end - val.start).(int) + 1
    }
    beaconsRow := getLocationsRows(beaconsRows, inputRow)
    sensorsRow := getLocationsRows(sensorsRows, inputRow)
    fmt.Println("sensorsRows:", sensorsRows)
    fmt.Println("beaconsRows:", beaconsRows)
    fmt.Printf("countIntervals: %v, beaconsRow: %v, sensorsRow: %v\n", countIntervalsRow, beaconsRow, sensorsRow)
    return  countIntervalsRow - beaconsRow - sensorsRow
}

func part2(sensorsRows map[int]int, beaconsRows map[int]int, rows map[int][]Interval) int {
    f := -1
    for y, _ := range rows {
        sortedIntervalsAtRow := sortIntervalsAtRow(rows, y)
        if len(sortedIntervalsAtRow) > 1 {
            for i := 1; i < len(sortedIntervalsAtRow); i++ {
                if sortedIntervalsAtRow[i].start - sortedIntervalsAtRow[i-1].end == 2 && y > 0{
                    x := sortedIntervalsAtRow[i-1].end + 1
                    if x < 0 || x > 4000000 || y < 0 || y > 4000000 {
                        continue
                    }
                    f = x*4000000 + y
                    fmt.Println("candidate:", x, y, f)
                    break
                }
            }
        }
    }
    return f
}

func main() {
    fmt.Println("part1:", beaconSensor(part1))
    fmt.Println("part2:", beaconSensor(part2))
}
