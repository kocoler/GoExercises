package main

// 回溯

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res [][]int
var ts int
func pathSum(root *TreeNode, targetSum int) [][]int {
	res = [][]int{}
	now = []int{}
	ts = targetSum

	if root == nil {
		return res
	}

	backtrace(root, 0)

	return res
}

var now []int
func backtrace(root *TreeNode, n int) {
	n = n + root.Val
	now = append(now, root.Val)

	// legal
	if root.Left == nil && root.Right == nil {
		if n == ts {
			nres := make([]int, len(now))
			copy(nres, now)
			res = append(res, nres)
		}
		return
	}

	// left / right backtrace
	if root.Left != nil {
		backtrace(root.Left, n)
		now = now[:len(now)-1]
	}
	if root.Right != nil {
		backtrace(root.Right, n)
		now = now[:len(now)-1]
	}
}

