package main

// 带备忘录的搜索
// 考虑几个 )) () 的情况
// 其他做法: stack 记录弹出个数

func longestValidParentheses(s string) int {
	dp := make([]int, len(s) + 1)
	lens := len(s)

	res := 0
	for i := 2; i <= lens; i ++ {
		index := i-1
		if s[index] == ')' {
			if s[index-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				p := index-dp[i-1]-1
				if  p >= 0 && s[p] == '(' {
					dp[i] = dp[i-1] + 2 + dp[p]
				}
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
