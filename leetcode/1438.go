package main

import "container/heap"

// 单调队列

type monotoneQueue struct {
	value []int
	len int
}

func (m monotoneQueue)pop() int {
	r := m.value[0]
	m.value = m.value[1:]

	return r
}

func (m monotoneQueue)push(x int) {
	m.value = append(m.value, x)
}

func (m monotoneQueue)length() int {
	return m.len
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func longestSubarray(nums []int, limit int) int {
	res := 0
	queue := monotoneQueue{}
	h := &intHeap{}
	heap.Init(h)

	lenn := len(nums)
	sum := 0

	for i := 0; i < lenn; i ++ {

	}

	return res
}

func abs(a, b int) int {
	c := a - b

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
