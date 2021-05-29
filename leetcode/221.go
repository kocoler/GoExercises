package main

// 最大正方形
// dp 维护当前的就可以了
// ！！注意 min 的取值！！
var dp [301][301]int
func maximalSquare(matrix [][]byte) int {
	lr := len(matrix)
	lc := len(matrix[0])
	res := 0
	dp = [301][301]int{}

	for i := 1; i <= lr; i ++ {
		for j := 1; j <= lc; j ++ {
			if matrix[i-1][j-1] == '1' {
				if dp[i-1][j] == 0 || dp[i][j-1] == 0 || dp[i-1][j-1] == 0 {
					dp[i][j] = 1
					res = max(res, 1)
					continue
				}
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				res = max(res, dp[i][j] * dp[i][j])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}

		return c
	}
	// a > b
	if b < c {
		return b
	}

	return c
}

