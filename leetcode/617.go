package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	value := 0
	var r1l, r2l, r1r, r2r *TreeNode

	if root1 != nil {
		value += root1.Val
		r1l = root1.Left
		r1r = root1.Right
	}
	if root2 != nil {
		value += root2.Val
		r2l = root2.Left
		r2r = root2.Right
	}
	node := TreeNode{Val: value}
	node.Left = mergeTrees(r1l, r2l)
	node.Right = mergeTrees(r1r, r2r)

	return &node
}
