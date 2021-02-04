package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	lenn := len(nums)
	if lenn == 0 {
		return false
	}
	record := make(map[int]int)

	count := k+1
	if count > lenn {
		count = lenn
	}
	for i := 0; i < count; i ++ {
		record[nums[i]] ++
		if record[nums[i]] > 1 {
			return true
		}
	}

	for i := count; i < lenn; i ++ {
		record[nums[i - count]] --
		record[nums[i]] ++
		fmt.Println(record)
		if record[nums[i]] > 1 {
			return true
		}
	}

	return false
}
