package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	lens := len(s)

	dp := make([]int, lens)

	dp[0] = 1
	for i := 1; i < lens; i ++ {
		dp[i] = dp[i-1]
		if ((s[i] <= '6' && s[i-1] <= '2') || s[i-1] == '1') && s[i] != '0' && s[i-1] != '0' {
			if i > 1 {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-1] + 1
			}
		}
		if s[i] == '0' {
			if s[i-1] > '2' {
				return 0
			}
			if i > 1 {
				dp[i] = dp[i-2]
			} else {
				dp[i] = 1
			}
		}

		if s[i] == '0' && s[i-1] == '0' {
			return 0
		}
	}
	// fmt.Println(dp)

	return dp[lens - 1]
}
