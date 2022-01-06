package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lenq := len(queue)

		res = append(res, queue[lenq-1].Val)

		for i := 0; i < lenq; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[lenq:]
	}

	return res
}
