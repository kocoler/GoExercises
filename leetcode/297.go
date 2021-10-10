package main

import (
	"bytes"
	"strconv"
	"strings"
)

// 更优的做法
// 我不会的
// 括号表示编码 + 递归下降解码
// 反正就是自己规定一些规则
// 文法吧算是

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	stack := []*TreeNode{root}
	lens := 1
	var strbuilder bytes.Buffer
	for len(stack) > 0 {
		lens = len(stack)
		for i := 0; i < lens; i++ {
			now := stack[i]
			if now == nil {
				strbuilder.WriteString("-1001 ")
			} else {
				strbuilder.WriteString(strconv.Itoa(now.Val) + " ")
				stack = append(stack, now.Left)
				stack = append(stack, now.Right)
			}
		}
		stack = stack[lens:]
	}

	return strbuilder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	strs := strings.Split(data, " ")
	if strs[len(strs)-1] == "" {
		strs = strs[:len(strs)-1]
	}
	ints := make([]int, len(strs))
	for k, v := range strs {
		ints[k], _ = strconv.Atoi(v)
	}
	leni := len(ints)

	var root TreeNode

	stack := []*TreeNode{&root}
	// lens := 1
	nows := 0
	nowi := 0

	for nowi < leni {
		if nows >= len(stack) {
			break
		}
		now := stack[nows]
		if nowi < leni {
			if ints[nowi] == -1001 {
				now.Val = -1001
			} else {
				now.Val = ints[nowi]
				now.Left = &TreeNode{}
				now.Right = &TreeNode{}
				stack = append(stack, now.Left)
				stack = append(stack, now.Right)
			}
			nows++
		}
		nowi++
	}
	fmtTree(&root)

	return &root
}

func fmtTree(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Val == -1001 {
		root.Left = nil
	}
	if root.Right != nil && root.Right.Val == -1001 {
		root.Right = nil
	}
	fmtTree(root.Left)
	fmtTree(root.Right)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
