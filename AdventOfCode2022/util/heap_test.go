package util

import (
    "testing"
    "container/heap"
)

func TestMaxIntHeap(t *testing.T) {
    h := &MaxIntHeap{[]int{2, 1, 10}}
    heap.Init(h)
    expecteds := []int {10,2,1}
    for _, expected := range expecteds {
        actual := heap.Pop(h)
        if actual != expected {
            t.Fatalf(`maxHeap.pop() : expected %v but got %v`, expected, actual)
        }
    }
}
