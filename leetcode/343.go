package main

// 拆分整数

var dp [60]int
func integerBreak(n int) int {
	dp = [60]int{}

	// dp[0] = 0
	// dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i ++ {
		for j := 1; j <= i; j ++ {
			// 这个地方要注意，还有一种情况是不拆
			dp[i] = max(dp[i], max(i-j, dp[i-j]) * j)
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
