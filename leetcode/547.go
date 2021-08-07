package main

import "fmt"

// 并查集
type UnionFind struct {
	parent map[int]int
	rank   map[int]int
	count  int
}

// 按秩合并
func (u *UnionFind) Merge(x, y int) {
	rootX, rootY := u.Find(x), u.Find(y)

	if rootX != rootY {
		if u.rank[rootX] == u.rank[rootY] {
			u.rank[rootY]++
		} else if u.rank[rootX] > u.rank[rootY] {
			rootX, rootY = rootY, rootX
		}
		u.parent[rootX] = rootY

		u.count--
	}
}

// 路径压缩
func (u *UnionFind) Find(x int) int {
	r := x

	for u.parent[r] != r {
		r = u.parent[r]
	}

	for x != u.parent[x] {
		temp := u.parent[x]
		u.parent[x] = r
		x = temp
	}

	return r
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := &UnionFind{
		parent: make(map[int]int, n),
		rank:   make(map[int]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}

	for index, province := range isConnected {
		for k, status := range province {
			if status == 1 {
				uf.Merge(k, index)
			}
		}
	}

	return uf.count
}

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}