package main

import "fmt"

func main() {
	fmt.Println(friendRequests(3, [][]int{{0, 1}}, [][]int{{0, 2}, {2, 1}}))
}

type UnionFind struct {
	parent map[int]int
	rank   map[int]int
	count  int
}

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

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	lenr := len(requests)
	res := make([]bool, lenr)
	uf := UnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}

	//restrictionsMap := make(map[int][]int)
	//for _, restriction := range restrictions {
	//	restrictionsMap[restriction[0]] = append(restrictionsMap[restriction[0]], restriction[1])
	//	restrictionsMap[restriction[1]] = append(restrictionsMap[restriction[1]], restriction[0])
	//}

	for k, request := range requests {
		p1 := uf.Find(request[1])
		p0 := uf.Find(request[0])

		flag := true
		// check 0 restriction
		for _, v := range restrictions {
			v0 := uf.Find(v[0])
			v1 := uf.Find(v[1])
			if (v1 == p1 && v0 == p0) || (v1 == p0 && v0 == p1) {
				flag = false
				break
			}
		}

		if flag {
			uf.Merge(request[0], request[1])
			res[k] = true
		}
	}

	return res
}
