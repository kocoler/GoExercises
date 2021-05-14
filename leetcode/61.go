package main

// 注意的就是 回环 优化
// **还可以直接反转尾部**
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	res := ListNode{Next: head}
	h := res.Next
	num := 1
	for h.Next != nil {
		h = h.Next
		num ++
	}
	n := k%num
	tail := h
	for i := 0; i < n; i ++ {
		tail.Next = res.Next
		res.Next = tail
		newTail := res.Next
		for newTail.Next != tail {
			newTail = newTail.Next
		}
		tail = newTail
	}
	tail.Next = nil

	return res.Next
}
