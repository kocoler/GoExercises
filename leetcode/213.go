package main

// 两个选择 每个都算一遍
// 滚动数组：
// 每次只有两个状态相关，可以做到常量时间复杂度 同 198
var dp [102]int
func rob(nums []int) int {
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	lenn := len(nums)
	if lenn <= 2 {
		if lenn == 1 {
			return nums[0]
		}
		return max(nums[0], nums[1])
	}
	dp = [102]int{}
	// dp[1] = nums[0]
	max1 := 0
	max2 := 0

	// 0 - lenn-1
	dp[1] = nums[0]
	for i := 2; i <= lenn-1; i ++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
	}
	max1 = dp[lenn-1]

	dp[1] = 0
	// 1 - lenn
	dp[2] = nums[1]
	for i := 3; i <= lenn; i ++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
	}
	max2 = dp[lenn]


	return max(max1, max2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

