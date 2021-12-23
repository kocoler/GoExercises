package main

import "fmt"

// 两个数之和可以被 K 整除
// 相当于
// 两个数对 K 取模，相加可以被 K 整除

func main() {
	fmt.Println(canArrange([]int{}, 2))
}

func canArrange(arr []int, k int) bool {
	m := make(map[int]int)

	for _, v := range arr {
		m[((v%k)+k)%k]++
	}

	for _, v := range arr {
		mo := ((v % k) + k) % k
		fmt.Println(mo, v, m)
		if mo != 0 {
			mo = k - mo
		}
		if m[mo] <= 0 {
			return false
		}
		m[mo]--
	}
	return true
}
