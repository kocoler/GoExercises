package main

import "fmt"

func main() {
	str := "abcd"
	b := []byte(str)

	for k, v := range b {
		v = 'a'
		b[k] = 'a'
		fmt.Println(k, string(v))
	}

	fmt.Println(string(b))
}

