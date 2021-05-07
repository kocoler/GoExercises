package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{0, head}
	pre := h
	t := h.Next

	for t != nil && t.Next != nil {
		val := t.Val
		flag := false
		for t.Next != nil && t.Next.Val == val {
			flag = true
			t = t.Next
		}

		if flag {
			pre.Next = t.Next
		} else {
			pre = pre.Next
		}

		t = t.Next
	}

	return h.Next
}
