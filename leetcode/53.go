package main

// 最大子序和
// 动归
// 同时维护 max 和 运行中 max onMax
func maxSubArray(nums []int) int {
	max := nums[0]
	onMax := nums[0]
	lenn := len(nums)
	for i := 1; i < lenn; i ++ {
		onMax = maxInt(onMax+nums[i], nums[i])
		max = maxInt(onMax, max)
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

