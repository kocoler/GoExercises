package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPaths(1,3,3,0,1))
}

var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
const mod int = 1e9 + 7

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	res := 0

	dp[startRow][startColumn] = 1

	// 返回增量
	dfs := func() {
		newdp := make([][]int, m)
		for i := 0; i < m; i++ {
			newdp[i] = make([]int, n)
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] != 0 {
					//fmt.Println("found!")
					for k := 0; k < 4; k++ {
						nxr := i + dis[k][0]
						nxc := j + dis[k][1]
						if nxr < 0 || nxr >= m || nxc < 0 || nxc >= n {
							res = (res + dp[i][j]) % mod
							//res %= m
						} else {
							newdp[nxr][nxc] = (newdp[nxr][nxc] + dp[i][j]) % mod
						}
					}
				}
			}
		}
		dp = newdp
	}

	for i := 0; i < maxMove; i++ {
		dfs()
	}

	return res
}
