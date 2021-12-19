package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var recursion func(*Node) *Node
	recursion = func(root *Node) *Node {
		if root == nil {
			return nil
		}
		if v, ok := m[root]; ok {
			return v
		}
		newNode := &Node{
			Val: root.Val,
		}
		m[root] = newNode
		newNode.Next = recursion(root.Next)
		newNode.Random = recursion(root.Random)
		return newNode
	}

	return recursion(head)
}
