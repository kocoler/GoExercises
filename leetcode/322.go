package main

// 二维 dp
var dp [13][10001]int
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	lenc := len(coins)

	for i := 0; i <= lenc; i ++ {
		for j := 0; j <= amount; j ++ {
			dp[i][j] = 0x3f3f3f
		}
	}

	dp[0][0] = 0

	for i := 1; i <= lenc; i ++ {
		value := coins[i-1]
		for j := 0; j <= amount; j ++ {
			if j < value {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-value] + 1)
			}
		}
	}

	if dp[lenc][amount] == 0x3f3f3f {
		return -1
	}

	return dp[lenc][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 一维 dp
var dp [10001]int
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	lenc := len(coins)

	for j := 0; j <= amount; j ++ {
		dp[j] = 0x3f3f3f
	}

	dp[0] = 0

	for i := 1; i <= lenc; i ++ {
		value := coins[i-1]
		for j := value; j <= amount; j ++ {
			dp[j] = min(dp[j], dp[j-value] + 1)
		}
	}

	if dp[amount] == 0x3f3f3f {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
