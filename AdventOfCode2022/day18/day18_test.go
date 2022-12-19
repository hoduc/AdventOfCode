package main

import (
    "testing"
)


func testCmp(t *testing.T , x, y, z float64, expected Cube) {
    actual := newCube(x, y, z)
    if len(actual.sides) != len(expected.sides) {
        t.Fatalf(`expected %v but got %v`, len(expected.sides), len(actual.sides))
    }

    for i := 0; i < len(expected.sides); i++ {
        expectedSide := expected.sides[i]
        actualSide := actual.sides[i]
        if actualSide != expectedSide {
            t.Fatalf("expected \n%v\n but got \n%v\n", expectedSide, actualSide)
        }
    }
}


func Test(t *testing.T) {
    testCmp(t, 1, 1, 1, Cube{
        sides: []Side{
            // top
            Side{
                topLeft: newPoint(0.5, 1.5, 0.5),
                topRight: newPoint(1.5, 1.5, 0.5),
                bottomLeft: newPoint(0.5, 1.5, 1.5),
                bottomRight: newPoint(1.5, 1.5, 1.5),
            },
            // bottom
            Side{
                topLeft: newPoint(0.5, 0.5, 0.5),
                topRight: newPoint(1.5, 0.5, 0.5),
                bottomLeft: newPoint(0.5, 0.5, 1.5),
                bottomRight: newPoint(1.5, 0.5, 1.5),
            },
            // left
            Side{
                topLeft: newPoint(0.5, 1.5, 1.5),
                topRight: newPoint(0.5, 1.5, 0.5),
                bottomLeft: newPoint(0.5, 0.5, 1.5),
                bottomRight: newPoint(0.5, 0.5, 0.5),
            },
            // right
            Side{
                topLeft: newPoint(1.5, 1.5, 1.5),
                topRight: newPoint(1.5, 1.5, 0.5),
                bottomLeft: newPoint(1.5, 0.5, 1.5),
                bottomRight: newPoint(1.5, 0.5, 0.5),
            },
            // front
            Side{
                topLeft: newPoint(0.5, 1.5, 1.5),
                topRight: newPoint(1.5, 1.5, 1.5),
                bottomLeft: newPoint(0.5, 0.5, 1.5),
                bottomRight: newPoint(1.5, 0.5, 1.5),
            },
            // back
            Side{
                topLeft: newPoint(0.5, 1.5, 0.5),
                topRight: newPoint(1.5, 1.5, 0.5),
                bottomLeft: newPoint(0.5, 0.5, 0.5),
                bottomRight: newPoint(1.5, 0.5, 0.5),
            },
        },
    })
}
