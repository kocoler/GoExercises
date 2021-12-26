package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)>>1
		// nums[mid] < nums[r], nums[mid] > nums[r]
		// nums[mid] < nums[l] => nums[mid] < nums[r] , nums[mid] > nums[l] => nums[mid] > nums[r]
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}

	return nums[l]
}
