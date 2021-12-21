package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	f := head
	s := head
	for f != nil {
		if f.Next == nil {
			return nil
		}
		f = f.Next.Next
		s = s.Next
		if f == s {
			ptr := head
			for ptr != s {
				ptr = ptr.Next
				s = s.Next
			}
			return ptr
		}
	}

	return nil
}
