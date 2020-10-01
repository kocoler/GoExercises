package main

import (
	"fmt"
)

var x = 'A'
var y = 'B'
var z = 'C'
var stp int = 0

func move(n int, a, b rune) {
	stp++
	fmt.Printf("第%d步：把第%d個盤子從%c到%c\n", stp, n, a, b)
}

func hanoi(n int, x, y, z rune) {
	if n == 0 {
		return
	} else {
		hanoi(n-1, x, z, y)
		move(n, x, z)
		hanoi(n-1, y, x, z)
	}
}

func main() {
	var n int
	fmt.Printf("輸入個數：")
	fmt.Scanf("%d", &n)
	hanoi(n, x, y, z)
}
