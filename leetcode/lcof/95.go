package main

func longestCommonSubsequence(text1 string, text2 string) int {
	lent1 := len(text1)
	lent2 := len(text2)
	dp := make([][]int, lent1+1)
	for i := 0; i <= lent1; i ++ {
		dp[i] = make([]int, lent2+1)
	}

	for i := 0; i < lent1; i ++ {
		c := text1[i]
		for j := 0; j < lent2; j ++ {
			if c == text2[j] {
				dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j] + 1)
			} else {
				// dp[i+1][j+1] = max(max(max(dp[i+1][j+1], dp[i][j]), dp[i][j+1]), dp[i+1][j])
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[lent1][lent2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
