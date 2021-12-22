package main

// 一遍过，好耶
// 这个是 自顶向下

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

// 1 -> 2 -> 3
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, s, f := &ListNode{Next: head}, head, head
	for f != nil && f.Next != nil {
		pre, s, f = pre.Next, s.Next, f.Next.Next
	}
	pre.Next = nil

	return merge(mergeSort(head), mergeSort(s))
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
