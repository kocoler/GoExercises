package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2))
}

var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var pimage [][]int
var res [][]int
var newc int
var lenm int
var lenn int
var visited [][]bool

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	lenm = len(image)
	lenn = len(image[0])
	pimage = make([][]int, lenm)
	res = make([][]int, lenm)
	visited = make([][]bool, lenm)

	for i := 0; i < lenm; i++ {
		pimage[i] = make([]int, lenn)
		copy(pimage[i], image[i])
		res[i] = make([]int, lenn)
		copy(res[i], image[i])
		visited[i] = make([]bool, lenn)
	}

	newc = newColor

	res[sr][sc] = newc

	dfs(sr, sc)

	return res
}

func dfs(sr int, sc int) {
	if visited[sr][sc] {
		return
	}
	visited[sr][sc] = true
	for i := 0; i < 4; i++ {
		nsr := sr + dis[i][0]
		nsc := sc + dis[i][1]
		fmt.Println(nsr, nsc, lenn, lenm)
		if nsr > -1 && nsc > -1 && nsr < lenm && nsc < lenn && pimage[sr][sc] == pimage[nsr][nsc] {
			res[nsr][nsc] = newc
			dfs(nsr, nsc)
		}
	}
}
