package main

import "fmt"

var dis string
var src [][]byte

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	dis = word
	src = board
	res = false

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if res {
				return res
			}
			bt(i, j, 0)
		}
	}

	return res
}

var stepMap [100][100]bool
var res bool
func bt(x, y int, begin int) {
	if res {
		return
	}
	if !check(x, y) {
		return
	}
	if begin >= len(dis) {
		return
	}
	if src[x][y] != dis[begin] {
		return
	}
	stepMap[x][y] = true
	fmt.Println(x, y, begin)
	if begin == len(dis)-1 {
		res = true
		return
	}
	// left
	bt(x-1, y, begin+1)
	// right
	bt(x+1, y, begin+1)
	// up
	bt(x, y-1, begin+1)
	// down
	bt(x, y+1, begin+1)
	stepMap[x][y] = false
}

func check(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= len(src) || y >= len(src[0]) {
		return false
	}
	if stepMap[x][y] {
		return false
	}

	return true
}


// 9.16

var words []byte
var visted [][]bool
var lenc int
var lenr int
var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var bo [][]byte
var lenw int
func exist(board [][]byte, word string) bool {
	words = []byte(word)
	bo = board
	lenw = len(word)
	lenc = len(board)
	lenr = len(board[0])

	visted = make([][]bool, lenc)
	for i := 0; i < lenc; i ++ {
		visted[i] = make([]bool, lenr)
	}

	for i := 0; i < lenc; i ++ {
		for j := 0; j < lenr; j ++ {
			if bo[i][j] == words[0] {
				visted[i][j] = true
				if dfs(i, j, 1) {
					return true
				}
				visted[i][j] = false
			}

		}
	}

	return false
}

func dfs(sc, sr int, index int) bool {
	if index == lenw {
		return true
	}

	for i := 0; i < 4; i++ {
		nsr := sr + dis[i][0]
		nsc := sc + dis[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && !visted[nsc][nsr] && words[index] == bo[nsc][nsr] {
			visted[nsc][nsr] = true
			find := dfs(nsc, nsr, index + 1)
			if find {
				return find
			}
			// rowback
			visted[nsc][nsr] = false
		}
	}

	return false
}
