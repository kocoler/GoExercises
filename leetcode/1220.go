package main

// 还可以矩阵快速幂
// 牛！

func countVowelPermutation(n int) int {
	mod := int(1e9+7)
	// a, e, i, o, u
	// a -> e
	// e -> a/i
	// i -> a/e/o/u
	// o -> i/u
	// u -> a
	dp := [5]int{1,1,1,1,1}

	for i := 1; i < n; i ++ {
		dp[0], dp[1], dp[2], dp[3], dp[4] =
			((dp[1] + dp[4])%mod + dp[2]) % mod,
			(dp[0] + dp[2]) % mod,
			(dp[1] + dp[3]) % mod,
			dp[2],
			(dp[2] + dp[3]) % mod
	}

	var res int
	for _, v := range dp {
		res += v
		res %= mod
	}

	return res
}
