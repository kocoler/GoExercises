package main

// 比较神奇的做法，直接将访问过的单位置 0
// 就可以省略 visited 数组了

// var visited [51][51]bool
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
	// visited = [51][51]bool{}
	res = 0
	lenr = len(grid)
	lenc = len(grid[0])
	g = grid

	for i := 0; i < lenr; i ++ {
		for j := 0; j < lenc; j ++ {
			if grid[i][j] == 1 {
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
	g[sr][sc] = 0

	n ++

	for i := 0; i < 4; i ++ {
		nsr := sr + direction[i][0]
		nsc := sc + direction[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && g[nsr][nsc] == 1 {
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
