package main

// 最长公共子序列问题

func minDistance(word1 string, word2 string) int {
	lenw1 := len(word1)
	lenw2 := len(word2)

	dp := make([][]int, lenw1+1)
	for k, _ := range dp {
		dp[k] = make([]int, lenw2+1)
	}

	for i := 0; i < lenw1; i ++ {
		for j := 0; j < lenw2; j ++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return lenw1 - dp[lenw1][lenw2] + lenw2 - dp[lenw1][lenw2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

