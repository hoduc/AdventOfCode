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


func TestMinPriorityQueue(t *testing.T) {
    h := &PriorityQueue{
        Items: []*Item{
            &Item{
                Value: 2,
                Priority: 2,
                Index: 0,
            },
            &Item{
                Value: 1,
                Priority: 1,
                Index: 1,
            },
            &Item{
                Value: 10,
                Priority: 10,
                Index: 2,
            },
        },
        LessFn: func(pi, pj int) bool {
            return pi < pj
        },
    }

    heap.Init(h)
    expecteds := []int {1, 2, 10}
    for _, expected := range expecteds {
        actual := heap.Pop(h).(*Item).Value
        if actual != expected {
            t.Fatalf(`minPriorityQueue.pop() : expected %v but got %v`, expected, actual)
        }
    }
}

func TestMaxPriorityQueue(t *testing.T) {
    h := &PriorityQueue{
        Items: []*Item{
            &Item{
                Value: 2,
                Priority: 2,
                Index: 0,
            },
            &Item{
                Value: 1,
                Priority: 1,
                Index: 1,
            },
            &Item{
                Value: 10,
                Priority: 10,
                Index: 2,
            },
        },
        LessFn: func(pi, pj int) bool {
            return pi > pj
        },
    }

    heap.Init(h)
    expecteds := []int {10, 2, 1}
    for _, expected := range expecteds {
        actual := heap.Pop(h).(*Item).Value
        if actual != expected {
            t.Fatalf(`minPriorityQueue.pop() : expected %v but got %v`, expected, actual)
        }
    }
}
