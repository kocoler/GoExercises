package main

import "fmt"

func main() {
	fmt.Println(isValid("aabcbc"))
}

func isValid(s string) bool {
	// 1: a 2: b 3: c
	var stack []int
	i := 0
	for i < len(s) {
		if i < len(s) {
			switch s[i] {
			case 'a':
				stack = append(stack, 1)
			case 'b':
				stack = append(stack, 2)
			case 'c':
				if len(stack) < 2 || stack[len(stack)-1] != 2 || stack[len(stack)-2] != 1 {
					return false
				} else {
					stack = stack[:len(stack)-2]
				}
			}
			i++
		}
	}

	return len(stack) == 0
}
