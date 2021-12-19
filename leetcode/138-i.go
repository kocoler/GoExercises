package main

// 这个迭代写法挺神奇的
// 这提供了一个思路：当链表进行操作时：
// 我们可以用 node.Next 来保存新的东西
// node.Next.Next 来保存之前的节点
// 之后就可以通过操作来还原

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
	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{
			Val:  node.Val,
			Next: node.Next,
		}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		// 还原！！！！
		pre := node.Next
		node.Next = node.Next.Next
		if pre != nil {
			pre.Next = pre.Next.Next
		}
	}

	return newHead
}
