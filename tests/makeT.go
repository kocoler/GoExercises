package main

import "fmt"

func main() {
	var a = make([]int, 1)
	fmt.Println(a, " ", len(a))
	a = append(a, 1)
	fmt.Println(a, " ", len(a))
	
}
