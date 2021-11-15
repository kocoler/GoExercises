package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{9,10,9,-7,-4,-8,2,-6}, 5))
}

func maxSlidingWindow(nums []int, k int) []int {
	lenn := len(nums)
	res := make([]int, lenn - k + 1)

	// max stack
	stack := []int{}
	for i := 0; i < k; i ++ {
		for len(stack) != 0 && stack[len(stack) - 1] < nums[i] {
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, nums[i])
	}
	res[0] = stack[0]
fmt.Println(stack)
	for i := k; i < lenn; i ++ {
		// num := nums[i]
		out := nums[i-k]
		if out >= stack[0] {
			for index, v := range stack {
				if v == out {
					stack = append(stack[:index], stack[index+1:] ...)
					break
				}
			}
		}
		fmt.Println(nums[i], stack)
		for len(stack) != 0 && stack[len(stack) - 1] < nums[i] {
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, nums[i])

		res[i-k+1] = stack[0]
	}

	return res
}
