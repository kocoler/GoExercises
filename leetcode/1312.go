package main

import "fmt"

func main() {
	fmt.Println(minInsertions("no"))
}

// LSC 问题

func minInsertions(s string) int {
	lens := len(s)
	dp := make([][]int, lens+1)
	for k := range dp {
		dp[k] = make([]int, lens+1)
	}
	r := lens - 1
	for i := 1; i <= lens; i++ {
		for j := 1; j <= lens; j++ {
			ii, ij := i-1, r-j+1
			if s[ii] == s[ij] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return lens - dp[lens][lens]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
