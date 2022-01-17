package main

// 可以加上位运算压缩状态

func solveSudoku(board [][]byte) {
	rm := [9][9]bool{}
	cm := [9][9]bool{}
	subset := [3][3][9]bool{}
	var bm [][2]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				bm = append(bm, [2]int{i, j})
			} else {
				rm[i][board[i][j]-'1'] = true
				cm[j][board[i][j]-'1'] = true
				subset[i/3][j/3][board[i][j]-'1'] = true
			}
		}
	}

	found := false
	var dfs func(int)
	dfs = func(pos int) {
		if pos == len(bm) {
			found = true
			return
		}
		r, c := bm[pos][0], bm[pos][1]
		for i := 0; i < 9; i++ {
			if !rm[r][i] && !cm[c][i] && !subset[r/3][c/3][i] {
				board[r][c] = byte('1' + i)
				rm[r][i] = true
				cm[c][i] = true
				subset[r/3][c/3][i] = true
				dfs(pos + 1)
				if found {
					return
				}
				rm[r][i] = false
				cm[c][i] = false
				subset[r/3][c/3][i] = false
			}
		}
	}

	dfs(0)
}
