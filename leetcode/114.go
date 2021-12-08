package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 迭代
	var node *TreeNode
	stack := []*TreeNode{root}

	var temp *TreeNode
	for len(stack) != 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if temp == nil {
			temp = node
		} else {
			temp.Left, temp.Right, temp = nil, node, node
		}
	}

	// return root
}
