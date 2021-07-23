package main

// 组合背包（考虑顺序的完全背包）的组合问题
func combinationSum4(nums []int, target int) int {
	lenn := len(nums)
	dp := make([]int, target+1)

	dp[0] = 1
	// 先 target
	for i := 1; i <= target; i ++ {
		// 再物品
		for j := 0; j < lenn; j ++ {
			value := nums[j]
			if value <= i {
				dp[i] += dp[i-value]
			}
		}
	}


	return dp[target]
}