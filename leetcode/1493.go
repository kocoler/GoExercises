package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 1}))
}

func longestSubarray(nums []int) int {
	// 滑动窗口

	l := 0
	r := 0
	res := 0
	others := 0

	for r < len(nums) {
		if nums[r] != 1 {
			others++
		}
		if others > 1 {
			for l < len(nums) && nums[l] == 1 {
				l++
			}
			l++
			others--
		}
		res = max(res, r-l+others-1)
		fmt.Println(l, r, res)

		r++
	}
	res = max(res, r-l+others-1)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
