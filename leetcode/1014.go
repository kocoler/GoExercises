package main

// O(n^2) => O(3n) => O(2n) O(n) => O(n) O(1)
// 注意 dp 算式的拆分
func maxScoreSightseeingPair(values []int) int {
	lenv := len(values)

	dp := make([]int, lenv)

	dp[0] = values[0] + 0
	// dp[0][1] = values[0] - 0

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 1; i < lenv; i ++ {
		dp[i] = max(dp[i-1], values[i] + i)
	}

	// dp[lenv-1][1] = values[lenv-1] - lenv + 1

	// for i := lenv-2; i >= 0; i -- {
	//     dp[i][1] = max(dp[i+1][1], values[i] - i)
	// }

	res := 0
	// for i := 0; i < lenv-1; i ++ {
	//     res = max(res, dp[i][0] + dp[i+1][1])
	// }
	r := values[lenv-1] - lenv + 1
	for i := lenv - 2; i >= 0; i -- {
		res = max(res, dp[i] + r)
		r = max(r, values[i] - i)
	}

	return res
}

