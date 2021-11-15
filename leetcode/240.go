package main

import "sort"

// 二分
func searchMatrix(matrix [][]int, target int) bool {
	lenr := len(matrix)
	lenc := len(matrix[0])

	if matrix[0][0] > target || matrix[lenr-1][lenc-1] < target {
		return false
	}

	// find rl
	indexrl := sort.Search(lenr, func(i int) bool {
		if matrix[i][0] > target {
			return true
		}
		return false
	})

	// find rr
	indexrr := sort.Search(lenr, func(i int) bool {
		if matrix[i][lenc - 1] >= target {
			return true
		}
		return false
	})

	indexrl --

	for j := indexrl; j >= indexrr; j ++ {
		r := j
		// find c
		indexc := sort.Search(lenc, func(i int) bool {
			if matrix[r][i] >= target {
				return true
			}
			return false
		})

		if indexc < lenc && matrix[r][indexc] == target {
			return true
		}
	}

	return false
}

// Z型搜索
func searchMatrix(matrix [][]int, target int) bool {
	lenr := len(matrix)
	lenc := len(matrix[0])

	if matrix[0][0] > target || matrix[lenr-1][lenc-1] < target {
		return false
	}

	x, y := 0, lenc - 1
	for x <= lenr - 1 && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x ++
		} else {
			y --
		}
	}

	return false
}
