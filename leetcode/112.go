package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ts int
func hasPathSum(root *TreeNode, targetSum int) bool {
	ts = targetSum

	return dfs(root, 0)
}

func dfs(root *TreeNode, n int) bool {
	if root == nil {
		return false
	}
	n = n + root.Val
	if root.Left == nil && root.Right == nil {
		if n == ts {
			return true
		}
	}

	return dfs(root.Left, n) || dfs(root.Right, n)
}

