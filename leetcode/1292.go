package main

import "sort"

func maxSideLength(mat [][]int, threshold int) int {
	maxr, maxc := len(mat), len(mat[0])
	dp := make([][]int, maxr+1)
	for k := range dp {
		dp[k] = make([]int, maxc+1)
	}

	for i := 0; i < maxr; i++ {
		for j := 0; j < maxc; j++ {
			// dp[x, y] = dp[x-1][y] + dp[x][y-1] - dp[x-1][y-1] + mat[x][y]
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + mat[i][j]
		}
	}

	res := sort.Search(min(maxc, maxr), func(length int) bool {
		for i := 0; i+length < maxr; i++ {
			for j := 0; j+length < maxc; j++ {
				// (x2, y2) - (x1, y1) = dp[x2][y2] - dp[x2][y1-1] - dp[x1-1][y2] + dp[x1-1][y1-1]
				x2, y2 := i+length+1, j+length+1
				x1, y1 := i+1, j+1
				x1 = max(1, x1)
				y1 = max(1, y1)
				x2 = min(maxr, x2)
				y2 = min(maxc, y2)

				sum := dp[x2][y2] - dp[x2][y1-1] - dp[x1-1][y2] + dp[x1-1][y1-1]
				if sum <= threshold {
					return false
				}
			}
		}
		return true
	})

	return res - 1
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
