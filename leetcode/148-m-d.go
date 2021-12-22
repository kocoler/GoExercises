package main

// 这个是自底向上

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(A, B *ListNode) *ListNode {
	t := &ListNode{}

	nx := t
	for A != nil && B != nil {
		if A.Val > B.Val {
			nx.Next, B = B, B.Next
			nx = nx.Next
		} else {
			nx.Next, A = A, A.Next
			nx = nx.Next
		}
	}
	for A != nil {
		nx.Next, A = A, A.Next
		nx = nx.Next
	}
	for B != nil {
		nx.Next, B = B, B.Next
		nx = nx.Next
	}

	return t.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for h := head; h != nil; h = h.Next {
		length++
	}

	res := &ListNode{Next: head}
	for subLen := 1; subLen < length; subLen <<= 1 {
		prev, cur := res, res.Next
		for cur != nil {
			head1 := cur

		}
	}

	return
}
