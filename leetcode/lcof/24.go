package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 3: 37 - 3: 41
// 5 min
func reverseList(head *ListNode) *ListNode {
	// 1 2 3 4 5 nil
	// 2 1 3 4 5 nil
	if head == nil {
		return nil
	}

	h := head
	next := head.Next
	tail := head
	for next != nil {
		h.Next, next.Next, next, tail = next.Next, tail, next.Next, next
	}

	return tail
}
