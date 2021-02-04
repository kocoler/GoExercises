package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchRange2([]int{6,9}, 6))
	//fmt.Println(searchRange([]int{2,2}, 1))
}

func searchRange2(nums []int, target int) []int {
	left := sort.Search(len(nums), func(i int) bool {
		fmt.Println("i:", i)
		return nums[i] >= target
	})
	fmt.Println(left)
	right := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
	fmt.Println(right)
	if left <=  right - 1 {
		if nums[left] != target {
			return []int{-1, -1}
		}
		return []int{left, right - 1}
	} else {
		return []int{-1, -1}
	}
}

func searchRange(nums []int, target int) []int {
	left := 0
	lens := len(nums)
	if lens == 0 {
		return []int{-1, -1}
	}
	right := lens - 1
	for left < right {
		mid := left + (right-left)/2
		//fmt.Println("mid:", mid)
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		fmt.Println(left, right)
		if right > left && nums[right] > target {
			right --
		}
		if right > left && nums[left] < target {
			left ++
		}
		if nums[left] == target && nums[right] == target {
			break
		}
	}

	if left > right {
		return []int{-1, -1}
	}
	if left == right {
		if nums[left] != target {
			return []int{-1, -1}
		}
	}

	//fmt.Println(left, right)
	for left - 1 >= 0 && nums[left - 1] == target {
		left--
	}
	for right + 1 <= lens - 1 && nums[right + 1] == target {
		right++
	}

	return []int{left, right}
}
