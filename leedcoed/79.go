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
