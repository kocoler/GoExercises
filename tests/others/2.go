package main

import "fmt"

func main() {
	a := make([]int, 10)
	for k, _ := range a {
		a[k] = 1
	}
	fmt.Println(a)
}
