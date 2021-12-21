package main

// 迭代

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var queue []*TreeNode
	var res []int
	node := root
	var prev *TreeNode

	for node != nil || len(queue) > 0 {
		for node != nil {
			queue = append(queue, node)
			node = node.Left
		}
		node = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node.Right == nil || node.Right == prev {
			res = append(res, node.Val)
			prev = node
			node = nil
		} else {
			queue = append(queue, node)
			node = node.Right
		}
	}

	return res
}
