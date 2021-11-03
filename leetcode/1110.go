package main

// 通过看别人的提交记录
// **TreeNode 可以从后面的状态来删除前面的状态！

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var m map[int]bool
var res []*TreeNode
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	m = make(map[int]bool)
	res = []*TreeNode{}

	for _, v := range to_delete {
		m[v] = true
	}

	if m[root.Val] {
		//if root.Left != nil {
		//	dfs(root.Left, true)
		//}
		//if root.Right != nil {
		//	dfs(root.Right, true)
		//}
		dfs(root, true)
	} else {
		res = append(res, root)
		dfs(root, false)
	}

	return res
}

func dfs(root *TreeNode, delete bool) {
	if root.Left != nil {
		if m[root.Left.Val] {
			dfs(root.Left, true)
			root.Left = nil
		} else {
			if delete {
				res = append(res, root.Left)
			}
			dfs(root.Left, false)
		}
	}
	if root.Right != nil {
		if m[root.Right.Val] {
			dfs(root.Right, true)
			root.Right = nil
		} else {
			if delete {
				res = append(res, root.Right)
			}
			dfs(root.Right, false)
		}
	}
}
