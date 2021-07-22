package main

// 前缀和 + HashMap

func numSubarraysWithSum(nums []int, goal int) int {
	lenn := len(nums)
	m := make(map[int]int)

	res := 0
	if (nums[0] == 1 && goal == 1) || (nums[0] == 0 && goal == 0) {
		res ++
	}
	m[0] ++

	for i := 1; i < lenn; i ++ {
		m[nums[i-1]] ++

		nums[i] += nums[i-1]

		res += m[nums[i] - goal]
	}

	return res
}
