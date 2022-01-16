package main

// 当然啦，也可以直接原地做除法运算
// rm[9][9]int{}
// cm[9][9]int{}
// sub[3][3][9]int{}
// sub[i/3][j/3]

func isValidSudoku(board [][]byte) bool {
	rm := [9][9]int{}
	cm := [9][9]int{}

	for si := 0; si < 9; si += 3 {
		for sj := 0; sj < 9; sj += 3 {
			nm := [9]int{}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					r, c := si+i, sj+j
					if board[r][c] == '.' {
						continue
					}
					b := board[r][c] - '0' - 1

					rm[r][b]++
					cm[c][b]++
					nm[b]++

					if rm[r][b] > 1 || cm[c][b] > 1 || nm[b] > 1 {
						return false
					}
				}
			}
		}
	}

	return true
}
