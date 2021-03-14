package main

import "fmt"

var dp [10001]int
var status [10001]bool // false -> bottom true -> top
var max int

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}

func wiggleMaxLength(nums []int) int {
	lenn := len(nums)
	if lenn <= 1 {
		return lenn
	}
	max = 0
	flag := true // true -> top, false -> bottom
	last := nums[0]
	// top -> > left(last) && > right(right)
	// bottom -> < left(last) && < right(right)
	for i := 0; i < lenn; i++ {
		if i == 0 {
			// lenn >= 2
			for i+2 < lenn && nums[i+1] == nums[i] {
				i++
			}
			if nums[i] > nums[i+1] {
				flag = true
			} else {
				flag = false
			}
			continue
		}
		if i+1 == lenn {
			if flag {
				if nums[i] < last {
					max++
				}
			} else {
				if nums[i] > last {
					max++
				}
			}
			continue
		}
		if flag {
			// pre: top
			if nums[i] < last && nums[i] < nums[i+1] {
				last = nums[i]
				max++
				flag = false
			}
		} else {
			// pre: bottom
			if nums[i] > last && nums[i] > nums[i+1] {
				last = nums[i]
				max++
				flag = true
			}
		}
	}

	return max + 1
}

func calMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
