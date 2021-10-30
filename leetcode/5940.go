package main

import "fmt"

func main() {
	fmt.Println(minimumDeletions([]int{-1, -53, 93, -42, 37, 94, 97, 82, 46, 42, -99, 56, -76, -66, -67, -13, 10, 66, 85, -28}))
}

func minimumDeletions(nums []int) int {
	mini := 0
	maxi := 0
	minnum := nums[0]
	maxnum := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < minnum {
			minnum = nums[i]
			mini = i
		}
		if nums[i] > maxnum {
			maxnum = nums[i]
			maxi = i
		}
	}

	minres := 0
	maxres := 0
	minres = min(mini+1, len(nums)-mini)
	maxres = min(maxi+1, len(nums)-maxi)

	// remove from left
	leftMin := mini + 1
	if maxi > mini {
		leftMin = maxi + 1
	}
	rightMin := len(nums) - mini
	if maxi < mini {
		rightMin = len(nums) - maxi
	}

	return min(min(leftMin, rightMin), minres+maxres)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 17 11
