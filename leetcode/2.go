package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	h := dummy
	for l1 != nil && l2 != nil {
		now := &ListNode{}
		now.Val = l1.Val + l2.Val + carry
		carry = now.Val / 10
		now.Val = now.Val % 10
		h.Next = now
		h = h.Next
		l1, l2 = l1.Next, l2.Next
	}

	temp := l1
	if l1 == nil {
		temp = l2
	}

	for temp != nil {
		now := &ListNode{}
		now.Val = temp.Val + carry
		carry = now.Val / 10
		now.Val = now.Val % 10
		h.Next = now
		h = h.Next
		temp = temp.Next
	}

	if carry != 0 {
		now := &ListNode{Val: carry}
		h.Next = now
	}

	return dummy.Next
}
