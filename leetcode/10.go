package main

import "fmt"

// 动态规划
// 比较经典了，竟然还做了好久，我没救了呜呜呜

func main() {
	fmt.Println(isMatch("a", ".*"))
}

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if i > 0 && ((s[i-1] == p[j-2]) || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j] || dp[i-1][j-1] || dp[i][j-1]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
