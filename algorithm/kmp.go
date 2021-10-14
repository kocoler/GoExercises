package main

import (
	"fmt"
)

func getNext(str string) []int {
	lens := len(str)
	next := make([]int, lens+1)

	i := 0; j := -1
	next[0] = -1

	for i < lens {
		if j == -1 || str[i] == str[j] {
			i ++
			j ++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}

func kmp(ori, p string, next []int) int {
	lens := len(ori)
	lenp := len(p)
	i := 0; j := 0

	for i < lens && j < lenp {
		if j == -1 || ori[i] == p[j] {
			i ++
			j ++
		} else {
			j = next[j]
		}
	}

	if j == lenp {
		return i - j + 1
	}

	return -1
}


func main() {
	ori := "ababababca"
	p := "abababca"
	next := getNext(p)
	fmt.Println(kmp(ori, p, next))
}
