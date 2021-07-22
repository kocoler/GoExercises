package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	res := 0
	for x != 0 && y != 0 {
		if x % 2 != y % 2 {
			res ++
		}
		x /= 2
		y /= 2
	}
	fmt.Println(res)
	for x != 0 {
		if x % 2 != 0 {
			res ++
		}
		x /= 2
	}
	fmt.Println(x, y, res)
	if y != 0 {
		if y % 2 != 0 {
			res ++
		}
		y /= 2
	}
	fmt.Println(x, y, res)

	return res
}
