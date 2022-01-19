package main

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	h *ListNode
}


func Constructor(head *ListNode) Solution {
	return Solution{h: head}
}


func (this *Solution) GetRandom() int {
	res := this.h.Val
	for h, index := this.h, 1; h != nil; h, index = h.Next, index + 1 {
		if rand.Intn(index) == 0 {
			res = h.Val
		}
	}

	return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
