package main

import "fmt"

//输入: A = [5,4,0,3,1,6,2]
//输出: 4
//解释:
//A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.
//
//其中一种最长的 S[K]:
//S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}
// 20000 * 20000
// dp[a][b]

func main() {
	fmt.Println(arrayNesting([]int{5,4,0,3,1,6,2}))
}

// 顺序的时候可以用 []bool 而不是 map
func arrayNesting(nums []int) int {
	lenn := len(nums)
	res := 1
	m := make([]bool, lenn)

	for k := range nums {
		if !m[k] {
			step := 1
			index := nums[k]
			for index != k {
				m[index] = true
				index = nums[index]
				step ++
			}

			if step > res {
				res = step
			}
		}
	}

	return res
}
