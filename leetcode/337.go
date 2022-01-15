package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dp1, dp2 := dfs(root)

	return max(dp1, dp2)
}

// return dp1, dp2
func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ldp1, ldp2 := dfs(root.Left)
	rdp1, rdp2 := dfs(root.Right)

	//dp1, dp2 := ldp1+rdp1, ldp2+rdp2

	// 这个为什么可以也很值得思考
	//if dp2+root.Val > dp1 {
	//	return dp2 + root.Val, dp1
	//}

	return ldp2 + rdp2 + root.Val, max(ldp1, ldp2) + max(rdp1, rdp2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
