package main

import "fmt"

func main() {
	fmt.Println(colorBorder([][]int{{1,1,1},{1,1,1},{1,1,1}},1,1,2))
}

var direction = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var visted [][]bool
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	visted = make([][]bool, len(grid))
	for k := range visted {
		visted[k] = make([]bool, len(grid[0]))
	}

	dfs(grid, row, col, color)

	return grid
}

func dfs(grid [][]int, r, c int, color int) {
	visted[r][c] = true
	now := grid[r][c]
	flag := false
	//grid[r][c] = color
	for i := 0; i < 4; i ++ {
		nx, ny := r + direction[i][0], c + direction[i][1]
		if nx < len(grid) && nx >= 0 && ny >= 0 && ny < len(grid[0]) {
			if grid[nx][ny] == now {
				if !visted[nx][ny] {
					dfs(grid, nx, ny, color)
				}
			} else if grid[nx][ny] != color {
				flag = true
			}
		} else {
			flag = true
		}
	}
	if flag {
		grid[r][c] = color
	}
}
