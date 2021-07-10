package main

import "fmt"

func main() {
	fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
}

func minFallingPathSum(matrix [][]int) int {
	lenr := len(matrix)
	lenc := len(matrix[0])
	//dp := make([]int, lenc)
	// 1
	res := 0x3f3f3f
	//for i := 0; i < lenc; i ++ {
	//	// res = min(res, matrix[0][i])
	//	dp[i] = matrix[0][i]
	//}

	// 1 -> lenr - 1
	for i := 1; i <= lenr-1; i++ {
		if lenc < 2 {
			//dp[0] = dp[0] + matrix[i][0]
			matrix[i][0] = matrix[i-1][0] + matrix[i][0]
			continue
		}
		//p1 := 0
		// p1 o o  o p1 o
		//p1 = min(dp[0], dp[1]) + matrix[i][0]
		matrix[i][0] = min(matrix[i-1][0], matrix[i-1][1]) + matrix[i][0]
		for j := 1; j < lenc-1; j++ {
			//mi := min(min(dp[j-1], dp[j]), dp[j+1])
			matrix[i][j] = min(min(matrix[i-1][j-1], matrix[i-1][j]), matrix[i-1][j+1]) + matrix[i][j]
			//dp[j-1] = p1
			//p1 = mi + matrix[i][j]
		}

		//dp[lenc-2] = p1
		//dp[lenc-1] = min(dp[lenc-1], dp[lenc-2]) + matrix[i][lenc-1]
		matrix[i][lenc-1] = min(matrix[i-1][lenc-1], matrix[i-1][lenc-2]) + matrix[i][lenc-1]
	}

	for i := 0; i < lenc; i++ {
		//res = min(res, dp[i])
		res = min(res, matrix[lenr-1][i])
	}

	fmt.Println(matrix)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
