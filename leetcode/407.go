package main

// bfs
// 还是要看 单调栈，比较好奇的
type pair struct {
	r int
	c int
}

func trapRainWater(heightMap [][]int) int {
	lenr := len(heightMap)
	lenc := len(heightMap[0])

	maxHeight := 0
	for i := range heightMap {
		for j := range heightMap[i] {
			maxHeight = max(maxHeight, heightMap[i][j])
		}
	}

	water := make([][]int, lenr)
	for i := range water {
		water[i] = make([]int, lenc)
		for j := range water[i] {
			water[i][j] = maxHeight
		}
	}

	// 构建队列
	queue := []pair{}
	for i := range heightMap {
		for j := range heightMap[i] {
			if i == 0 || j == 0 || i == lenr - 1 || j == lenc - 1 {
				if water[i][j] > heightMap[i][j] {
					water[i][j] = heightMap[i][j]
					queue = append(queue, pair{i, j})
				}
			}
		}
	}

	// 开始搜索
	dist := []int{-1, 0, 1, 0, -1}
	for len(queue) != 0 {
		now := queue[0]
		queue = queue[1:]
		r := now.r
		c := now.c
		for i := 0; i < 4; i ++ {
			nr := r + dist[i]
			nc := c + dist[i + 1]
			if nr >= 0 && nc >= 0 && nr < lenr && nc < lenc && water[r][c] < water[nr][nc] && water[nr][nc] > heightMap[nr][nc] {
				water[nr][nc] = max(water[r][c], heightMap[nr][nc])
				queue = append(queue, pair{nr, nc})
			}
		}
	}

	res := 0
	for i := range water {
		for j := range water[i] {
			res += water[i][j] - heightMap[i][j]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

