package main

import (
	"container/heap"
)

// 每一步都要小心处理好就挺麻烦的

type maxHeap []int

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m *maxHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	var hp maxHeap
	heap.Init(&hp)
	var i int

	height := heights[0]
	for i = 1; i < len(heights); i++ {
		if height < heights[i] {
			diff := heights[i] - height
			if bricks < diff {
				if ladders <= 0 {
					break
				} else {
					if len(hp) > 0 && hp[0] >= diff {
						bricks += heap.Pop(&hp).(int)
						bricks -= diff
						heap.Push(&hp, diff)
					}
					ladders--
				}
			} else {
				bricks -= diff
				heap.Push(&hp, diff)
			}
		}
		height = heights[i]
	}

	return i - 1
}
