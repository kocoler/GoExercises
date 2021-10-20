package main

import "fmt"

func main() {
	var i int
	i = 0
	println(&i)
	b := []byte{1, 2, 3, 4, 6, 67, 7, 7, 7, 7, 7, 7, 7, 7, 10}
	fmt.Println(b)
	println(&b, &b[0])
}
