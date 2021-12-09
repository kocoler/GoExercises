package main

// 决策 dp 问题
func maximumSum(arr []int) int {
	// 1 with no delete contained
	// 2 with one delete contained
	dp := make([][2]int, len(arr))
	lena := len(arr)

	dp[0][0] = arr[0]
	dp[0][1] = arr[0] // if all less than zero
	res := arr[0]
	for i := 1; i < lena; i++ {
		dp[i][0] = max(dp[i-1][0]+arr[i], arr[i])
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]+arr[i])
		res = max(res, max(dp[i][0], dp[i][1]))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
