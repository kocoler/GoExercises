package main

import "sort"

func targetIndices(nums []int, target int) []int {
	var res []int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
		}
	}

	return res
}
