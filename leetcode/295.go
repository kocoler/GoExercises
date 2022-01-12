package main

import "container/heap"

type intHeap []int

type minHeap struct {
	intHeap
}

type maxHeap struct {
	intHeap
}

func (m intHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (i intHeap) Len() int {
	return len(i)
}

func (m minHeap) Less(i, j int) bool {
	return m.intHeap[i] < m.intHeap[j]
}

func (m maxHeap) Less(i, j int) bool {
	return m.intHeap[i] > m.intHeap[j]
}

func (i *intHeap) Pop() interface{} {
	x := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return x
}

func (i *intHeap) Push(x interface{}) {
	*i = append(*i, x.(int))
}

type MedianFinder struct {
	minHeap minHeap
	maxHeap maxHeap
}

func Constructor() MedianFinder {
	var m minHeap
	var m2 maxHeap
	heap.Init(&m)
	heap.Init(&m2)
	return MedianFinder{m, m2}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 {
		heap.Push(&this.maxHeap, num)
		return
	}

	if num < this.maxHeap.intHeap[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}

	if this.minHeap.Len() > this.maxHeap.Len()+1 {
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap).(int))
	} else if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(&this.minHeap, heap.Pop(&this.maxHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.maxHeap.intHeap[0]+this.minHeap.intHeap[0]) / 2
	}
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.intHeap[0])
	}
	return float64(this.minHeap.intHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
