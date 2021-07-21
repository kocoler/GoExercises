package main

func maxProfit(prices []int, fee int) int {
	lenp := len(prices)
	dp := make([][2]int, lenp)

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < lenp; i ++ {
		dp[i][0] = max(dp[i-1][1] + prices[i] - fee, dp[i-1][0])
		dp[i][1] = max(dp[i-1][0] - prices[i], dp[i-1][1])
	}

	return max(dp[lenp-1][0], dp[lenp-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
