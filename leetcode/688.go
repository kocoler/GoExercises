package main

import "math"

var dis = [8][2]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, n)
	}

	dp[row][column] = 1

	dfs := func() {
		newdp := make([][]float64, n)
		for i := 0; i < n; i++ {
			newdp[i] = make([]float64, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] != 0 {
					for k := 0; k < 8; k++ {
						nxr := i + dis[k][0]
						nxc := j + dis[k][1]
						if nxr < 0 || nxr >= n || nxc < 0 || nxc >= n {
							// res = (res + dp[i][j]) % mod
							// res += dp[i][j]
						} else {
							// 这里！每次都先除 8 计算出这一步的概率
							newdp[nxr][nxc] += dp[i][j] / 8
						}
					}
				}
			}
		}
		dp = newdp
	}

	for i := 0; i < k; i++ {
		dfs()
	}

	var res float64
	for i := 0; i < n; i ++ {
		for j := 0; j < n; j ++ {
			if dp[i][j] != 0 {
				res += dp[i][j]
			}
		}
	}

	return res
}