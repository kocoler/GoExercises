package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	return reverse(root)
}

func reverse(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = reverse(root.Left)
	}
	if root.Right != nil {
		root.Right = reverse(root.Right)
	}
	root.Left, root.Right = root.Right, root.Left

	return root
}
