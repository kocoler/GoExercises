package main

// 回溯
// 经典子集
func subsets(nums []int) [][]int {
	var res [][]int
	lenn := len(nums)
	var cur []int
	var dfs func(n int)
	dfs = func(n int) {
		if n == lenn {
			r := make([]int, len(cur))
			copy(r, cur)
			res = append(res, r)
			return
		}
		cur = append(cur, nums[n])
		dfs(n + 1)
		lenc := len(cur)
		cur = cur[:lenc-1]
		dfs(n + 1)
	}

	dfs(0)

	return res
}

