package main

// [2,1,5,0,4,6]
// 0 4 6
// 1 2 0 4
// 1 0 3
//
func increasingTriplet(nums []int) bool {
	stack := make([]int, 0, 2)
	stack = append(stack, nums[0])

	for i := 1; i < len(nums); i ++ {
		if len(stack) == 1 {
			if stack[0] < nums[i] {
				stack = append(stack, nums[i])
			} else {
				stack[0] = nums[i]
			}
		} else {
			if stack[1] < nums[i] {
				return true
			} else if stack[0] < nums[i] {
				stack[1] = nums[i]
			} else {
				stack[0] = nums[i]
			}
		}
	}

	return false
}
