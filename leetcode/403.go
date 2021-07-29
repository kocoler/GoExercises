package main

var dp[2001][2001]bool
func canCross(stones []int) bool {
	dp = [2001][2001]bool{}
	lens := len(stones)
	// last := stones[lens-1]
	dp[0][0] = true

	for i := 1; i < lens; i ++ {
		// 这里的 j 是上一个跳跃的起点 ！！
		// 这个转化很神奇
		for j := i-1; j >= 0; j -- {
			// 这里的 k 就可以转化为两点间距离了
			k := stones[i] - stones[j]
			if k > j + 1 {
				break
			}
			dp[i][k] = dp[j][k] || dp[j][k-1] || dp[j][k+1]
			if i == lens-1 && dp[i][k] {
				return true
			}
		}
	}

	return false
}
