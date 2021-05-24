package main

import "sort"

// 1 2 2
// 1 2 2
// 1 2
// 1
// 2
// 2 2
func subsetsWithDup(nums []int) [][]int {
	lenn := len(nums)
	sort.Ints(nums)
	dfs := func(int){}
	var res [][]int
	var cur []int
	dfs = func(n int) {
		if n >= lenn {
			t := make([]int, len(cur))
			copy(t, cur)
			res = append(res, t)
			return
		}
		n ++
		for n < lenn && nums[n] == nums[n-1] {
			n ++
		}
		dfs(n)
		cur = cur[:len(cur)-1]
		n ++
		for n < lenn && nums[n] == nums[n-1] {
			n ++
		}
		dfs(n)
	}
	dfs(0)
	return res
}
