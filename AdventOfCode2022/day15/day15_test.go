package main

import (
    "testing"
)


func testMergeIntervalsWrapper(t *testing.T, a Interval, b Interval, expected Interval) {
    t.Log("---\n")
    actual := mergeIntervals(a, b)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
    t.Log("---\n")
}

func testMergeIntervals(t *testing.T, as, ae, bs, be, es, ee int) {
    testMergeIntervalsWrapper(t, Interval{start: as, end: ae}, Interval{start: bs, end: be}, Interval{start: es, end: ee})
}

func Test(t *testing.T) {
    testMergeIntervals(t, 0, 0, 0, 0, 0, 0)
    testMergeIntervals(t, 0, 5, 0, 2, 0, 5)
    testMergeIntervals(t, 0, 2, 0, 5, 0, 5)
    testMergeIntervals(t, 0, 5, 1, 2, 0, 5)
    testMergeIntervals(t, 1, 2, 0, 5, 0, 5)
    testMergeIntervals(t, 0, 5, -1, 0, -1, 5)
    testMergeIntervals(t, -1, 0, 0, 5, -1, 5)
    testMergeIntervals(t, 0, 5, 5, 9, 0, 9)
    testMergeIntervals(t, 5, 9, 0, 5, 0, 9)
    testMergeIntervals(t, 0, 5, 2, 3, 0, 5)
}
