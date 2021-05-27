package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(13))
}

// 丑数
// 注意处理过程 ...
var dp [1700]int
func nthUglyNumber(n int) int {
	// 6/2 5/3 5/5
	// 1 2 3 4 5 6 7 8 9  10
	// 1 2 3 4 5 6 8 9 10 12
	// [0 1 2 3 4 5 6 8 9 10 9 12 15 16 15 18 15 15]
	dp = [1700]int{0, 1}
	p2 := 1; p3 := 1; p5 := 1
	for i := 2; i <= n; i ++ {
		val2 := dp[p2] * 2
		val3 := dp[p3] * 3
		val5 := dp[p5] * 5
		dp[i] = min(min(val2, val3), val5)
		if dp[i] == val2 {
			p2 ++
		}
		if dp[i] == val3 {
			p3 ++
		}
		if dp[i] == val5 {
			p5 ++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
