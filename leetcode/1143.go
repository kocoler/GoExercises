package main

var dp [1001][1001]int

func longestCommonSubsequence(text1 string, text2 string) int {
	// db two-string
	len1 := len(text1)
	len2 := len(text2)

	dp = [1001][1001]int{}
	for i := 0; i < len1; i ++ {
		for j := 0; j < len2; j ++ {
			if text1[i] == text2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
					continue
				}
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if i == 0 || j == 0 {
					if i == 0 {
						if j == 0 {
							dp[i][j] = 0
						} else {
							dp[i][j] = dp[i][j-1]
						}
					} else {
						dp[i][j] = dp[i-1][j]
					}
					continue
				}
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len1-1][len2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
