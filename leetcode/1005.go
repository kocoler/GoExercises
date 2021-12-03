package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{-5, -6, -2, -8, 6, 9}, 4))
}

func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	sort.Ints(nums)
	nk := k
	for i := 0; i < k; i++ {
		if nums[i] >= 0 {
			break
		}
		nk--
		nums[i] = -nums[i]
	}
	sort.Ints(nums)
	for _, v := range nums {
		sum += v
	}
	//fmt.Println(nums, k, sum)
	if nk <= 0 || nk%2 == 0 {
		return sum
	}
	return sum - 2*nums[0]
}
