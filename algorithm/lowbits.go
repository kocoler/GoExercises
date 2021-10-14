package main

import "fmt"

func main() {
	fmt.Println(lowbit(8))
}

func lowbit(n int) int {
	// n & (-n)
	return n & (^ n + 1)
}
