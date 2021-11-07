package main

import "fmt"

// 这个方法有点笨拙，需要再做一遍嘞

func main() {
	fmt.Println(waysToMakeFair([]int{1,2,3}))
}

func waysToMakeFair(nums []int) int {
	//   2 1 6 4
	//0: 2 0 8 0 -> 8
	//1: 0 1 0 5 -> 5
	//remove 2
	//   2 6 4
	//0: 2 0 6 -> 6
	//1: 0 6 0 -> 6
	//odd: even[lenn] - even[i] + odd[i](- nums[i])
	//even: odd[lenn] - odd[i] + even[i]
	// dp[n][0] => odd
	// dp[n][1] => even
	res := 0
	lenn := len(nums)
	if lenn == 1 {
		return 1
	}

	dp := make([][2]int, lenn)

	dp[0][0] = nums[0]
	dp[0][1] = 0
	dp[1][0] = nums[0]
	dp[1][1] = nums[1]
	for i := 2; i < lenn; i ++ {
		// todo: optimization
		if (i-1) % 2 == 0 {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-2][1] + nums[i]
		} else {
			dp[i][0] = dp[i-2][0] + nums[i]
			dp[i][1] = dp[i-1][1]
		}
	}

	odd := dp[lenn-1][0]
	even := dp[lenn-1][1]
	for i := 0; i < lenn; i ++ {
		o := even - dp[i][1] + dp[i][0]
		e := odd - dp[i][0] + dp[i][1]
		if i % 2 != 0 {
			e -= nums[i]
		} else {
			o -= nums[i]
		}
		if o == e {
			res ++
		}
	}

	return res
}
