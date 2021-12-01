package main

// 哈希树
// 我自认为对数据量比较大的时候可能比较好用

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var mod int = 1000000007 // 10^9 + 7

func hash(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return (node.Val + hash(node.Left)*31 + hash(node.Right)*131) % mod
}

func treeHash(node *TreeNode, target int) int {
	if node == nil || res {
		return 0
	}

	p := (node.Val + treeHash(node.Left, target)*31 + treeHash(node.Right, target)*131) % mod
	if p == target {
		res = true
	}
	return p
}

var res bool

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	res = false
	target := hash(t2)
	treeHash(t1, target)

	return res
}
