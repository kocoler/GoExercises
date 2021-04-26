package main

func generateMatrix(n int) [][]int {
	// lenc := len(matrix) // 竖
	// lenm := len(matrix[0]) // 横
	lenc := n
	lenm := n
	num := 0
	max := n * n
	var res [][]int
	for i := 0; i < n; i ++ {
		temp := make([]int, n)
		res = append(res, temp)
	}

	// right 0
	// down 1
	// left 2
	// up 3
	dir := 0
	i := 0; j := 0
	temp := 0
	lenc --
	for num != max {
		// res = append(res, matrix[i][j])
		num ++
		res[i][j] = num
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
