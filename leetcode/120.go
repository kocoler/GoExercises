package main

// 可以倒推 使用滚动数组
// 不需要 dp 数组
// 直接返回 triangle[0][0]

// var dp [201] int
func minimumTotal(triangle [][]int) int {
	maxNum := 0x3f3f3f
	lent := len(triangle)
	dp := make([]int, lent)
	lenm := len(triangle[lent-1])
	dp[0] = triangle[0][0]
	for i := 1; i < lent; i ++ {
		lenn := len(triangle[i])
		dp[lenn-1] = dp[lenn-2] + triangle[i][lenn-1]
		for j := lenn-2; j > 0; j -- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] = dp[0] + triangle[i][0]
	}
	// 2
	// 3 4
	// 6 5 7
	// 4 1 8 3
	res := maxNum
	for i := 0; i < lenm; i ++ {
		res = min(res, dp[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// double space
// var dp [201][201] int
// func minimumTotal(triangle [][]int) int {
//     maxNum := 0x3f3f3f
//     lent := len(triangle)
//     lenm := len(triangle[lent-1])
//     dp[0][0] = triangle[0][0]
//     for i := 1; i < lent; i ++ {
//         lenn := len(triangle[i])
//         dp[i][0] = dp[i-1][0] + triangle[i][0]
//         for j := 1; j < lenn-1; j ++ {
//             dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
//         }
//         dp[i][lenn-1] = dp[i-1][lenn-2] + triangle[i][lenn-1]
//     }

//     res := maxNum
//     for i := 0; i < lenm; i ++ {
//         // fmt.Println(dp[lent-1][i])
//         res = min(res, dp[lent-1][i])
//     }

//     return res
// }
