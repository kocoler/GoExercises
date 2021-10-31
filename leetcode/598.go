package main

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) <= 0 {
		return m*n
	}

	x := ops[0][0]
	y := ops[0][1]
	for k := range ops {
		x = min(x, ops[k][0])
		y = min(y, ops[k][1])
	}

	return x * y
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
