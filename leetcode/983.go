package main

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1,4,6,7,8,20}, []int{2,7,15}))
}

// 双百了！
// 组合背包 的最值问题
// dp[i] 是能 i 天的最小值

// 更神奇的官方做法是：递归，从后往前做，dp[i]的语义是能到达最后一天旅游的最小值

var dp[367]int

var spans = []int{1, 7, 30}
func mincostTickets(days []int, costs []int) int {
	dp = [367]int{}
	lend := len(days)

	md := days[lend-1]
	for i := 0; i <= md; i ++ {
		dp[i] = 0x3f3f3f
	}
	for i := 0 ; i < days[0]; i ++ {
		dp[i] = 0
	}

	for j := 0; j < lend; j ++ {
		n := days[j]
		for i := 0; i < 3; i ++ {
			cost := costs[i]
			s := min(n, spans[i])
			// dp[n] = min(dp[n], dp[n-s]+cost)

			if dp[n] > dp[n-s] + cost {
				dp[n] = dp[n-s] + cost
				if j + 1 != lend {
					nx := days[j+1] - 1
					for k := n+1; k <= nx; k ++ {
						dp[k] = min(dp[k-1], dp[k])
					}
				}
			}
		}
	}

	return dp[md]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
