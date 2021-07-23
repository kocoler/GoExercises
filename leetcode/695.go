package main

var visited [51][51]bool
var res int
var g [][]int
var direction = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var lenr int
var lenc int
// [[1,1,0,0,0]
// ,[1,1,0,0,0]
// ,[0,0,0,1,1]
// ,[0,0,0,1,1]]
func maxAreaOfIsland(grid [][]int) int {
	visited = [51][51]bool{}
	res = 0
	lenr = len(grid)
	lenc = len(grid[0])
	g = grid

	for i := 0; i < lenr; i ++ {
		for j := 0; j < lenc; j ++ {
			if !visited[i][j] && grid[i][j] == 1 {
				dfs(i, j)
				res = max(res, n)
				n = 0
			}
		}
	}

	return res
}

var n int
func dfs(sr, sc int) {
	visited[sr][sc] = true

	n ++

	for i := 0; i < 4; i ++ {
		nsr := sr + direction[i][0]
		nsc := sc + direction[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && !visited[nsr][nsc] && g[nsr][nsc] == 1 {
			dfs(nsr, nsc)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
