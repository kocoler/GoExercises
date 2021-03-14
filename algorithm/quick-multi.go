package main

import "fmt"

func main() {
	fmt.Println(quickMulti(10, 20))
}

func quickMulti(A, B int) int {
	ans := 0
	for ; B > 0; B >>= 1 {
		if B & 1 > 0 {
			ans += A
		}
		A <<= 1
	}
	return ans
}
