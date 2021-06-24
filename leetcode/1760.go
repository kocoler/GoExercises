package main

import "fmt"

// 袋子里最少数目的球

// 二分寻找可以的值

func main() {
	fmt.Println(minimumSize([]int{9}, 2))
}

func minimumSize(nums []int, maxOperations int) int {
	l := 1
	right := 1000000000

	mid := (l + right) / 2
	res := mid

	for l <= right {
		if legal(nums, maxOperations, mid) {
			right = mid - 1
			res = mid
		} else {
			l = mid + 1
		}
		mid = (l+right)/2
	}

	return res
}

func legal(nums []int, maxOperations int, min int) bool {
	for _, v := range nums {
		if v % min == 0{
			maxOperations -= v / min - 1
		} else {
			maxOperations -= v/min
		}

	}

	return maxOperations >= 0
}

