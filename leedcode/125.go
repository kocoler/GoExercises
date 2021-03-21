package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("1p2"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	lens := len(s)
	left := 0
	right := lens - 1
	for left < right {
		for !isAlpha(s[left]) && left < right {
			left ++
		}
		for !isAlpha(s[right]) && left < right {
			right --
		}

		if s[left] != s[right] {
			return false
		}

		left ++
		right --
	}

	return true
}

func isAlpha(b byte) bool {
	if 'a' <= b && b <= 'z' {
		return true
	}
	if '0' <= b && b <= '9' {
		return true
	}

	return false
}
