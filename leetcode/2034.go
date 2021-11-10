package main

import "container/heap"

type stockprice struct {
	timestamp int
	parice    int
}

type stockprices []stockprice

func (s stockprices) Len() int { return len(s) }

func (s stockprices) Less(i, j int) bool {
	return s[i].parice < s[j].parice
}

func (s stockprices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type minStockHeap struct {
	stockprices
}

func (h *minStockHeap) Push(x interface{}) {
	(*h).stockprices = append((*h).stockprices, x.(stockprice))
}

func (h *minStockHeap) Pop() interface{} {
	old := (*h).stockprices
	n := len(old)
	x := old[n-1]
	(*h).stockprices = old[0 : n-1]
	return x
}

type maxStockHeap struct {
	stockprices
}

func (h maxStockHeap) Less(i, j int) bool {
	return h.stockprices[i].parice > h.stockprices[j].parice
}

func (h *maxStockHeap) Push(x interface{}) {
	(*h).stockprices = append((*h).stockprices, x.(stockprice))
}

func (h *maxStockHeap) Pop() interface{} {
	old := (*h).stockprices
	n := len(old)
	x := old[n-1]
	(*h).stockprices = old[0 : n-1]
	return x
}

type StockPrice struct {
	lastTimestamp int
	lastPrice     int

	// 可以用 -price 和 price 来表示
	// 这样就不用写两个堆了
	minQueue minStockHeap
	maxQueue maxStockHeap

	m map[int]int
}

func Constructor() StockPrice {
	return StockPrice{
		lastTimestamp: 0,
		lastPrice:     0,
		minQueue:      minStockHeap{},
		maxQueue:      maxStockHeap{},
		m:             make(map[int]int),
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if timestamp >= this.lastTimestamp {
		this.lastTimestamp = timestamp
		this.lastPrice = price
	}

	this.m[timestamp] = price
	heap.Push(&this.minQueue, stockprice{timestamp, price})
	heap.Push(&this.maxQueue, stockprice{timestamp, price})
}

func (this *StockPrice) Current() int {
	return this.lastPrice
}

func (this *StockPrice) Maximum() int {
	top := this.maxQueue.stockprices[0]
	for top.parice != this.m[top.timestamp] {
		heap.Pop(&this.maxQueue)
		top = this.maxQueue.stockprices[0]
	}

	return top.parice
}

func (this *StockPrice) Minimum() int {
	top := this.minQueue.stockprices[0]
	for top.parice != this.m[top.timestamp] {
		heap.Pop(&this.minQueue)
		top = this.minQueue.stockprices[0]
	}

	return top.parice
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
