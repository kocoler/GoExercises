// 还原二叉树
// 递归求解
// go 直接切片，很方便


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    lenp := len(preorder)
    leni := len(inorder)
    if lenp == 0 || leni == 0 {
        return nil
    }
    n := TreeNode{Val: preorder[0]}
    i := 0
    for ; i < leni; i ++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    n.Left = buildTree(preorder[1:], inorder[:i])
    n.Right = buildTree(preorder[i+1:], inorder[i+1:])

    return &n
}
