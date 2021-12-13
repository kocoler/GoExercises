package main

import "fmt"

// 第一次用 并查集 解决图的搜索问题
// 第一次写其实还蛮困难的
// 面试的时候还是用 DFS 吧

type unionFind struct {
	parent []int
	rank   []int // uf size as rank
	count  int
}

func (uf *unionFind) init(size int) {
	uf.parent = make([]int, size)
	uf.rank = make([]int, size)
	uf.count = size
	for k := range uf.parent {
		uf.parent[k] = k
		uf.rank[k] = 1
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	xp, yp := uf.find(x), uf.find(y)
	//fmt.Println("merge", x, y, xp, yp)
	if xp == yp {
		return
	}

	uf.count--

	if uf.rank[xp] >= uf.rank[yp] {
		// merge simpler/equal yp to xp
		uf.parent[yp] = xp
		uf.rank[xp] += uf.rank[yp]
	} else {
		uf.parent[xp] = yp
		uf.rank[yp] += uf.rank[xp]
	}
}

func main() {
	fmt.Println(closedIsland([][]int{{0, 0, 1, 1, 0, 1, 0, 0, 1, 0}, {1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, {1, 0, 1, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 1, 0, 0, 0, 0, 1, 0, 1}, {0, 0, 0, 0, 0, 0, 1, 1, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 1, 1, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 1, 1, 0, 1, 1, 0, 1, 1, 0}}))
}

func closedIsland(grid [][]int) int {
	dis := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var res int
	var uf unionFind

	maxr, maxc := len(grid), len(grid[0])
	uf.init(maxr*maxc + 1)

	visited := make([][]bool, maxr)
	for k := range visited {
		visited[k] = make([]bool, maxc)
	}

	var dfs func(int, int, int)
	dfs = func(r, c int, p int) {
		uf.union(r*maxc+c, p)
		if r == 0 || c == 0 || r == maxr-1 || c == maxc-1 {
			uf.union(p, maxr*maxc)
		}
		for i := 0; i < 4; i++ {
			nxr, nxc := r+dis[i][0], c+dis[i][1]
			if nxr >= 0 && nxc >= 0 && nxr < maxr && nxc < maxc {
				if !visited[nxr][nxc] && grid[nxr][nxc] == grid[r][c] {
					visited[nxr][nxc] = true
					dfs(nxr, nxc, p)
				}
			}
		}
	}

	for i := 0; i < maxr; i++ {
		for j := 0; j < maxc; j++ {
			if !visited[i][j] {
				visited[i][j] = true
				if grid[i][j] == 1 {
					uf.union(i*maxc+j, maxr*maxc)
				} else {
					dfs(i, j, i*maxc+j)
				}
			}
		}
	}

	res = uf.count

	return res - 1
}
