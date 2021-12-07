package main

import "fmt"

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
}

var diff [1002]int

func carPooling(trips [][]int, capacity int) bool {
	final := 0

	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
		final = max(final, trip[2])
	}
	fmt.Println(diff[:final+2])
	now := 0
	for i := 0; i <= final; i++ {
		now += diff[i]
		if now > capacity {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
