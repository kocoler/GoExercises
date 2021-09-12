package main

import (
	"fmt"
	"sort"
)

// 康托展开
// X = a[n]*(n-1)! + a[n-1]*(n-2)! + ... + a[i]*(i-1)! + ... + a[1]*0!

var FAC = []int{1,1,2,6,24,120,720,5040,40320,362800}

// 34152
// 2 * 4! + 2 * 3! + 0 * 2! + 1 * 1! + 0 * 0!
func cantor(a string) int {
	res := 0
	lena := len(a)

	for i := 0; i < lena; i ++ {
		smaller := 0
		for j := i + 1; j < lena; j ++ {
			if a[j] < a [i] {
				smaller ++
			}
		}
		res = res + FAC[lena-i-1] * smaller
	}

	return res
}

// 61, lena = 5
// 61 / 4! = 2 ··· 13  => 2
// 13 / 3! = 2 ··· 1  => 2
// 1 / 2! = 0 ··· 1 => 0
// 1 / 1! = 1 ··· 0 => 1
func decantor(a []int, n int) []int {
	sort.Ints(a)
	lena := len(a)
	res := make([]int, lena)

	for i := lena - 1; i >= 0; i -- {
		index := n / FAC[i]
		n = n % FAC[i]
		res[lena - 1 - i] = a[index]
		a = append(a[:index], a[index+1:] ...)
	}

	return res
}

func main() {
	fmt.Println(cantor("34152"))

	//fmt.Println(decantor([]int{1,2,3,4,5}, 61))
}
