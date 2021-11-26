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
var resDeep int
// var resL int
func findBottomLeftValue(root *TreeNode) int {
	res = root.Val
	resDeep = 0

	dfs(root, 1)

	return res
}

func dfs(node *TreeNode, deep int) {
	if resDeep < deep {
		res = node.Val
		resDeep = deep
	}

	if node.Left != nil {
		dfs(node.Left, deep + 1)
	}
	if node.Right != nil {
		dfs(node.Right, deep + 1)
	}
}
