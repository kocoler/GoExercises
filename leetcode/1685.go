package main

import "fmt"

func main() {
	fmt.Println(getSumAbsoluteDifferences([]int{2,3,5}))
}

func getSumAbsoluteDifferences(nums []int) []int {
	// 突然发现是有序的 ... 那就前缀和了
	res := make([]int, len(nums))
	presum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		presum[i+1] = presum[i] + nums[i]
	}

	lenn := len(nums)
	for i := 0; i < lenn; i++ {
		res[i] = nums[i]*i - presum[i] + presum[lenn] - presum[i+1] - (lenn-i-1)*nums[i]
	}

	return res
}
