package main

// 二分是不可以的哦，这道题是不可以二分的哦
// 因为二分的目标只是：不断的缩小范围
// 但是这道题需要的是所有情况的遍历结果
// 最大最小动态规划问题
// i-j 是 i-j 所需的最小金额，**每一步都要是最坏的情况**
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			now := 0x3f3f3f
			for k := i; k < j; k++ {
				now = min(now, k+max(dp[i][k-1], dp[k+1][j]))
			}
			dp[i][j] = now
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
