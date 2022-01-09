package main

import "math/rand"

func partition(nums []int, l, r int) int {
	i := l - 1
	pivot := nums[r]

	for j := l; j < r; j++ {
		if nums[j] > pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]

	return i + 1
}

func randomPartition(nums []int, l, r int) int {
	i := rand.Intn(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]

	return partition(nums, l, r)
}

func qsort(nums []int, l, r, k int) {
	if l < r {
		i := randomPartition(nums, l, r)
		if i == k-1 {
			return
		} else if i < k {
			qsort(nums, i+1, r, k)
		} else {
			qsort(nums, l, i-1, k)
		}
	}
}

func findKthLargest(nums []int, k int) int {
	qsort(nums, 0, len(nums)-1, k)
	return nums[k-1]
}
