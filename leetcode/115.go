package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for k := range dp {
		dp[k] = make([]int, len(t)+1)
		dp[k][len(t)] = 1
	}

	for i := len(s)-1; i >= 0; i -- {
		for j := len(t)-1; j >= 0; j -- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}
