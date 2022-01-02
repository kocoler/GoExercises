package main

func massage(nums []int) int {
	lenn := len(nums)
	if lenn == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)

	dp[1] = nums[0]
	for i := 2; i <= lenn; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[lenn]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
