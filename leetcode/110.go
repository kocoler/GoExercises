package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 考虑如何计算高度
// 无法避免重复的计算
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(calDeep(root.Left, 0) - calDeep(root.Right, 0))) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func calDeep(node *TreeNode, deep int) int {
	if node == nil {
		return deep
	}
	return max(calDeep(node.Left, deep + 1), calDeep(node.Right, deep+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


