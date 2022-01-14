package main

// 能往 4 个方向走的时候 dp 会丢掉一些状态
// 所以 4 个方向 dfs + 记忆化 会好一点

// 但其实也可以 dp, 用 拓扑排序 的思想来做这个状态转移的过程
// 因为他的初始状态是 入度 为 0 那个点，逐渐往外走的~
// 而大于的条件可以看作是 一条边
// 那么把 dp[r][c] 定义成以 r, c 开始的最长上升序列就可以了
// 反过来定义也可以的

var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func longestIncreasingPath(matrix [][]int) int {
	maxr, maxc := len(matrix), len(matrix[0])
	dp := make([][]int, maxr+1) // 这里的 dp 就是记忆化了
	for k := range dp {
		dp[k] = make([]int, maxc+1)
	}
	var res int

	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if dp[r][c] != 0 {
			return dp[r][c]
		}

		nx := 0
		for i := 0; i < 4; i++ {
			nxr, nxc := r+dis[i][0], c+dis[i][1]

			if nxr >= 0 && nxc >= 0 && nxr < maxr && nxc < maxc && matrix[nxr][nxc] > matrix[r][c] {
				nx = max(nx, dfs(nxr, nxc))
			}
		}

		dp[r][c] = nx + 1

		return nx + 1
	}

	for i := 0; i < maxr; i++ {
		for j := 0; j < maxc; j++ {
			res = max(res, dfs(i, j))
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
