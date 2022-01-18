package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kSmallestPairs([]int{1,2}, []int{3}, 3))
}

type info struct {
	num int
	x, y int
}

type minHeap []info

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j],m[i]
}

func (m minHeap) Less(i, j int) bool {
	return m[i].num < m[j].num
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(info))
}

func (m *minHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := make([][]int, 0, k)

	var hp minHeap
	heap.Init(&hp)
	heap.Push(&hp, info{num: nums1[0]+nums2[0], x: 0, y: 0})
	mi := 0 // max index -> max index > nx -> mi ++
	for k > 0 && len(hp) > 0 {
		min := heap.Pop(&hp).(info)
		fmt.Println(min)
		if min.y < len(nums2)-1 {
			heap.Push(&hp, info{nums1[min.x]+nums2[min.y+1], min.x, min.y+1})
		}
		if mi < len(nums1)-1 && mi == min.x {
			mi ++
			heap.Push(&hp, info{nums1[mi]+nums2[0], mi, 0})
		}

		k --
		res = append(res, []int{min.x, min.y})
	}

	return res
}
