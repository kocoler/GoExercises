package main

// 单调栈

func findUnsortedSubarray(nums []int) int {
	var stack []int
	// var res int

	lenn := len(nums)

	top := -0x3f3f3f
	left := 0x3f3f3f
	right := 0
	for i := 0; i < lenn; i ++ {
		if top <= nums[i] {
			stack = append(stack, nums[i])
			top = nums[i]
		} else {
			right = i
			// index := 0
			// flag := false

			// 优化
			if left == 0 || (left != 0x3f3f3f && stack[left] <= nums[i]) {
				continue
			}
			if stack[0] > nums[i] {
				left = 0
				continue
			}

			for j := len(stack) - 1; j >= 0; j -- {
				if stack[j] <= nums[i] {
					// index = j
					left = min(left, j+1)
					// flag = true
					break
				}
			}
			// if !flag {
			//     left = 0
			// }
		}
	}

	if left == 0x3f3f3f {
		return right - 0
	}

	return right - left + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
