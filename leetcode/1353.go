package main

import (
	"container/heap"
	"sort"
)

type item struct {
	start int
	end   int
}

type items []item

type minHeap []int

func (i items) Len() int {
	return len(i)
}

func (i items) Less(j int, k int) bool {
	return i[j].start < i[k].start
}

func (i items) Swap(j int, k int) {
	i[j], i[k] = i[k], i[j]
}

func (i minHeap) Len() int {
	return len(i)
}

func (i minHeap) Less(a, b int) bool {
	return i[a] < i[b]
}

func (i minHeap) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i *minHeap) Push(x interface{}) {
	*i = append(*i, x.(int))
}

func (i *minHeap) Pop() interface{} {
	old := *i
	n := len(old)
	x := old[n-1]
	*i = old[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {
	var eventsItems items
	lene := len(events)
	eventsItems = make(items, lene)
	m := 0
	for i := 0; i < lene; i++ {
		eventsItems[i] = item{
			start: events[i][0],
			end:   events[i][1],
		}
		if m < events[i][1] {
			m = events[i][1]
		}
	}
	sort.Sort(eventsItems)

	var hp minHeap
	var num int
	heap.Init(&hp)

	l := 0
	for i := 1; i <= m; i++ {
		for ; l < lene && eventsItems[l].start <= i; l++ {
			heap.Push(&hp, eventsItems[l].end)
		}
		for len(hp) > 0 && hp[0] < i {
			heap.Pop(&hp)
		}
		if len(hp) > 0 {
			heap.Pop(&hp)
			num++
		}
	}

	return num
}
