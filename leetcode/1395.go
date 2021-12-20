package main

func numTeams(rating []int) int {
	// 定长上升子序列 + 定长下降子序列
	dp := make([][2]int, len(rating))
	res := 0

	for i := 1; i < len(rating); i++ {
		// dp[i] 是这个位置比这个数大的数量
		cnt0 := 0
		cnt1 := 0
		for j := i - 1; j >= 0; j-- {
			if rating[i] > rating[j] {
				cnt0++
				if dp[j][0] > 0 {
					res += dp[j][0]
				}
			} else {
				cnt1++
				if dp[j][1] > 0 {
					res += dp[j][1]
				}
			}
		}

		dp[i][0] = cnt0
		dp[i][1] = cnt1
	}

	return res
}
