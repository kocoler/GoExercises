package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAA", 0))
}

func characterReplacement(s string, k int) int {
	lens := len(s)

	left := 0
	right := 0
	res := 0
	var record [26]int
	maxC := 0
	for i := 0; i < lens; i++ {
		w := s[i] - 'A'
		record[w]++
		if maxC < record[w] {
			maxC = record[w]
		}
		if maxC + k > res {
			right ++
			res ++
		} else {
			record[s[left]-'A']--
			left ++
		}
		fmt.Println(left, right)
	}

	return res
}