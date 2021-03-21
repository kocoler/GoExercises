package main

var dp[501][501]int
func PredictTheWinner(nums []int) bool {
	dp  = [501][501]int{}
	lenn := len(nums)
	for i := lenn-1; i >= 0; i -- {
		for j := i; j < lenn; j ++ {
			if i == j {
				dp[i][j] = nums[i]
			} else {
				dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
			}
		}
	}

	return dp[0][lenn-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
