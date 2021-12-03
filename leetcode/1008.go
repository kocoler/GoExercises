package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1)
}

func build(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	node := TreeNode{Val: preorder[left]}
	i := sort.Search(right-left+1, func(i int) bool {
		return preorder[left+i] > preorder[left]
	})

	node.Left = build(preorder, left+1, left+i-1)
	node.Right = build(preorder, left+i, right)

	return &node
}
