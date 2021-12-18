package main

import "math"

// 这其实也是后序遍历的思想吧！！！！

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func maxSumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res = 0

	dfs(root)

	return res
}

// 贡献值，最大值，最小值，是否为BST
// 如果这个节点也是BST，那么它的贡献值相加
func dfs(root *TreeNode) (int, int, int, bool) {
	if root == nil {
		return 0, math.MinInt64, math.MaxInt64, true
	}

	sum := root.Val
	leftSum, leftMax, leftMin, leftOk := dfs(root.Left)
	if root.Left == nil {
		leftMin = root.Val
	}
	rightSum, rightMax, rightMin, rightOk := dfs(root.Right)
	if root.Right == nil {
		rightMax = root.Val
	}
	if leftOk && rightOk && root.Val > leftMax && root.Val < rightMin {
		sum += leftSum + rightSum
		if sum > res {
			res = sum
		}
		return sum, rightMax, leftMin, true
	}

	return 0, 0, 0, false
}
