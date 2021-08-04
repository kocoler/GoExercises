package main

// 三色标记 dfs

var g [][]int
var m map[int]int // 0 -> none 1 -> true 2 -> false

func eventualSafeNodes(graph [][]int) []int {
	m = make(map[int]int)
	g = graph

	num := len(graph)

	for i := 0; i < num; i ++ {
		if m[i] == 0 {
			dfs(i)
		}
	}

	var res []int
	for i := 0; i < num; i ++ {
		if m[i] == 1 {
			res = append(res, i)
		}
	}

	return res
}

func dfs(i int) bool {
	if m[i] == 1 {
		return true
	}
	if m[i] == 2 {
		return false
	}

	// if len(g[i]) == 0 {
	//     m[i] = 1
	//     return true
	// }

	flag := true
	m[i] = 2
	for _, v := range g[i] {
		if dfs(v) {
			m[v] = 1
		} else {
			flag = false
			m[v] = 2
		}
	}

	if flag {
		m[i] = 1
	} else {
		m[i] = 2
	}

	return flag
}
