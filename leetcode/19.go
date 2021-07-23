package main

// 删除连标的倒数第N个节点：双指针
// 挺神奇的

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	len := 0
	for fast.Next != nil {
		len ++
		// ** 这里，长度大于倒数的情况下，慢指针可以开始走 **
		if len > n {
			slow = slow.Next
		}
		fast = fast.Next
	}
	if slow == head {
		if len + 1 == n {
			return head.Next
		}
	}
	slow.Next = slow.Next.Next
	return head
}
