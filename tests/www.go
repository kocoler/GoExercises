package main

import "fmt"

func Q(n int) int {

	if n == 1 {
		return 1
	}
	num := Q(n - 1) + 2 * n - 1
	fmt.Println(n, num)
	return num
}

func main() {
	fmt.Println(Q(3))
}
