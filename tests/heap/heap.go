// https://golang.google.cn/pkg/container/heap/
// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// An IntHeap is a min-heap of ints.
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

type StringHeap []string

func (s *StringHeap) Len() int {
	return len(*s)
}

func (s *StringHeap) Less(i, j int) bool {
	return strings.Compare((*s)[i], (*s)[j]) > 0
}

func (s *StringHeap) Swap(i, j int) {
	var temp string
	temp = (*s)[i]
	(*s)[i] = (*s)[j]
	(*s)[j] = temp

	return
}

func (s *StringHeap) Push(x interface{}) {
	*s = append(*s, x.(string))

	return
}

func (s *StringHeap) Pop() interface{} {
	//old := *s
	//n := len(old)
	//x := old[n-1]
	//*s = old[0 : n-1]
	x := (*s)[len(*s) - 1]
	*s = (*s)[0:len(*s) - 1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	testInt()
	testString()
}

func testString() {
	s := "h"
	h := &StringHeap{"b", "d", "c", "z", "y", "g", s}
	heap.Init(h)
	heap.Push(h, "a")

	fmt.Printf("minimum: %s\n", (*h)[0])
	(*h)[0] = "x"
	//heap.Fix(h, 1)
	for h.Len() > 0 {
		fmt.Printf("%s ", heap.Pop(h))
	}
}

func testInt() {
	a := []int{1, 2, 3}
	var b *IntHeap
	b = (*IntHeap)(&a)
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Init(b)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
