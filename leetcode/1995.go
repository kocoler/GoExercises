package main

import (
	"fmt"
)

// 用这道题复习一下 三数之和
// 复习不了 三数之和是有序的，这道题不能排序，debug 半天 尬住了

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

func countQuadruplets(nums []int) int {
	var res int

	m := make(map[int]int)                // fourth
	for t := len(nums) - 2; t >= 2; t-- { // third
		m[nums[t+1]]++
		for f := 0; f < t-1; f++ { // first
			for s := f + 1; s < t; s++ { // second
				sum := nums[t] + nums[f] + nums[s]
				res += m[sum]
			}
		}
	}

	return res
}
