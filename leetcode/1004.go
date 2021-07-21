package main

// 滑动窗口
// 更神奇的做法: 前缀和+二分查找
func longestOnes(nums []int, k int) int {
	lenn := len(nums)
	res := 0
	l := 0
	r := 0
	b := 0
	for r < lenn {
		if nums[r] == 1 {
			r ++
		} else {
			// nums[r] == 0
			if b < k {
				b ++
				r ++
			} else {
				res = max(res, r - l)
				for nums[l] != 0 {
					l ++
				}
				l ++
				b --
			}
		}
	}
	res = max(res, r - l)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
