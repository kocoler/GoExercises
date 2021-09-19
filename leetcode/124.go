package main

// 计算当前节点最大贡献值！

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
func maxPathSum(root *TreeNode) int {
    res = -0x3f3f3f

    recursion(root)

    return res
}

func recursion(root *TreeNode) int {
    if root == nil {
        return 0
    }

    now := root.Val

    l := max(recursion(root.Left), 0)
    r := max(recursion(root.Right), 0)

    now = now + l + r

    res = max(res, now)

    // fmt.Println(root.Val, now + max(l, r))

    return root.Val + max(l, r)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

