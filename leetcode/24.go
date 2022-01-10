package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	h, pre := head, dummy
	for h != nil && h.Next != nil {
		h.Next, h.Next.Next, pre.Next, pre = h.Next.Next, h, h.Next, h
		h = h.Next
	}

	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	ret := head.Next
	head.Next.Next, head.Next = head, swapPairs(head.Next.Next)

	return ret
}
