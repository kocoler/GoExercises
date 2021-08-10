package main

func numberOfArithmeticSlices(nums []int) int {
	lenn := len(nums)
	if lenn < 3 {
		return 0
	}

	dp := make([]int, lenn)

	var res int

	diff := nums[0] - nums[1]
	for i := 2; i < lenn; i ++ {
		if diff == nums[i-1] - nums[i] {
			dp[i] = dp[i-1] + 1
			res += dp[i]
		} else {
			diff = nums[i-1] - nums[i]
		}
	}

	return res
}
