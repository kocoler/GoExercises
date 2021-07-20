package main

// 基本的相交链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA := headA
	hB := headB
	ALen := 0
	BLen := 0
	for hA != nil {
		hA = hA.Next
		ALen ++
	}
	for hB != nil {
		hB = hB.Next
		BLen ++
	}
	for ALen > BLen {
		headA = headA.Next
		ALen --
	}
	for BLen > ALen {
		headB = headB.Next
		BLen --
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

