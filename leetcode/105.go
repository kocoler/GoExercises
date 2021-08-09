package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var p []int
var i []int
func buildTree(preorder []int, inorder []int) *TreeNode {
	p = preorder
	i = inorder

	return build(0, 0, len(i) - 1)
}

func build(now, from, end int) *TreeNode {
	if end < from || now < 0 {
		return nil
	}

	nowValue := p[now]
	node := &TreeNode{Val: nowValue}

	index := 0
	for j := from; j <= end; j ++ {
		if i[j] == nowValue {
			index = j
			break
		}
	}

	leftLen := index - from

	node.Left = build(now+1, from, index - 1)
	node.Right = build(now + 1 + leftLen, index + 1, end)

	return node
}
