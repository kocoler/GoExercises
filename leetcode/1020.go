package main

var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var lenr int
var lenc int
var g [][]int
func numEnclaves(grid [][]int) int {
	lenr = len(grid)
	lenc = len(grid[0])
	g = grid
	for i := 0; i < lenr; i ++ {
		if g[i][0] != 0 {
			dfs(i, 0)
		}
		if g[i][lenc-1] != 0 {
			dfs(i, lenc-1)
		}
	}
	for j := 0; j < lenc; j ++ {
		if g[0][j] != 0 {
			dfs(0, j)
		}
		if g[lenr-1][j] != 0 {
			dfs(lenr-1, j)
		}
	}

	res := 0
	for i := 1; i < lenr; i ++ {
		for j := 1; j < lenc; j ++ {
			if g[i][j] == 1 {
				res ++
			}
		}
	}

	return res
}

func dfs(r, c int) {
	g[r][c] = 2
	for i := 0; i < 4; i++ {
		nsr := r + dis[i][0]
		nsc := c + dis[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && g[nsr][nsc] == 1 {
			dfs(nsr, nsc)
		}
	}
}
