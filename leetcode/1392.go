package main

import "fmt"

func main() {
	fmt.Println(longestPrefix("aab"))
}

func longestPrefix(s string) string {
	// next
	next := make([]int, len(s)+1)

	next[0] = -1
	for i, j := 0, -1; i < len(s); {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	fmt.Println(next)

	return s[:next[len(s)]]
}
