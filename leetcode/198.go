package main

// 打家劫舍 I
// 动归，当前状态取决于之前两个状态，拿还是不拿
var dp [102]int
func rob(nums []int) int {
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	lenn := len(nums)
	dp = [102]int{}
	dp[1] = nums[0]
	for i := 2; i <= lenn; i ++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
	}

	return dp[lenn]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 优化成常数 dp
func rob(nums []int) int {
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	lenn := len(nums)
	last := 0
	second := 0
	second = nums[0]
	for i := 2; i <= lenn; i ++ {
		second, last = max(second, last + nums[i-1]), second
	}

	return second
}

