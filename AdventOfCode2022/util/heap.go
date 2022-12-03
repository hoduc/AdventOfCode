package util


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
