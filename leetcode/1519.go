package main

import "fmt"

// ERROR

// DFS

// 0 - 2 - 1
/// - 3

func main() {
	//fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(countSubTrees(4, [][]int{{0, 2}, {0, 3}, {1, 2}}, "aeed"))
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	res := make([]int, n)

	// build graph
	graph := make(map[int][]int)
	visited := map[int]bool{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	// dfs
	var dfs func(int) [26]int
	dfs = func(node int) [26]int {
		visited[node] = true
		count := [26]int{}
		count[labels[node]-'a']++
		for _, child := range graph[node] {
			//if visited[child] || child < node {
			if visited[child] {
				continue
			}
			for k, v := range dfs(child) {
				count[k] += v
			}
		}

		res[node] = count[labels[node]-'a']

		return count
	}

	// 最重要的就是从 0 开始，这样就保证了这个树的构造！
	dfs(0)

	return res
}
