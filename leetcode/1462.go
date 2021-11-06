package main

// Floyd 的判断路径存在用法

var graph [101][101]bool

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	lenq := len(queries)
	res := make([]bool, lenq)
	graph = [101][101]bool{}

	var x, y int
	for k := range prerequisites {
		x, y = prerequisites[k][0], prerequisites[k][1]
		graph[x][y] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				if graph[i][j] || (graph[i][k] && graph[k][j]) {
					graph[i][j] = true
				}
			}
		}
	}

	for k := range queries {
		x, y = queries[k][0], queries[k][1]
		if graph[x][y] || x == y {
			res[k] = true
		}
	}

	return res
}
