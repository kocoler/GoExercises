package main

// 每次操作既可以理解为使 n-1 个元素增加 1，也可以理解使 1 个元素减少 1
// ...

func minMoves(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}

	res := 0
	for _, v := range nums {
		res += v - min
	}

	return res
}
