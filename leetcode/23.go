package main

import "container/heap"

type ListNode struct {
	Val int
	Next *ListNode
}

type intHeap []*ListNode

type MinIntHeap struct {
	intHeap
}

func (h intHeap) Len() int              { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h.intHeap[i].Val < h.intHeap[j].Val }
func (h intHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap MinIntHeap
	heap.Init(&minHeap)

	for _, v := range lists {
		heap.Push(&minHeap, v)
	}

	head := &ListNode{}
	p := head

	for minHeap.Len() > 0 {
		node := heap.Pop(&minHeap).(*ListNode)
		// fmt.Println(node.Val)
		p.Next = node
		p = p.Next
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
	}

	return head.Next
}
