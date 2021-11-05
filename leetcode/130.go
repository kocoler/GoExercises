package main

import "fmt"

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}

// dfs
// 可以反向入手，就是说边界入手

var direction = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
var lenr int
var lenc int
func solve(board [][]byte) {
	lenr = len(board)
	lenc = len(board[0])

	for i := 0; i < lenr; i ++ {
		dfs(&board, i, 0)
		dfs(&board, i, lenc - 1)
	}
	for i := 0; i < lenc; i ++ {
		dfs(&board, 0, i)
		dfs(&board, lenr - 1, i)
	}

	for i := 0; i < lenr; i ++ {
		for j := 0; j < lenc; j ++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board *[][]byte, sr, sc int) {
	if sr < 0 || sc < 0 || sc > lenc - 1 || sr > lenr - 1 || (*board)[sr][sc] != 'O'{
		return
	}

	(*board)[sr][sc] = 'A'

	for i := 0; i < 4; i++ {
		nsr := sr + direction[i][0]
		nsc := sc + direction[i][1]
		dfs(board, nsr, nsc)
	}
}
