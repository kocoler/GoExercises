package main

import "fmt"

func main() {
	fmt.Println(findNumberOfLIS([]int{1,3,5,4,7}))
}

func findNumberOfLIS(nums []int) int {
	lenn := len(nums)
	dp := make([]int, lenn)
	dp2 := make([]int, lenn)
	maxlen := 1
	res := 0

	for i := 0; i < lenn; i ++ {
		dp[i] = 1
		dp2[i] = 1
		for j := i - 1; j >= 0; j -- {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					dp2[i] = dp2[j]
				} else if dp[j] + 1 == dp[i] {
					dp2[i] += dp2[j]
				}
			}
		}

		if maxlen < dp[i] {
			maxlen = dp[i]
			res = 0
		}
		if maxlen == dp[i] {
			res += dp2[i]
		}
	}

	return res
}
