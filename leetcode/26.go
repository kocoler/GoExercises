package main

func removeDuplicates(nums []int) int {
	// 还是有序的 ...
	// 怪不得是 easy
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			l++
			nums[l] = nums[i]
		}
	}
	return l + 1
}
