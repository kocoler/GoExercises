package main

// 前缀和 + 回溯

// 更聪明的办法: 前缀和 + 回溯， 用 HashMap 来维护中间值 => now - target in HashMap

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int
var now []int
var ts int
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	pSum(root)

	res = 0
	now = []int{0}
	ts = targetSum

	backtrace(root)

	return res
}

func backtrace(root *TreeNode) {
	n := root.Val
	// fmt.Println(n)
	lenn := len(now)
	for i := 1; i <= lenn; i ++ {
		if n - now[i-1] == ts {
			res ++
		}
	}
	now = append(now, n)
	// left / right backtrace
	if root.Left != nil {
		backtrace(root.Left)
		now = now[:len(now)-1]
	}
	if root.Right != nil {
		backtrace(root.Right)
		now = now[:len(now)-1]
	}
}

func pSum(root *TreeNode) {
	if root.Right != nil {
		root.Right.Val += root.Val
		pSum(root.Right)
	}
	if root.Left != nil {
		root.Left.Val += root.Val
		pSum(root.Left)
	}
}
