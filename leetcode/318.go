package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

// 4:14
func maxProduct(words []string) int {
	lenw := len(words)
	m := make([]int, lenw)

	for k := range words {
		now := 0
		for _, b := range words[k] {
			now |= 1 << int(b-'a')
		}
		m[k] = now
	}

	res := 0
	for i := 0; i < lenw; i++ {
		for j := 0; j < lenw; j++ {
			if m[i]&m[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
