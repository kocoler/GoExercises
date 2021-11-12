package main

import (
	"fmt"
	"sort"
)

func main() {
	// 0 2 3 6 10 15
	// 0 1 2 3 4  5
	fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3}))
}

// 前缀和+!二分查找! 或者 滑动窗口
// 这里复习一下前缀和和二分查找
// 面试的话肯定就滑动窗口了，时间复杂度低还容易写
func minSubArrayLen(target int, nums []int) int {
	res := 0x3f3f3f
	lenn := len(nums)

	dp := make([]int, lenn+1)
	for i := 1; i <= lenn; i ++ {
		dp[i] = dp[i-1] + nums[i-1]
		if dp[i] >= target {
			find := dp[i] - target
			index := sort.Search(i+1, func(j int) bool {
				return dp[j] >= find
			})

			// 这里还是要判断一下
			// 如果刚好有相等的，那可以直接不算
			// 如果没有，则需要加上他~
			if dp[index] == find {
				res = min(res, i - index)
			} else {
				res = min(res, i - index + 1)
			}
		}
	}

	if res == 0x3f3f3f {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
