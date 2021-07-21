package main

var lena int
var visited []bool
func canReach(arr []int, start int) bool {
	lena = len(arr)

	visited = make([]bool, len(arr))

	return dfs(arr, start)
}

func dfs(arr []int, start int) bool {
	if start < 0 || start >= lena || visited[start] {
		return false
	}

	if arr[start] == 0 {
		return true
	}

	visited[start] = true

	return dfs(arr, start - arr[start]) || dfs(arr, start + arr[start])
}
