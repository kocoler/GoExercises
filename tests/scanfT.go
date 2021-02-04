package main

import "fmt"

func main() {
	s := ""
	var a, b byte

	n, err := fmt.Scanf("%s", &s)
	fmt.Println(n, err)
	fmt.Println()
	n, err = fmt.Scanf("%d", &a)
	fmt.Println(n, err)
	fmt.Println()
	n, err = fmt.Scanf("%d", &b)
	fmt.Println(n, err)
	fmt.Println()

	fmt.Println(s, "a: ", string(a), "b: ", string(b))
}
