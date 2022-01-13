package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	oddList, evenList := &ListNode{}, &ListNode{}
	oddTail, evenTail := oddList, evenList
	flag := true // true if odd node

	h := head
	for h != nil {
		if flag {
			oddTail.Next, oddTail = h, h
		} else {
			evenTail.Next, evenTail = h, h
		}

		flag = !flag
		h = h.Next
	}

	oddTail.Next = evenList.Next
	evenTail.Next = nil

	return oddList.Next
}
