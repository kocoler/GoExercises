package main

import "fmt"

func returnT(a []int, res []int) []int {
	if len(a) != 0 {
		res = append(res, a[len(a)-1])
		return returnT(a[:len(a)-1], res)
	}

	return res
}

func returnE(a []int, res *[]int) {
	if len(a) != 0 {
		*res = append(*res, a[len(a)-1])
		returnE(a[:len(a)-1], res)
	}
}

func main() {
	res := []int{}
	returnE([]int{1, 2, 3, 4}, &res)
	fmt.Println(res)
}
