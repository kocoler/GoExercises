package main

func matrixBlockSum(mat [][]int, k int) [][]int {
	maxr, maxc := len(mat), len(mat[0])

	dp := make([][]int, maxr+1)
	res := make([][]int, maxr)
	for k := range dp {
		dp[k] = make([]int, maxc+1)
	}
	for k := range res {
		res[k] = make([]int, maxc)
	}

	for r := 0; r < maxr; r++ {
		for c := 0; c < maxc; c++ {
			// (x1, y1) = dp[x1-1][y1] + dp[x1][y1-1] - dp[x1-1][y1-1] + num[x1][y1]
			dp[r+1][c+1] = dp[r+1][c] + dp[r][c+1] - dp[r][c] + mat[r][c]
		}
	}

	for r := 0; r < maxr; r++ {
		for c := 0; c < maxc; c++ {
			// (x2, y2) - (x1, y1) = dp[x2][y2] - dp[x1-1][y2] - dp[x2][y1-1] + dp[x1-1][y1-1]
			x2, y2 := r+k+1, c+k+1
			x1, y1 := r-k+1, c-k+1
			if x1 < 1 {
				x1 = 1
			}
			if y1 < 1 {
				y1 = 1
			}
			if x2 > maxr {
				x2 = maxr
			}
			if y2 > maxc {
				y2 = maxc
			}

			res[r][c] = dp[x2][y2] - dp[x1-1][y2] - dp[x2][y1-1] + dp[x1-1][y1-1]
		}
	}

	return res
}
