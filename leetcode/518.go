package main

// 完全背包的组合问题
func change(amount int, coins []int) int {
	lenc := len(coins)
	dp := make([]int, amount+1)

	dp[0] = 1
	for i := 0; i < lenc; i ++ {
		value := coins[i]
		for j := value; j <= amount; j ++ {
			dp[j] += dp[j - value]
		}
	}

	return dp[amount]
}
