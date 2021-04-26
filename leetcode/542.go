package main

// dp row -> colume
// var matrixW [][]int
var r int
var c int

func updateMatrix(matrix [][]int) [][]int {
	max := 0x3f3f3f
	var dp [][]int
	// matrixW = matrix
	dp = make([][]int, len(matrix))
	r = len(matrix)
	c = len(matrix[0])
	for i := 0; i < r; i++ {
		dp[i] = make([]int, c)
	}
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if matrix[j][i] == 0 {
				dp[j][i] = 0
				continue
			} else {
				dp[j][i] = max
			}
			if i > 0 {
				dp[j][i] = min(dp[j][i-1]+1, dp[j][i])
			}
			if j > 0 {
				dp[j][i] = min(dp[j-1][i]+1, dp[j][i])
			}
		}
	}
	for i := c - 1; i >= 0; i-- {
		for j := r - 1; j >= 0; j-- {
			if i < c-1 {
				dp[j][i] = min(dp[j][i+1]+1, dp[j][i])
			}
			if j < r-1 {
				dp[j][i] = min(dp[j+1][i]+1, dp[j][i])
			}
		}
	}

	return dp
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// func findMinDis(r, c int) {
//     var up, down, left, right int
//     if r-1 > 0 {
//         up = matrix[r-1][c]
//     }
//     if r+1

//     down := matrix[r+1][c]
//     left := matrix[r][c-1]
//     right := matrix[r][c+1]
//     if ()
// }

// func min(a, b) {

// }
