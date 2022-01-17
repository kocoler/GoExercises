package main

import "sort"

// LIS DP 问题

func largestDivisibleSubset(nums []int) []int {
	dp := make([]int, len(nums))
	p := make([]int, len(nums))
	sort.Ints(nums)
	maxIndex := 0
	maxNum := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		p[i] = i
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					p[i] = j
				}
			}
		}
		if dp[i] > maxNum {
			maxIndex = i
			maxNum = dp[i]
		}
	}

	var res []int

	res = append(res, nums[maxIndex])
	for p[maxIndex] != maxIndex {
		res = append(res, nums[p[maxIndex]])
		maxIndex = p[maxIndex]
	}

	return res
}
