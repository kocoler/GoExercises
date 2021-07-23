package main

// 双指针

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// s := head
	f := head
	n := 0

	for f != nil {
		n ++
		if n % 2 == 0 {
			head = head.Next
		}
		f = f.Next
	}

	return head
}
