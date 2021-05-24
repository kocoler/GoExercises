package main

import "fmt"

func reSolve(s, e int) int {
	if s >= e {
		return 0
	}
	mid := (e - s) / 2 + s
	fmt.Println(mid)
	return reSolve(mid+1, e) + mid
}

func getMoneyAmount(n int) int {
	return reSolve(1, n)
}

// 1 2 3 4 5 6 7 8 9 10
func main() {
	fmt.Println(getMoneyAmount(10))
}
