package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	h := &ListNode{-0x3f3f3f, head}
	t := head

	tail := h.Next
	for t != nil {
		pre := h
		i := h.Next

		for i != nil && i.Val < t.Val {
			pre = pre.Next
			i = i.Next
		}

		if i == nil || (i == t) {
			tail = t
		}

		pre.Next, t.Next, t = t, i, t.Next

		if tail.Next != nil {
			tail.Next = nil
		}
	}
	tail.Next = nil

	return h.Next
}
