package main

import "fmt"

var dp [12][12]bool
var n int
var str string
var dp2 [12]int

func main() {
	fmt.Println(minCut("aabbccddee"))
}

func minCut(s string) int {
	lens := len(s)
	n = lens
	s = fmt.Sprintf(" %s", s)
	str = s
	dp = [12][12]bool{}
	dp2 = [12]int{}
	for i := lens; i >= 1; i-- {
		for j := i; j <= lens; j++ {
			if i == j {
				dp[i][j] = true
			}
			if s[i] == s[j] {
				if (i+1 == j) || dp[i+1][j-1] == true {
					dp[i][j] = true
				}
			}
		}
	}

	for i := 1; i <= lens; i ++ {
		if dp[1][i] {
			dp2[i] = 1
		} else {
			min := dp2[i-1]+1
			for j := i-1; j >= 1; j -- {
				if dp[j][i] {
					min = calmin(min, dp2[j-1]+1)
				}
			}
			dp2[i] = min
		}
	}
	fmt.Println(dp2)

	return dp2[lens] - 1
}

func calmin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
