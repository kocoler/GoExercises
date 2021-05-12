package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var res [][]int
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		lenq := len(queue)
		temp := make([]int, lenq)
		for i := 0; i < lenq; i++ {
			temp[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, temp)
		queue = queue[lenq:]
	}

	return res
}
