// 构建平衡树
// 递归

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var array []int
func sortedArrayToBST(nums []int) *TreeNode {
	array = nums
	return build(0, len(nums) - 1)
}

func build(start, end int) *TreeNode {
	var node TreeNode
	mid := (start + end) / 2
	node.Val = array[mid]

	if start == end {
		return &node
	}
	if start > end {
		return nil
	}

	node.Left = build(start, mid - 1)
	node.Right = build(mid + 1, end)

	return &node
}
