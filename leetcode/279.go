package main

// 完全背包的最小值问题
func numSquares(n int) int {
	lenn := int(math.Sqrt(float64(n))) + 1
	dp := make([]int, n+1)

	for i := 0; i <= n; i ++ {
		dp[i] = 0x3f3f3f
	}

	dp[0] = 0
	for i := 1; i <= lenn; i ++ {
		sqrtI := i * i
		for j := sqrtI; j <= n; j ++ {
			dp[j] = min(dp[j], dp[j-sqrtI] + 1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
