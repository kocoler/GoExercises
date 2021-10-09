package main

import (
	"fmt"
	"strings"
)

// backtrace
// 注意一下 check 的条件
// 两个对角线如何在运行的时候标记好

func main() {
	fmt.Println(solveNQueens(4))
}

var res [][]string
var N int
func solveNQueens(n int) [][]string {
	res = [][]string{}
	N = n
	chess = make([]string, n)
	col = make([]bool, n)
	diagonal = make([]bool, 2*n)
	diagonal2 = make([]bool, 2*n)
	for i := 0; i < n; i ++ {
		chess[i] = strings.Repeat(".", n)
	}

	backtrace(0)

	return res
}

var chess []string
var col []bool
var diagonal []bool
var diagonal2 []bool
func backtrace(n int) {
	if n == N {
		temp := make([]string, N)
		copy(temp, chess)
		res = append(res, temp)
		return
	}

	s := []byte(chess[n])
	for i := 0; i < N; i ++ {
		if check(n, i) {
			s[i] = 'Q'
			chess[n] = string(s)
			col[i] = true
			fmt.Println(i, n, i-n+N-1, n-i+N-1)
			diagonal[i-n+N-1] = true
			diagonal2[n+i] = true

			backtrace(n+1)

			// backtrace
			col[i] = false
			diagonal[i-n+N-1] = false
			diagonal2[n+i] = false

			s[i] = '.'
			chess[n] = string(s)
		}
	}
}

func check(n, i int) bool {
	if col[i] {
		return false
	}
	if diagonal[i-n+N-1] {
		return false
	}
	if diagonal2[n+i] {
		return false
	}

	return true
}

