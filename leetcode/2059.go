package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{70,83,-93,47,-81,94,64,84,4,28,37,99,42,74}, 95, -25))
}

func minimumOperations(nums []int, start int, goal int) int {
	if start == goal {
		return 0
	}
	// bfs
	queue := []int{start}
	// m := make(map[int]struct{})
	// 可以用数组充当 map 性能会好一点
	m := make([]int, 1001)
	for k := range m {
		m[k] = -1
	}
	m[start] = 0

	res := 0
	for len(queue) > 0 {
		res ++
		lenq := len(queue)

		for i := 0; i < lenq; i ++ {
			for k := range nums {
				numadd := nums[k] + queue[i]
				numsub := queue[i] - nums[k]
				numxor := nums[k] ^ queue[i]
				if numadd == goal || numsub == goal || numxor == goal {
					return res
				}
				if numadd <= 1000 && numadd >= 0 && m[numadd] == -1 {
					m[numadd] = res
					queue = append(queue, numadd)
				}
				if numsub <= 1000 && numsub >= 0 && m[numsub] == -1 {
					m[numsub] = res
					queue = append(queue, numsub)
				}
				if numxor <= 1000 && numxor >= 0 && m[numxor] == -1{
					m[numxor] = res
					queue = append(queue, numxor)
				}
			}
		}

		queue = queue[lenq:]
	}

	return -1
}
