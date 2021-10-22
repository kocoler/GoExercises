package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 1:12
// 1:23
//var res int
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0

	var calNum func(*TreeNode) int
	calNum = func(root *TreeNode) int {
		left, right := 0, 0
		if root.Left != nil {
			left = calNum(root.Left)
		}
		if root.Right != nil {
			right = calNum(root.Right)
		}

		res += abs(left - right)
		return root.Val + left + right
	}

	calNum(root)

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
