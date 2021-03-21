package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var h *ListNode

func reverseList(head *ListNode) *ListNode {
	return reverseList2(head)
}

func reverseList2(head *ListNode) *ListNode {
	cur := head
	new := &ListNode{}
	for cur != nil {
		next := cur.Next
		cur.Next = new.Next
		new.Next = cur
		cur = next
	}

	return new.Next
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	f := rec(head)
	f.Next = nil

	return h
}

func rec(head *ListNode) *ListNode {
	if head.Next == nil {
		h = head
		return head
	}
	p := rec(head.Next)
	p.Next = head
	return head
}

