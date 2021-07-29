package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// var min [2]int
var min int
var res int
func findSecondMinimumValue(root *TreeNode) int {
	min = root.Val
	res = -1

	dfs(root)

	return res
}

func dfs(root *TreeNode) {
	if root.Left == nil {
		if root.Val > min {
			if res == -1 {
				res = root.Val
			} else {
				res = minInt(res, root.Val)
			}
		}
	} else {
		dfs(root.Left)
		dfs(root.Right)
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
