package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代的中序遍历
// 要杀 pre 节点嘞

type BSTIterator struct {
	node  *TreeNode
	queue []*TreeNode
	//pre   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var queue []*TreeNode
	node := root
	for node != nil {
		queue = append(queue, node)
		node = node.Left
	}

	return BSTIterator{node: node, queue: queue}
}

func (this *BSTIterator) Next() int {
	ret := math.MinInt64

	node := this.node
	//for {
	for node != nil {
		this.queue = append(this.queue, node)
		node = node.Left
	}
	node = this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	//if node.Right == nil || this.pre == node.Right {
	//	this.pre = node
	//	this.node = nil
	//} else {
	ret = node.Val
	this.node = node.Right
	//	break
	//}
	//}

	return ret
}

func (this *BSTIterator) HasNext() bool {
	if this.node == nil && len(this.queue) == 0 {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
