package main

import "fmt"

var INF = 0x3f3f3f

func main() {
	//fmt.Println(maxDistance([][]int{{0,0,1,1,1}, {0,1,1,0,0}, {0,0,1,1,0}, {1,0,0,0,0}, {1,1,0,0,1}}))
	fmt.Println(maxDistance([][]int{{0,0,1}, {0,0,0}, {0,0,1}}))
}

// 这题就是双向 dp
// 但是 INF 的作用比较重要
// 因为这个取的是 **最小值**
// 如果取最大值 我认为可以 考虑 -INF
var dp [101][101]int
func maxDistance(grid [][]int) int {
	res := -1
	n := len(grid)
	// c := len(grid[0])
	// dp := make([][]int, r+1)
	// for i := 0; i <= r; i ++ {
	//     dp[i] = make([]int, c+1)
	// }
	dp = [101][101]int{}
	// 0, 0 to r, c
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				dp[i][j] = INF
			}
		}
	}
	for i := 1; i < n; i ++ {
		if grid[0][i] != 1 {
			if grid[0][i-1] == 1 {
				dp[0][i] = 1
			} else {
				dp[0][i] = min(dp[0][i-1] + 1, dp[0][i])
			}
		}
	}
	for i := 1; i < n; i ++ {
		if grid[i][0] != 1 {
			if grid[i-1][0] == 1 {
				dp[i][0] = 1
			} else {
				dp[i][0] = min(dp[i-1][0] + 1, dp[i][0])
			}
		}
	}
	for i:= 1; i < n; i ++ {
		for j := 1; j < n; j ++ {
			if grid[i][j] != 1 {
				if grid[i-1][j] == 1 || grid[i][j-1] == 1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
				}
			}
		}
	}

	res = max(res, dp[n-1][n-1])
	for i := n-2; i >= 0; i -- {
		if grid[n-1][i] != 1 {
			if grid[n-1][i+1] == 1 {
				dp[n-1][i] = 1
			} else {
				dp[n-1][i] = min(dp[n-1][i+1]+1, dp[n-1][i])
			}
			res = max(res, dp[n-1][i])
		}
	}
	for i := n-2; i >= 0; i -- {
		if grid[i][n-1] != 1 {
			if grid[i+1][n-1] == 1 {
				dp[i][n-1] = 1
			} else {
				dp[i][n-1] = min(dp[i+1][n-1] + 1, dp[i][n-1])
			}
			res = max(res, dp[i][n-1])
		}
	}

	for i:= n-2; i >= 0; i -- {
		for j := n-2; j >= 0; j -- {
			if grid[i][j] != 1 {
				if grid[i+1][j] == 1 || grid[i][j+1] == 1 {
					dp[i][j] = 1
				} else {
					if dp[i+1][j] > 0 && dp[i][j+1] > 0 {
						dp[i][j] = min(min(dp[i+1][j], dp[i][j+1]) + 1, dp[i][j])
					}
				}
				res = max(res, dp[i][j])
			}
		}
	}
	res = max(res, dp[0][0])

	if res >= INF || res == 0 {
		return -1
	}

	return res
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
