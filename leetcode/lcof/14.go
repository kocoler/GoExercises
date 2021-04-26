package lcof

// 1 1 1 1 1
//  2   2  1
var dp [59]int
func cuttingRope(n int) int {
	dp = [59]int{0, 1, 1}

	for i := 3; i <= n; i ++ {
		dp[i] = dp[i-1]
		for j := 1; j <= i/2; j ++ {
			l := max(dp[j], j)
			r := max(dp[i-j], i-j)
			dp[i] = max(dp[i], l*r)
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
