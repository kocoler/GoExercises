package main

type MinIntHeap struct {
	intHeap
}

type MaxIntHeap struct {
	intHeap
}

// An intHeap is a min-heap of ints.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h.intHeap[i] < h.intHeap[j] }
func (h MaxIntHeap) Less(i, j int) bool { return h.intHeap[i] > h.intHeap[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
