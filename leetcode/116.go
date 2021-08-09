package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// 层序遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) != 0 {
		lenq := len(queue)
		// pop
		for i := 0; i < lenq - 1; i ++ {
			queue[i].Next = queue[i+1]
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if queue[lenq - 1].Left != nil {
			queue = append(queue, queue[lenq - 1].Left)
		}
		if queue[lenq - 1].Right != nil {
			queue = append(queue, queue[lenq - 1].Right)
		}
		queue = queue[lenq:]
	}

	return root
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// 空间复杂度常数(不考虑递归栈)的递归
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	buildNode(root, nil)

	return root
}

func buildNode(left, right *Node) {
	var nx *Node
	if right == nil {
		nx = nil
	} else {
		nx = right.Left
	}
	if left.Left != nil {
		left.Left.Next = left.Right
		buildNode(left.Left, left.Right)
	}
	if left.Right != nil {
		left.Right.Next = nx
		buildNode(left.Right, nx)
	}
}
