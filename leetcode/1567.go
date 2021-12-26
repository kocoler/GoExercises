package main

import "fmt"

func main() {
	fmt.Println(getMaxLen([]int{-1, -2, -3, 0, 1}))
}

func getMaxLen(nums []int) int {
	// 0: + length
	// 1: - length
	dp := make([][2]int, len(nums)+1)
	var res int

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			res = max(res, dp[i-1][0])
			dp[i][0] = 0
			dp[i][1] = 0
		} else if nums[i-1] > 0 {
			dp[i][0] = dp[i-1][0] + 1
			if dp[i-1][1] > 0 {
				dp[i][1] = dp[i-1][1] + 1
			}
		} else {
			// nums[i] < 0
			res = max(res, dp[i-1][0])
			if dp[i-1][1] > 0 {
				dp[i][0] = dp[i-1][1] + 1
			}
			dp[i][1] = dp[i-1][0] + 1
		}
	}

	return max(res, dp[len(nums)][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
