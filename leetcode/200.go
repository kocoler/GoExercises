package main

var res int
var g [][]byte
var direction = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var lenr int
var lenc int
func numIslands(grid [][]byte) int {
	res = 0
	lenr = len(grid)
	lenc = len(grid[0])
	g = grid

	for i := 0; i < lenr; i ++ {
		for j := 0; j < lenc; j ++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				res ++
			}
		}
	}

	return res
}

func dfs(sr, sc int) {
	g[sr][sc] = '0'

	for i := 0; i < 4; i ++ {
		nsr := sr + direction[i][0]
		nsc := sc + direction[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && g[nsr][nsc] == '1' {
			dfs(nsr, nsc)
		}
	}
}

