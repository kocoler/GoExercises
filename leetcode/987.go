package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type node struct {
	r   int
	c   int
	val int
}

var queue []node

func verticalTraversal(root *TreeNode) [][]int {
	dfs(root, 0, 0)

	sort.Slice(queue, func(i, j int) bool {
		if queue[i].c == queue[j].c {
			if queue[i].r == queue[j].r {
				return queue[i].val < queue[j].val
			}
			return queue[i].r < queue[j].r
		}

		return queue[i].c < queue[j].c
	})

	fmt.Println(queue)

	var res [][]int
	lenq := len(queue)
	for i := 0; i < lenq; i++ {
		p := queue[i].c
		temp := []int{}
		for ; i < lenq && queue[i].c == p; i++ {
			temp = append(temp, queue[i].val)
		}
		i--
		res = append(res, temp)
	}

	return res
}

func dfs(root *TreeNode, r, c int) {
	if root == nil {
		return
	}

	queue = append(queue, node{r, c, root.Val})

	dfs(root.Left, r+1, c-1)
	dfs(root.Right, r+1, c+1)
}
