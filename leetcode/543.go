package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
func diameterOfBinaryTree(root *TreeNode) int {
	res = 0

	reverse(root)

	return res
}

func reverse(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	r := 0
	l := 0
	if root.Left != nil {
		l = reverse(root.Left) + 1
	}
	if root.Right != nil {
		r = reverse(root.Right) + 1
	}

	if res < l + r {
		res = l + r
	}

	if l > r {
		return l
	}

	return r
}
