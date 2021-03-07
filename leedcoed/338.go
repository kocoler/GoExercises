package main

import "fmt"

func main() {
	fmt.Println(countBits(4))
}

func countBits(num int) []int {
	dp := make([]int, num+2)
	pre := 1
	b := 3
	dp[1]=1; dp[2]=1
	for i := 3; i <= num; i ++ {
		fmt.Println(i, b, pre)
		if b >= i {
			dp[i] = dp[pre] + dp[i-pre]
		} else if b < i {
			dp[i] = 1
			pre = b
			b *= 2
			b += 1
			fmt.Println(b)
		}
	}

	return dp[:num+1]
}
