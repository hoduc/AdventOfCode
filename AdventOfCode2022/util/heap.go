package util

import(
    "container/heap"
)

// https://pkg.go.dev/container/heap
type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


/*
   Why can I not able to do this ?
   type MaxIntHeap = IntHeap

   func (h MaxIntHeap) Less(i, j int) bool {
    return h[i] > h[j]
   }
*/

type MaxIntHeap struct {
    IntHeap
}

func (h MaxIntHeap) Less(i, j int) bool {
    return h.IntHeap[i] > h.IntHeap[j]
}

type Item struct {
    Value any
    Priority int
    Index int
}

type PriorityQueue struct {
    Items []*Item
    LessFn func(pi, pj int) bool
    Value any
}

func (pq PriorityQueue) Len() int {
    return len(pq.Items)
}


func (pq PriorityQueue) Less(i, j int) bool {
    return pq.LessFn(pq.Items[i].Priority, pq.Items[j].Priority)
}


func (pq PriorityQueue) Swap(i, j int) {
	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
	pq.Items[i].Index = i
	pq.Items[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(pq.Items)
	item := x.(*Item)
	item.Index = n
	pq.Items = append(pq.Items, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	pq.Items = old.Items[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value any, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
