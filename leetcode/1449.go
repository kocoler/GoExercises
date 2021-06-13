package main

import (
	"strconv"
)

// 数位成本和为目标值的最大数字
// 完全背包 + 倒推

// var dp [10000]int
func largestNumber(cost []int, target int) string {
	// dp = [10000]int{}
	dp := make([]int, target+1)
	m := make(map[int]int)
	for k, v := range cost {
		m[v] = k + 1
	}
	for i := 0; i <= target; i++ {
		dp[i] = -0x3f3f3f
	}
	dp[0] = 0
	for k, _ := range m {
		for i := k; i <= target; i++ {
			if i >= k {
				dp[i] = max(dp[i-k]+1, dp[i])
			}
		}
	}
	// fmt.Println(dp[target+1])
	if dp[target] <= 0 {
		return "0"
	}

	l := dp[target]
	res := ""
	t := target
	lenc := len(cost)
	for i := l; i > 0; i-- {
		for k := lenc; k > 0; k-- {
			v := cost[k-1]
			if v > t {
				continue
			}
			if dp[t-v] == dp[t]-1 {
				res += strconv.Itoa(k)
				t -= v
				break
			}
		}
		// fmt.Println(i, res)
	}

	return res
}

func max(a, b int) int {
	// fmt.Println(a, b)
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
