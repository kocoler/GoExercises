package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1, 3, 2, 2, 5, 2, 3, 7
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

// 8:55
func findLHS(nums []int) int {
	lenn := len(nums)
	res := 0

	sort.Ints(nums)

	left := 0
	right := 0

	for right < lenn {
		for left < right && nums[right]-nums[left] > 1 {
			left++
		}
		if nums[right]-nums[left] == 1 {
			res = max(res, right-left+1)
		}
		right++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
