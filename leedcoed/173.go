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
//
//type BSTIterator struct {
//	root *TreeNode
//	path []*TreeNode
//	lens int
//	dir bool // false: left true: right
//	now  *TreeNode
//}
//
//// 中序遍历
//func initTree(root *TreeNode, res []int) []int {
//	for root.Left != nil {
//		initTree(root, res)
//	}
//	res = append(res, root.Val)
//	for root.Right != nil {
//		initTree(root, res)
//	}
//	return res
//}
//
//func Constructor(root *TreeNode) BSTIterator {
//	temp := root
//	stack := []*TreeNode{}
//	stack = append(stack, root)
//	for temp.Left != nil {
//		stack = append(stack, temp)
//		temp = temp.Left
//	}
//	return BSTIterator{
//		path: stack, // stack
//		lens: len(stack),
//		dir: false,
//		now:  temp,
//	}
//}
//
//func (this *BSTIterator) Next() int {
//	if this.dir {
//
//	} else {
//		temp := this.path[this.lens-1]
//		if this.now.Left
//		if this.lens == 0 {
//			this.dir = true
//		}
//	}
//
//}
//
//func (this *BSTIterator) HasNext() bool {
//
//}


type BSTIterator struct {
	nodes []int
	lenn  int
}

func initTree(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	initTree(root.Left, res)
	res = append(res, root.Val)
	initTree(root.Right, res)

	return res
}


func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		nodes: initTree(root, []int{}),
	}
	iterator.lenn = len(iterator.nodes)
	return iterator
}


func (this *BSTIterator) Next() int {
	res := this.nodes[this.lenn - 1]
	this.nodes = this.nodes[:this.lenn-1]
	this.lenn --
	return res
}


func (this *BSTIterator) HasNext() bool {
	return this.lenn == 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
