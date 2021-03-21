package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var l int
var h int
var res int

func rangeSumBST(root *TreeNode, low int, high int) int {
	l = low
	h = high
	res = 0
	dfs(root)

	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Val < l {
		dfs(root.Right)
		return
	}

	if root.Val > h {
		dfs(root.Left)
		return
	}

	res += root.Val
	dfs(root.Left)
	dfs(root.Right)
}
