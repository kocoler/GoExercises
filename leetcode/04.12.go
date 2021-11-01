package main

import "fmt"

var res int
var target int
var m map[string]struct{}

func pathSum(root *TreeNode, sum int) int {
	res = 0
	target = sum
	m = make(map[string]struct{})

	dfs(root, "", 0)

	return res
}

func dfs(root *TreeNode, path string, now int) {
	if root == nil {
		return
	}
	npath := fmt.Sprintf("%s %d", path, now)
	if _, ok := m[npath]; ok {
		return
	}
	m[npath] = struct{}{}
	fmt.Println("dfs", root.Val, now)
	now += root.Val
	if now == target {
		res++
		// fmt.Println(now, root.Val)
	}
	l := path + "0"
	r := path + "1"
	dfs(root.Left, l, now)
	dfs(root.Right, r, now)
	dfs(root.Left, l, 0)
	dfs(root.Right, r, 0)
}
