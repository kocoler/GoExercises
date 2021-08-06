package main

// 乘最多水的容器
// 经典双指针，要想到双指针的条件
// 两边开始，判定条件

func maxArea(height []int) int {
	lenh := len(height)

	l := 0
	r := lenh - 1

	res := 0
	for l < r {
		// res = max(min(height[l], height[r]) * (r - l), res)
		if height[l] < height[r] {
			res = max(height[l] * (r - l), res)
			l ++
		} else {
			res = max(height[r] * (r - l), res)
			r --
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
