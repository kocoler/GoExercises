package main

// ！如果正向存放的话，用一个栈来维护！

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位递归
	return recursion(l1, l2, 0)
}

func recursion(l1 *ListNode, l2 *ListNode, add int) *ListNode {
	if l1 == nil && l2 == nil && add == 0 {
		return nil
	}

	var ret *ListNode
	var val1, val2 int
	var l1Next, l2Next *ListNode
	if l1 == nil {
		ret = l2
	} else {
		val1 = l1.Val
		l1Next = l1.Next
	}
	if l2 == nil {
		ret = l1
	} else {
		val2 = l2.Val
		l2Next = l2.Next
	}
	if ret == nil {
		ret = &ListNode{Val: 0, Next: nil}
	}
	ret.Val = 0

	ret.Val = val1 + val2 + add
	carry := ret.Val / 10
	ret.Val %= 10

	ret.Next = recursion(l1Next, l2Next, carry)
	return ret
}
