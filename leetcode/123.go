package main

func main() {
	// 7 1 5 3 6 4
	println(maxProfit([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	// 状态机：0:持有 1:不持有 2:2轮持有 3:2轮卖出
	// 两轮同时进行，如果第一轮出现比第二轮大的，那肯定可以刷新了
	lenp := len(prices)
	dp := make([][]int, lenp)
	for k := range dp {
		dp[k] = make([]int, 4)
	}

	dp[0][0] = -prices[0]
	dp[0][2] = -prices[0]
	for i := 1; i < lenp; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	return max(dp[lenp-1][3], max(dp[lenp-1][2], max(dp[lenp-1][1], dp[lenp-1][0])))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
