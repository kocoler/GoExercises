package main

// 除自身以外数组的乘积
// 优化到时间复杂度 O(1) 的办法: 把答案数组当作 dp 数组来用
var dp [100001][2]int
func productExceptSelf(nums []int) []int {
	lenn := len(nums)
	dp = [100001][2]int{}
	res := make([]int, lenn)
	dp[0][0] = 1
	dp[lenn-1][1] = 1
	for i := 1; i < lenn; i ++ {
		dp[i][0] = dp[i-1][0] * nums[i-1]
	}
	for i := lenn-2; i >= 0; i -- {
		dp[i][1] = dp[i+1][1] * nums[i+1]
		res[i] = dp[i][0] * dp[i][1]
	}
	res[lenn-1] = dp[lenn-1][0] * dp[lenn-1][1]
	return res
}

// 前缀和做法
func productExceptSelf(nums []int) []int {
	lenn := len(nums)

	res := make([]int, lenn)
	res[lenn-1] = nums[lenn-1]
	for i := lenn - 2; i > 0; i -- {
		res[i] = nums[i] * res[i+1]
	}

	res[0] = res[1]
	for i := 1; i < lenn - 1; i ++ {
		res[i] = nums[i-1] * res[i+1]
		nums[i] = nums[i] * nums[i-1]
	}

	res[lenn-1] = nums[lenn-2]

	return res
}
