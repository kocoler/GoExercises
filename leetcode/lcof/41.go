package lcof

import "container/heap"

type intHeap []int

type MinIntHeap struct {
	intHeap
}

type MaxIntHeap struct {
	intHeap
}

func (h intHeap) Len() int              { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h.intHeap[i] < h.intHeap[j] }
func (h MaxIntHeap) Less(i, j int) bool { return h.intHeap[i] > h.intHeap[j] }
func (h intHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

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

type MedianFinder struct {
	// true 奇数  odd ｜ false 偶数 even
	odd bool
	minIntHeap MinIntHeap
	maxIntHeap MaxIntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var mf MedianFinder
	heap.Init(&mf.minIntHeap)
	heap.Init(&mf.maxIntHeap)

	return mf
}


func (this *MedianFinder) AddNum(num int)  {
	this.odd = !this.odd

	if this.minIntHeap.Len() == 0 {
		heap.Push(&this.maxIntHeap, num)
		return
	}

	if this.minIntHeap.intHeap[0] > num {
		heap.Push(&this.maxIntHeap, num)
		for this.maxIntHeap.Len() - this.minIntHeap.Len() > 1 {
			heap.Push(&this.minIntHeap, heap.Pop(&this.maxIntHeap))
			//this.minIntHeap.Push(this.maxIntHeap.Pop())
		}
	} else {
		heap.Push(&this.minIntHeap, num)
		for this.minIntHeap.Len() - this.maxIntHeap.Len() > 1 {
			//this.maxIntHeap.Push(this.minIntHeap.Pop())
			heap.Push(&this.maxIntHeap, heap.Pop(&this.minIntHeap))
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.odd {
		if this.minIntHeap.Len() > this.maxIntHeap.Len() {
			return float64(this.minIntHeap.intHeap[0])
		}
		return float64(this.maxIntHeap.intHeap[0])
	}

	return (float64(this.maxIntHeap.intHeap[0]) + float64(this.minIntHeap.intHeap[0])) / 2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
