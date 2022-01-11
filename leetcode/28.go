package main

import "fmt"

// kmp 模版题呀，多写写

func main() {
	fmt.Println(strStr("aabaaabaaac", "ghiabcdefhelloadamhelloabcdefghi"))
}

func strStr(haystack string, needle string) int {
	return kmp(haystack, needle)
}

func kmp(str, pattern string) int {
	lens, lenp := len(str), len(pattern)

	// get next array
	next := make([]int, lenp+1)
	next[0] = -1
	for i, j := 0, -1; i < lenp; {
		if j == -1 || pattern[i] == pattern[j] {
			j++
			i++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	fmt.Println(next)

	i, j := 0, 0
	for i < lens && j < lenp {
		if j == -1 || str[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == lenp {
		return i - j
	}

	return -1
}
