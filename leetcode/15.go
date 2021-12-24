package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	lenn := len(nums)
	var res [][]int
	for f := 0; f < lenn; f++ {
		if f > 0 && nums[f] == nums[f-1] {
			continue
		}
		s := f + 1
		t := lenn - 1
		target := -nums[f]
		for ; s < lenn; s++ {
			if s > f+1 && nums[s] == nums[s-1] {
				continue
			}
			// !!! nums[t] + nums[s] > target
			// 因为 s 是从左到右的，如果大于，那么肯定是没有合适的，那么 s ++ 之后再从这里开始继续匹配置
			// 这样单次复杂度是 n
			// 总复杂度接近 n^2
			for t > s && nums[t]+nums[s] > target {
				t--
			}
			if s == t {
				break
			}
			if nums[s]+nums[t] == target {
				res = append(res, []int{nums[f], nums[s], nums[t]})
			}
		}
	}

	return res
}
