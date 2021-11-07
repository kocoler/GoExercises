package main

import (
	"container/heap"
	"fmt"
)

// 优先队列（大顶堆）
// 需要注意的是 heap 的判断条件
// 这里的判断条件是 下一个操作 和 当前状态之间 的判断
// 也就是计算如果进行下一个操作的收益
// 提前计算下一个操作

type RateNode struct {
	Val float64
	Pass float64
	Total float64
}

type rateHeap []*RateNode

type MaxRateHeap struct {
	rateHeap
}

func (h rateHeap) Len() int              { return len(h) }
func (h MaxRateHeap) Less(i, j int) bool { return h.rateHeap[i].Val > h.rateHeap[j].Val }
func (h rateHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

func (h *rateHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*RateNode))
}

func (h *rateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(maxAverageRatio([][]int{{1,2},{3,5},{2,2}}, 2))
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var rateheap MaxRateHeap
	heap.Init(&rateheap)
	for k := range classes {
		p, t := float64(classes[k][0]), float64(classes[k][1])
		heap.Push(&rateheap, &RateNode{
			Val:  ((p + 1) / (t + 1)) - (p / t),
			Pass:  p,
			Total: t,
		})
	}

	for extraStudents != 0 {
		rateheap.rateHeap[0].Pass ++
		rateheap.rateHeap[0].Total ++
		rateheap.rateHeap[0].Val = ((rateheap.rateHeap[0].Pass + 1) / (rateheap.rateHeap[0].Total + 1)) - rateheap.rateHeap[0].Pass / rateheap.rateHeap[0].Total
		heap.Fix(&rateheap, 0)
		extraStudents --
	}

	var res float64
	for k := range rateheap.rateHeap {
		fmt.Println(rateheap.rateHeap[k].Val, rateheap.rateHeap[k].Pass, rateheap.rateHeap[k].Total)
		res += rateheap.rateHeap[k].Pass / rateheap.rateHeap[k].Total
	}

	return res / float64(len(rateheap.rateHeap))
}
