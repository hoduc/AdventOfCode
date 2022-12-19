package main

import(
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type Point struct {
    x float64
    y float64
    z float64
}


type Side struct {
    topLeft Point
    topRight Point
    bottomLeft Point
    bottomRight Point
}

type Cube struct {
    sides []Side
}

func newPoint(x, y, z float64) Point {
    return Point{x: x, y: y, z: z}
}

func topSide(x, y, z float64) Side {
    return Side{
        topLeft: newPoint(x - 0.5, y + 0.5, z - 0.5),
        topRight: newPoint(x + 0.5, y + 0.5, z - 0.5),
        bottomLeft: newPoint(x - 0.5, y + 0.5, z + 0.5),
        bottomRight: newPoint(x + 0.5, y + 0.5, z + 0.5),
    }
}

func downSide(x, y, z float64) Side {
    return topSide(x, y - 1, z)
}

// TODO
func leftSide(x, y, z float64) Side {
    return Side {
        topLeft: newPoint(x - 0.5, y + 0.5, z + 0.5),
        topRight: newPoint(x - 0.5, y + 0.5, z - 0.5),
        bottomLeft: newPoint(x - 0.5, y - 0.5, z + 0.5),
        bottomRight: newPoint(x - 0.5, y - 0.5, z - 0.5),
    }
}

func rightSide(x, y, z float64) Side {
    return leftSide(x + 1, y, z)
}

// TODO
func frontSide(x, y, z float64) Side {
    return Side {
        topLeft: newPoint(x - 0.5, y + 0.5, z + 0.5),
        topRight: newPoint(x + 0.5, y + 0.5, z + 0.5),
        bottomLeft: newPoint(x - 0.5, y - 0.5, z + 0.5),
        bottomRight: newPoint(x + 0.5, y - 0.5, z + 0.5),
    }
}

func backSide(x, y, z float64) Side {
    return frontSide(x, y, z - 1)
}


func newCube(x, y, z float64) Cube {
    return Cube{sides: []Side{
        topSide(x, y, z),
        downSide(x, y, z),
        leftSide(x, y, z),
        rightSide(x, y, z),
        frontSide(x, y, z),
        backSide(x, y, z),
    }}
}

func lineToCube(line string) (Cube, error){
    splits := strings.Split(line, ",")
    var cube Cube
    var x, y, z int
    var err error
    x, err = strconv.Atoi(splits[0])
    if err != nil {
        return cube, err
    }
    y, err = strconv.Atoi(splits[1])
    if err != nil {
        return cube, err
    }
    z, err = strconv.Atoi(splits[2])
    if err != nil {
        return cube, err
    }
    // fmt.Println("x:", x, "y:", y, "z:", z)
    cube = newCube(float64(x), float64(y), float64(z))
    return cube, nil
}

//go:embed day18.txt
var day18txt string

func surfaceArea() int {
    cubes := []Cube{}
    sidesFreq := make(map[Side]float64)
    onLine := func(line string) error {
        fmt.Println(line)
        if len(line) > 0 {
            cube, err := lineToCube(line)
            // fmt.Println("cube:", cube)
            if err != nil {
                return err
            }
            cubes = append(cubes, cube)
            for _, side := range cube.sides {
                if _, ok := sidesFreq[side]; !ok {
                    sidesFreq[side] = 0
                }
                sidesFreq[side] = sidesFreq[side] + 1
            }
        }
        return nil
    }

    if err := util.ReadLinesEmbed(day18txt, onLine); err != nil {
        return -1
    }
    fmt.Println(sidesFreq, len(sidesFreq))
    fmt.Println("len(cubes):", len(cubes))
    surfaceArea := 0
    for _, cube := range cubes {
        exposedSides := 0
        for _, side := range cube.sides {
            f, ok := sidesFreq[side]
            if !ok || f <= 1 {
                exposedSides += 1
            }
        }
        fmt.Println(cube, exposedSides)
        surfaceArea += exposedSides
    }
    return surfaceArea
}

func main() {
    fmt.Println("part1:", surfaceArea())
}
