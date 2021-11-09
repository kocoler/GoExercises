package main

import "math"

func gridGame(grid [][]int) int64 {
	// 前缀和，属实没想到
	// 不知道回溯行不行呐，似乎不行诶
	// 或者枚举在哪里转弯
	leng := len(grid[0])
	var dp [2][]int64
	for i := 0; i < 2; i++ {
		dp[i] = make([]int64, leng+1)
	}
	// dp[0][i] = grid[0][0] + grid[0][1] + ... + grid[0][i-1]
	// dp[0][i] - dp[0][j] = grid[0][j] + grid[0][j+1] + ... + grid[0][i-1]
	for i := 1; i <= leng; i++ {
		dp[0][i] = dp[0][i-1] + int64(grid[0][i-1])
		dp[1][i] = dp[1][i-1] + int64(grid[1][i-1])
	}

	// 0x3f3f3f 都小了 qwq
	var res int64 = math.MaxInt64
	for i := 1; i <= leng; i++ {
		res = min(res, max(dp[0][leng]-dp[0][i], dp[1][i-1]))
	}

	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}
