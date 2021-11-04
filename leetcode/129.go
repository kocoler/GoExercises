package main

// 感觉今天思路不是很清晰诶
// 主要就是考虑一下怎么不加重的问题

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, 0)
}

func dfs(root *TreeNode, num int) int {
	if root == nil {
		return num
	}

	if root.Left == nil && root.Right == nil {
		return num * 10 + root.Val
	}

	if root.Left == nil {
		return dfs(root.Right, num*10 + root.Val)
	}

	if root.Right == nil {
		return dfs(root.Left, num*10 + root.Val)
	}

	return dfs(root.Left, num*10 + root.Val) + dfs(root.Right, num*10 + root.Val)
}

