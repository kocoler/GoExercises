package main

import (
	"container/heap"
)

type KthLargest struct {
	k int
	nums IntHeap
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	if len(nums) > k {
		nums = nums[len(nums)-k:]
	}
	heapNums := IntHeap(nums)
	heap.Init(&heapNums)
	return KthLargest{k:k,nums:heapNums}
}


func (this *KthLargest) Add(val int) int {
	// heap.Push(&this.nums, val)
	// if this.nums.Len() > this.k {
	//     heap.Pop(&this.nums)
	// }
	// res := heap.Pop(&this.nums)
	// heap.Push(&this.nums, res)
	// return res.(int)
	if this.nums.Len() < this.k {
		heap.Push(&this.nums, val)
		return this.nums[0]
	}
	if val > this.nums[0] {
		this.nums[0] = val
		heap.Fix(&this.nums, 0)
	}
	return this.nums[0]
}