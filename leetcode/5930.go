package main

import "fmt"

func main() {
	fmt.Println(maxDistance([]int{1, 1, 1, 6, 1, 1, 1}))
}

func maxDistance(colors []int) int {
	m := make(map[int]int)
	res := 0

	for k := range colors {
		if _, ok := m[colors[k]]; !ok {
			m[colors[k]] = k
		}
		for index, v := range m {
			if index != colors[k] {
				fmt.Println("now", k, index, v, k-v)
				res = max(res, k-v)
			}
		}
	}
	fmt.Println(m)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
