package main

// 494. 目标和
// 这道题最初想用 dp[i][j] 表示 i 个元素 和为 j 的数目
// 但是要考虑负数 (j-num[i-1]) 的情况，所以行不通
// 参考官方题解，引入 **neg** 即 负数之和，这样就不需要考虑负数的情况
// 还可以优化滚动数组，懒了
var dp [21][1002]int
func findTargetSumWays(nums []int, target int) int {
	// dp
	// dp[i][j]
	sum := 0
	for _, v := range nums {
		sum += v
	}
	lenn := len(nums)

	if sum - target < 0 || (sum - target) % 2 != 0 {
		return 0
	}

	dp = [21][1002]int{}

	neg := (sum - target) / 2
	//[[1 0 0 0 0 0 0 0 0 0]
	// [0 1 0 0 0 0 0 0 0 0]
	// [1 0 1 0 0 0 0 0 0 0]
	// [0 2 0 1 0 0 0 0 0 0]
	// [2 0 3 0 1 0 0 0 0 0]
	// [0 5 0 4 0 1 0 0 0 0]]

	dp[0][0] = 1
	for i := 1; i <= lenn; i ++ {
		for j := 0; j <= neg; j ++ {
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			}
		}
	}
	// fmt.Println(dp)

	return dp[lenn][neg]
}

