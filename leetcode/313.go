package main

import "container/heap"

type MinIntHeap struct {
	intHeap
}

// An intHeap is a min-heap of ints.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h.intHeap[i] < h.intHeap[j] }
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

func nthSuperUglyNumber(n int, primes []int) int {
	var minHeap MinIntHeap
	minHeap.intHeap = make([]int, len(primes))
	copy(minHeap.intHeap, primes)
	heap.Init(&minHeap)
	n --
	res := 1
	for n > 0 {
		num := heap.Pop(&minHeap).(int)
		for _, v := range primes {
			heap.Push(&minHeap, num * v)
		}
		res = num
		n --
	}

	return res
}
