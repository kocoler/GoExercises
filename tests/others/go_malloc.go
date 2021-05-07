package main

import "fmt"

func main() {
	A := al(1)
	al(2)
	fmt.Println(A)
}

func al(num int) *int {
	a := 0
	return &a
}
