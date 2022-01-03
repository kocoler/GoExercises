package main

import (
	"container/heap"
)

type info struct {
	num int
	exp int
}

type maxHeap []info

func (hp maxHeap) Less(i, j int) bool {
	return hp[i].exp < hp[j].exp
}

func (hp maxHeap) Len() int {
	return len(hp)
}

func (hp maxHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *maxHeap) Pop() interface{} {
	x := (*hp)[len(*hp)-1]
	*hp = (*hp)[:len(*hp)-1]

	return x
}

func (hp *maxHeap) Push(x interface{}) {
	*hp = append(*hp, x.(info))
}

func eatenApples(apples []int, days []int) int {
	if len(apples) == 0 {
		return 0
	}
	var res int

	var hp maxHeap
	heap.Init(&hp)
	day := 1
	//heap.Push(&hp, info{apples[0], days[0]})
	for hp.Len() > 0 || len(apples) >= day {
		// push today apple
		if day <= len(apples) && apples[day-1] != 0 {
			heap.Push(&hp, info{apples[day-1], day + days[day-1]})
		}

		// clean
		for hp.Len() > 0 && (hp[0].exp <= day || hp[0].num <= 0) {
			heap.Pop(&hp)
		}

		// eat
		if hp.Len() > 0 {
			hp[0].num--
			res++
		}

		day++
	}

	return res
}
