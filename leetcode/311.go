package main

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	res := make([][]int, len(mat1))
	for k := range res {
		res[k] = make([]int, len(mat2[0]))
	}

	cal := func(r, c int) int {
		res := 0

		for k := range mat1[r] {
			res += mat1[r][k] * mat2[k][c]
		}

		return res
	}

	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat2[0]); j++ {
			res[i][j] = cal(i, j)
		}
	}

	return res
}
