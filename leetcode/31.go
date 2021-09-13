package main

import "sort"

func nextPermutation(nums []int)  {
	lenn := len(nums)

	for i := lenn - 1; i > 0; i -- {
		if nums[i] > nums[i-1] {
			p := nums[i-1]

			for j := lenn - 1; j > 0; j -- {
				if nums[j] > p {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					h := i
					t := lenn - 1
					for h < t {
						nums[h], nums[t] = nums[t], nums[h]
						h ++
						t --
					}
					break
				}
			}

			return
		}
	}

	sort.Slice(nums, func (m, n int) bool { return nums[m] < nums[n] })
}
