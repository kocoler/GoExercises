package main

// dfs

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
var m map[int]*Node
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	m = make(map[int]*Node)

	root := &Node{Val: node.Val}
	m[node.Val] = root

	root.Neighbors =  dfs(node.Neighbors)

	return root
}

func dfs(neighbors []*Node) []*Node {
	res := []*Node{}

	for k := range neighbors {
		if _, ok := m[neighbors[k].Val]; ok {
			res = append(res, m[neighbors[k].Val])
		} else {
			node := &Node{Val: neighbors[k].Val}
			m[node.Val] = node
			node.Neighbors = dfs(neighbors[k].Neighbors)
			res = append(res, node)
		}
	}

	return res
}
