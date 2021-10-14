package main

import "fmt"

func main() {
	fmt.Println(binsearch([]int{5,7,8,8,10}, 10))
}

func binsearch(nums[]int, target int) int {
	l := 0
	r := len(nums) - 1

	// search left
	for l < r {
		// left
		mid := l + (r - l + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}

	if nums[r] != target {
		return -1
	}

	return r
}
