package main

import (
	"sort"
	"strconv"
)

// 字符串动规
// 感觉比较麻烦，要处理 1，2，2+长度
// 还有遍历顺序，很巧妙，这样同时限制了长度

func checkValidString(s string) bool {
	lens := len(s)

	dp := make([][]bool, lens)
	for i := 0; i < lens; i ++ {
		dp[i] = make([]bool, lens)
	}

	for i := 0; i < lens; i ++ {
		if s[i] == '*' {
			dp[i][i] = true
		}
	}

	for i := 1; i < lens; i ++ {
		if (s[i-1] == '(' || s[i-1] == '*') && (s[i] == '*' || s[i] == ')' )  {
			dp[i-1][i] =  true
		}
	}

	for i := lens - 3; i >= 0; i -- {
		f := s[i]
		for j := i + 2; j < lens; j ++ {
			e := s[j]
			if (f == '(' || f == '*') && (e == '*' || e == ')' )  {
				dp[i][j] =  dp[i+1][j-1]
			}
			for k := i; k < j; k ++ {
				if dp[i][j] {
					break
				}
				dp[i][j] = dp[i][k] && dp[k+1][j]
			}
			sort.Slice()
		}
	}


	return dp[0][lens - 1]
}
