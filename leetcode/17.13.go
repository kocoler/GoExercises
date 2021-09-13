package main

// 动规条件是不可分割的！最小！
// 更优的做法是 trim 树优化，但是我不会

func respace(dictionary []string, sentence string) int {
	m := make(map[string]bool)
	lens := len(sentence)
	sb := []byte(sentence)

	for _, v := range dictionary {
		m[v] = true
	}

	dp := make([]int, lens+1)

	dp[0] = 0

	for i := 1; i <= lens; i ++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j ++ {
			if m[string(sb[j:i])] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}

	return dp[lens]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
