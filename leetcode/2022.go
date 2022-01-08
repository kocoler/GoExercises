package main

func construct2DArray(original []int, m int, n int) [][]int {
	leno := len(original)
	var res [][]int
	if leno != m*n {
		return res
	}
	res = make([][]int, m)
	for k := range res {
		res[k] = make([]int, n)
	}

	for k, v := range original {
		r, c := k/n, k%n
		res[r][c] = v
	}

	return res
}
