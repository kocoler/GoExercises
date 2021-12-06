package main

// 下次或许我们可以用更优的 分治 + 中序遍历 来解决这道问题

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return build(head, nil)
}

func build(node *ListNode, right *ListNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Next == right {
		return &TreeNode{
			Val: node.Val,
		}
	}

	s, f := node, node
	for f != right && f.Next != right {
		s = s.Next
		f = f.Next.Next
	}

	ret := &TreeNode{
		Val:  s.Val,
		Left: build(node, s),
	}
	if s.Next != right {
		ret.Right = build(s.Next, right)
	}

	return ret
}
