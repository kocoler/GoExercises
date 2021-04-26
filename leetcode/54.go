package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

// var record [10001][10001]bool
func spiralOrder(matrix [][]int) []int {
	lenc := len(matrix) // 竖
	lenm := len(matrix[0]) // 横
	// record = [10001][10001]bool{}
	num := 0
	max := lenm * lenc
	var res []int

	// right 0
	// down 1
	// left 2
	// up 3
	dir := 0
	i := 0; j := 0
	temp := 0
	lenc --
	for num != max {
		fmt.Println(i, j, dir, temp)
		res = append(res, matrix[i][j])
		num ++
		temp ++

		if dir == 0 || dir == 2 {
			if temp >= lenm {
				lenm --
				dir ++
				temp = 0
			}
		} else {
			if temp >= lenc {
				lenc --
				if dir == 1 {
					dir ++
				} else {
					dir = 0
				}
				temp = 0
			}
		}

		switch dir {
		case 0:
			j ++
		case 1:
			i ++
		case 2:
			j --
		case 3:
			i --
		}
	}

	return res
}
