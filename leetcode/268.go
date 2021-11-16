package main

func missingNumber(nums []int) int {
	lenn := len(nums)

	sum := 0
	for k := range nums {
		sum += nums[k]
	}

	return lenn * (lenn+1) / 2 - sum
}
