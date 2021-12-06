package main

import "fmt"

func main() {
	fmt.Println(shortestCommonSupersequence("abacd", "cab"))
}

// 太牛了， SCS 问题
// 这个初始化也十分神奇
// 当然，也可以直接 LCS 问题，再推导出最终结果，因为 LCS 一定是其中的子序列，只要不断把中间值加进来就好

func shortestCommonSupersequence(str1 string, str2 string) string {
	dp := make([][]int, len(str1)+1)
	for k := range dp {
		dp[k] = make([]int, len(str2)+1)
	}

	for k := range dp {
		dp[k][0] = k
	}
	for k := range dp[0] {
		dp[0][k] = k
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	//fmt.Println(dp)
	var str []byte
	i, j := len(str1), len(str2)
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			str = append(str, str1[i-1])
			i--
			j--
		} else if dp[i][j] == dp[i-1][j]+1 {
			str = append(str, str1[i-1])
			i--
		} else {
			str = append(str, str2[j-1])
			j--
		}
	}
	for i > 0 {
		str = append(str, str1[i-1])
		i--
	}
	for j > 0 {
		str = append(str, str2[j-1])
		j--
	}
	reverse(str)
	return string(str)
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

func reverse(str []byte) {
	for i, j := 0, len(str)-1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}
