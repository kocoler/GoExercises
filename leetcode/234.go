package main

// 回文链表
// 先找中点
// 再翻转后面
// 再比较
// 缺少复原

// 递归解法：
// 外部全局变量
// 里面递归，递归一定是从尾到头，全局变量一定是从头到尾
// 这个思路挺神奇的
// 当然最优解还是双指针

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
	for tail != nil && tail.Next != nil {
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
