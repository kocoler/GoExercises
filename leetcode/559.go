package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 8:16
// 8:19
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	// 层序遍历
	queue := []*Node{root}

	res := 0
	for len(queue) > 0 {
		res++
		lenq := len(queue)

		for i := 0; i < lenq; i++ {
			// node := queue[i]
			if queue[i] != nil && len(queue[i].Children) > 0 {
				queue = append(queue, queue[i].Children...)
			}

		}

		queue = queue[lenq:]
	}

	return res
}
