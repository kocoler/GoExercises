package main

import "fmt"

// Greatest Common Divisor

func main() {
	fmt.Println(gcd(5, 15))
	fmt.Println(gcd(615, 1997))
}

func gcd(a, b int) int {
	c := a % b
	for c != 0 {
		a, b = b, c
		c = a % b
	}

	return b
}

func gcdSub(a, b int) int {
	c := a - b
	for c != 0 {
		a, b = b, c
		c = a - b
	}

	return b
}
