package main

import (
	"fmt"
)

func main() {
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	bit := 0
	for n := num; n != 0; n -= n & (^ n + 1) {
		bit = n
	}
	fmt.Println(bit)

	return (bit - 1) & (-num)
}
