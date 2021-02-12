package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	stackT := []*TreeNode{root}
	count := 0
	for {
		for root != nil {
			stackT = append(stackT, root)
			root = root.Left
		}
		temp := stackT[len(stackT)-1]
		count ++
		if count == k {
			return temp.Val
		}
		stackT = stackT[:len(stackT)-1]
		root = temp.Right
	}

	return 0
}
