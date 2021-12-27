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
func isEvenOddTree(root *TreeNode) bool {
	// 层序遍历
	queue := []*TreeNode{root}

	odd := true
	lastValue := math.MinInt64
	for len(queue) > 0 {
		lenq := len(queue)

		for i := 0; i < lenq; i++ {
			node := queue[i]

			if odd {
				if node.Val%2 == 0 || node.Val <= lastValue {
					return false
				}
			} else {
				if node.Val%2 != 0 || node.Val >= lastValue {
					return false
				}
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Left != nil {
				queue = append(queue, node.Right)
			}
		}

		odd = !odd
		if odd {
			lastValue = math.MinInt64
		} else {
			lastValue = math.MaxInt64
		}

		queue = queue[lenq:]
	}

	return true
}
