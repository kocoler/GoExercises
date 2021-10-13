package main

import (
	"fmt"
	"math"
)

// 就，找规律吧 ...

func main() {
	fmt.Println(countOrders(3))
}

var mod = int(math.Pow(10, 9) + 7)

func countOrders(n int) int {
	if n == 1 {
		return 1
	}

	res := 1
	for i := 2; i <= n; i ++ {
		temp := 2 * (i - 1) + 1 // 2(n - 1) + 1
		temp = temp * (temp - 1) / 2 + temp // 2(n - 1) + 1 * 2(n - 1)
		res = res * temp % mod
	}

	return res
}
