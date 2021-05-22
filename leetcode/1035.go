package main

// 1035. 不相交的线
// 唉，竟然第一想法不是最长子字符串，太菜了
var dp [501][501]int
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	for i := 0; i < len1; i ++ {
		for j := 0; j < len2; j ++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[len1][len2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

