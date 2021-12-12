package main

// 滑动窗口

func numberOfSubarrays(nums []int, k int) int {
	num := 0
	res := 0
	r := 0
	l := 0

	for r < len(nums) {
		if nums[r]%2 != 0 {
			num++
			if num == k {
				res++
				for nums[l]%2 == 0 {
					l++
				}
				l++
			}
		}
		r++
	}

	return res
}
