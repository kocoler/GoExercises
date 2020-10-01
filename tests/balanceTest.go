package main

import "fmt"

func main() {
	balance := 123456
	//56 34 12
	fmt.Println(balance % 100)
	fmt.Println((balance % 10000) / 100)
	fmt.Println(balance / 10000)
}
