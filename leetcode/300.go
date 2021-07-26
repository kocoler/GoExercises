package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
}

// LIS 问题 最长递增子序列问题
// 除了传统 dp, 还有贪心 + 单调栈/二分查找
func lengthOfLIS(nums []int) int {
	lenn := len(nums)
	d := []int{nums[0]}
	lend := 0
	for i := 1; i < lenn; i++ {
		if nums[i] > d[lend] {
			lend++
			d = append(d, nums[i])
		} else {
			index := sort.SearchInts(d, nums[i])
			d[index] = nums[i]
		}
	}

	return lend + 1
}
