package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 只能说 离谱哦
// 应该是可以优化的

type refInfo struct {
	index int
	num   int
}

type unionFind struct {
	ref    map[string]refInfo
	parent map[int]int
	rank   map[int]int
	count  int
}

func (u *unionFind) init(n int) {
	u.ref = make(map[string]refInfo)
	u.parent = make(map[int]int)
	u.rank = make(map[int]int)
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.rank[i] = 1
	}
	u.count = n
}

// 路径压缩
func (u *unionFind) find(p int) int {
	if u.parent[p] != p {
		u.parent[p] = u.find(u.parent[p])
	}

	return u.parent[p]
}

// 按秩合并
// 两个一起使用有可能破坏秩的准确性
func (u *unionFind) union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)

	if rootP == rootQ {
		return
	}

	// 把简单的树往复杂的树上合并
	if u.rank[rootP] < u.rank[rootQ] {
		u.parent[rootP] = rootQ
	} else if u.rank[rootP] > u.rank[rootQ] {
		u.parent[rootQ] = rootP
	} else {
		// 相等时，新树深度增加
		u.parent[rootQ] = rootP
		u.rank[rootP]++
	}

	u.count--
}

func main() {
	fmt.Println(trulyMostPopular([]string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"}, []string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"}))
}

func trulyMostPopular(names []string, synonyms []string) []string {
	var u unionFind
	lenn := len(names)
	u.init(lenn)
	i := 0
	for i, _ = range names {
		a := strings.Index(names[i], "(")
		b := strings.Index(names[i], ")")
		num, _ := strconv.Atoi(names[i][a+1 : b])
		u.ref[names[i][:a]] = refInfo{index: i, num: num}
	}

	for _, v := range synonyms {
		a := strings.Index(v, "(")
		c := strings.Index(v, ")")
		b := strings.Index(v, ",")
		u.union(u.ref[v[a+1:b]].index, u.ref[v[b+1:c]].index)
	}

	res := make([]string, u.count)

	resm := make(map[int]string)
	resnum := make(map[int]int)

	for _, v := range u.ref {
		p := u.find(v.index)
		if max, ok := resm[p]; !ok {
			resm[p] = names[v.index]
		} else {
			if max > names[v.index] {
				resm[p] = names[v.index]
			}
		}
		resnum[p] += v.num
	}

	index := 0
	for k, v := range resm {
		a := strings.Index(v, "(")
		res[index] = v[:a] + "(" + strconv.Itoa(resnum[k]) + ")"
		index++
	}

	return res
}
