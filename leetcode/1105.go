package main

// 第二类基本 DP 问题
// 枚举前一个状态

func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books)+1)
	for i := 1; i <= len(books); i++ {
		maxHeight := 0
		width := shelfWidth
		dp[i] = dp[i-1] + books[i-1][1]
		for j := i; j >= 1; j-- {
			width -= books[j-1][0]
			if width < 0 {
				break
			}
			maxHeight = max(maxHeight, books[j-1][1])
			dp[i] = min(dp[i], maxHeight+dp[j-1])
		}
	}

	return dp[len(books)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
