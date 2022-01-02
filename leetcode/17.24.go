package main

import (
	"fmt"
	"math"
)

// 最大子序和 变种

func main() {
	fmt.Println(getMaxMatrix([][]int{{9, -8, 1, 3, -2}, {-3, 7, 6, -2, 4}, {6, -4, -4, 8, -7}}))
}

func getMaxMatrix(matrix [][]int) []int {
	r := len(matrix)
	c := len(matrix[0])

	maxNum := math.MinInt64
	res := [4]int{}
	for i := 0; i < r; i++ {
		sum := make([]int, c+1)
		for j := i; j < r; j++ {
			for k := 0; k < c; k++ {
				sum[k+1] += matrix[j][k]
			}

			// 最大子序和 dp
			startPos := 1
			num := 0
			for v := 1; v <= c; v++ {
				if num < 0 {
					startPos = v
					num = sum[v]
				} else {
					num = num + sum[v]
				}

				if num > maxNum {
					maxNum = num
					res[0] = i
					res[1] = startPos - 1
					res[2] = j
					res[3] = v - 1
				}
			}
		}
	}

	return res[:]
}
