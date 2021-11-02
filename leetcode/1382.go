package main

// 比起节点旋转
// 还是建表再建树
// 所以他才是 medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	var nums []int
	var mf func(*TreeNode)
	mf = func(root *TreeNode) {
		if root != nil {
			mf(root.Left)
			nums = append(nums, root.Val)
			mf(root.Right)
		}
	}
	mf(root)

	var build func(int, int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		var node TreeNode
		mid := (l + r) / 2
		node.Val = nums[mid]
		node.Left = build(l, mid-1)
		node.Right = build(mid+1, r)
		return &node
	}
	return build(0, len(nums)-1)
}
