package main

// 贪心 DFS/BFS

func canJump(nums []int) bool {
	m := nums[0]
	if len(nums) == 1 && nums[0] == 0 {
		return true
	}
	for i := 0; i <= m ; i ++ {
		if m >= len(nums) {
			return true
		}
		if nums[i] + i > m {
			m = nums[i] + i
		}
	}


	return m+1 >= len(nums)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
