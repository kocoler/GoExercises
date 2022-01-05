package main

import (
	"fmt"
	"sort"
)

// 记住双指针的套路哦

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	lenn := len(nums)
	var res [][]int
	sort.Ints(nums)
	for f := 0; f < lenn-3; f++ {
		// 这个地方有个剪枝
		if f > 0 && (nums[f] == nums[f-1] || nums[f]+nums[lenn-3]+nums[lenn-2]+nums[lenn-1] < target) {
			continue
		}
		for s := f + 1; s < lenn-2; s++ {
			// 这个地方有个剪枝
			if s > f+1 && (nums[s] == nums[s-1] || nums[f]+nums[s]+nums[lenn-2]+nums[lenn-1] < target) {
				continue
			}
			psum := nums[f] + nums[s]
			l := s + 1
			r := lenn - 1
			for l < r {
				sum := psum + nums[l] + nums[r]
				if sum < target {
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else if sum > target {
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else {
					res = append(res, []int{nums[f], nums[s], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}

	return res
}
