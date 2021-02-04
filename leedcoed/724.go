package main

// var forward [10000]int
var reverse [10000]int

func pivotIndex(nums []int) int {
	// forward = [10000]int{}
	reverse = [10000]int{}
	lenn := len(nums)
	if lenn < 2 {
		if lenn == 0 {
			return -1;
		}
		return 0;
	}

	// forward[0] = 0
	// forward[1] = nums[0]
	// for i := 2; i < lenn; i ++ {
	//     forward[i] = forward[i-1] + nums[i - 1]
	// }
	reverse[lenn - 1] = 0
	reverse[lenn - 2] = nums[lenn - 1]
	for i := lenn - 3; i >= 0; i -- {
		reverse[i] = reverse[i+1] + nums[i+1]
	}

	if reverse[0] == 0 {
		return 0
	}

	var count int
	for i := 1; i < lenn; i ++ {
		count += nums[i-1]
		if count == reverse[i] {
			return i
		}
	}

	return -1
}
