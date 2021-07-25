package main

import "math"

// 多重 01 背包 的 组合问题
// 三重循环
// 有溢出问题，需要每次都取模

var dp[31][1001]int
var mod = math.Pow(10, 9) + 7
var m = int(mod)
func numRollsToTarget(d int, f int, target int) int {
	dp = [31][1001]int{}

	dp[0][0] = 1
	for i := 1; i <= d; i ++ {
		for j := 1; j <= f; j ++ {
			for k := j; k <= target; k ++ {
				dp[i][k] += dp[i-1][k-j]
				dp[i][k] %= m
			}
		}
	}

	return dp[d][target]
}
