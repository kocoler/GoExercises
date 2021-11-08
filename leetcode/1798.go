package main

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	lenc := len(coins)
	res := 1
	for i := 0; i < lenc; i++ {
		// 如果 coins[i] 比 res + 1 大，也就是现在能生成的数 + 1 还大的话
		// 那么肯定会有 res + 1 是生成不了的，因为组合不出来
		// 但是如果是小的话，可以通过 coins[i] - n 来组合
		if coins[i] > res+1 {
			break
		}
		res += coins[i]
	}

	return res
}
