package main

// 同时递归两个指针，代表两部分树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return recursion(root.Left, root.Right)
}

func recursion(l, r *TreeNode) bool {
	if l == nil || r == nil {
		if l == nil && r == nil {
			return true
		}
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return recursion(l.Left, r.Right) && recursion(l.Right, r.Left)
}
