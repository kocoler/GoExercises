package main

import "fmt"

func main() {
	fmt.Println(fastPower(2, 3, 10000))
}

// fastPower calculates x^n
func fastPower(x, n int64, mod int64) int64 {
	var res int64 = 1

	for n > 0 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}

	return res
}
