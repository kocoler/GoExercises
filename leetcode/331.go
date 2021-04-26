package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("#,7,6,9,#,#,#"))
}

func isValidSerialization(preorder string) bool {
	ss := strings.Split(preorder, ",")
	lenss := len(ss)
	num := 1
	for i := 0; i < lenss; i ++ {
		if num <= 0 {
			return false
		}
		if isDigits(ss[i]) {
			num --
			num += 2
		} else {
			num --
		}
		fmt.Println(ss[i], num)
	}
	return num == 0
}

func isDigits(b string) bool {
	if b == "#" {
		return false
	}

	return true
}
