package main

import "fmt"

func main() {
	fmt.Println(palindromePartition("abc", 2))
}

func palindromePartition(s string, k int) int {
	// dp[i][k] = min(dp[j-1][k-1]+count(s[j:i]))
	// ans: dp[len(s)][k]
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for index := range dp[i] {
			dp[i][index] = i
		}
	}

	//dp[0][0] = 0
	for i := 1; i <= len(s); i ++ {
		for kk := 1; kk <= min(i, k); kk ++ {
			for j := i; j >= kk; j -- {
				dp[i][kk] = min(dp[i][kk], dp[j-1][kk-1]+count(s, j-1, i-1))
			}
		}
	}

	return dp[len(s)][k]
}

func count(s string, i, j int) int {
	ret := 0
	for i < j {
		if s[i] != s[j] {
			ret ++
		}
		i ++
		j --
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
