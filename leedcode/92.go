package main


type ListNode struct {
     Val int
     Next *ListNode
 }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	h := &ListNode{Next: head}
	temp := h
	for i := 1; i < left; i ++ {
		temp = temp.Next
	}

	// ori | left | right
	ori := temp
	temp = temp.Next

	var next *ListNode
	var ri *ListNode
	if temp != nil {
		ri = temp.Next
	} else {
		ri = temp
	}

	for i := left; i < right; i ++ {
		if ri == nil {
			break
		}
		next = ri
		ri = next.Next
		next.Next = ori.Next
		ori.Next = next
	}
	if temp != nil {
		temp.Next = ri
	}

	return h.Next
}