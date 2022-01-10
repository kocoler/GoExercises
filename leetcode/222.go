package main

import (
	"fmt"
	"math"
)

// 这道题很有意思
// 完全树上的二分

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var mask int
var h int

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	node := root
	h = 1
	for node.Left != nil {
		h++
		node = node.Left
	}

	l := int(math.Pow(2, float64(h-1)))
	r := int(math.Pow(2, float64(h))) - 1
	mask = int(math.Pow(2, float64(h-1)))
	for l < r {
		mid := l + (r-l)/2
		if findMid(root, mid) == nil {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if findMid(root, l) == nil {
		return l - 1
	}
	return l
}

func findMid(root *TreeNode, mid int) *TreeNode {
	node := root
	f := 1
	nowMask := mask >> 1
	for f < h && node != nil {
		f++
		if mid&nowMask == 0 {
			node = node.Left
		} else {
			node = node.Right
		}
		nowMask >>= 1
	}

	return node
}

func findMid(root *TreeNode, mid int) *TreeNode {
	node := root
	for mid > 0 {
		if mid&1 == 0 {
			// <<
			mid >>= 1
			node = node.Left
		} else {
			mid <<= 1
			node = node.Right
		}
	}

	if node == nil {
		fmt.Println("?")
	} else {
		fmt.Println("!", node.Val)
	}

	return node
}
