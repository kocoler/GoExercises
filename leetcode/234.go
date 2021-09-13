package main

// 回文链表
// 先找中点
// 再翻转后面
// 再比较
// 缺少复原

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	f := head.Next
	s := head
	for f != nil && f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	f = head
	s = s.Next
	tail := s
	for tail != nil &&tail.Next != nil {
		s, tail.Next, tail.Next.Next = tail.Next, tail.Next.Next, s
	}

	for s != nil {
		if f.Val != s.Val {
			return false
		}

		f = f.Next
		s = s.Next
	}

	return true
}
