package main

func setZeroes(matrix [][]int)  {
	lenn := len(matrix) // 列
	lenm := len(matrix[0]) // 行
	// matrix[0][0] 只负责竖排 => 第二回合
	ori := false
	for i := 0; i < lenm; i ++ {
		if matrix[0][i] == 0 {
			ori = true
			break
		}
	}

	for i := 1; i < lenn; i ++ {
		for j := 0; j < lenm; j ++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < lenn; i ++ {
		if matrix[i][0] == 0 {
			for j := 0; j < lenm; j ++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < lenm; i ++ {
		if matrix[0][i] == 0 {
			for j := 0; j <lenn; j ++ {
				matrix[j][i] = 0
			}
		}
	}

	if ori {
		// for i := 0; i < lenn; i ++ {
		//     matrix[i][0] = 0
		// }
		for i := 0; i < lenm; i ++ {
			matrix[0][i] = 0
		}
	}

	return
}
