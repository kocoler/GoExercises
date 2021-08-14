package main

var m [][]int
var n int
func rotate(matrix [][]int)  {
	n = len(matrix) - 1
	m = matrix

	dfs(0, 0)
}

func dfs(ar, ac int) {
	if ar > n / 2 {
		return
	}
	br := ar; bc := n - ac
	cr := n - ar; cc := n - ac
	dr := n - ar; dc := ac

	end := bc - ar
	for i := 0; i < end; i ++ {
		m[ar][ac+i], m[br+i][bc], m[cr][cc-i], m[dr-i][dc] = m[dr-i][dc], m[ar][ac+i], m[br+i][bc], m[cr][cc-i]
	}

	dfs(ar + 1, ac + 1)
}
