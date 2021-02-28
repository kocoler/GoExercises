package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(accountsMerge([][]string{
		{"Alex","Alex5@m.co","Alex4@m.co","Alex0@m.co"},
		{"Ethan","Ethan3@m.co","Ethan3@m.co","Ethan0@m.co"},
		{"Kevin","Kevin4@m.co","Kevin2@m.co","Kevin2@m.co"},
		{"Gabe","Gabe0@m.co","Gabe3@m.co","Gabe2@m.co"},
		{"Gabe","Gabe3@m.co","Gabe4@m.co","Gabe2@m.co"},
	}))
}

var mrecord = map[string]string{}
var mparent = map[string]int{}
var parent [10000]int

func find(a int) int {
	if parent[a] != a {
		parent[a] = find(parent[a])
	}

	return parent[a]
}

func union(a, b int) {
	//roota := find(a)
	//rootb := find(b)
	//if roota != rootb {
	//	// 按秩合并
	//	if rank[rootx] < rank[rooty] {
	//		swap(rootx, rooty);
	//	}
	//	parent[rooty] = rootx;
	//	count--;
	//	// 如果秩相等，将父节点rootx秩 + 1
	//	if (rank[rootx] == rank[rooty]) rank[rootx] += 1;
	//}
	parent[find(a)] = find(b)
}

func accountsMerge(accounts [][]string) [][]string {
	lena := len(accounts)
	if lena < 2 {
		return accounts
	}

	var res [][]string
	mrecord = make(map[string]string) // record name
	mparent = make(map[string]int) // record parent
	parent = [10000]int{} // parent
	count := 0
	for _, v := range accounts {
		name := v[0]
		lenv := len(v)
		for i := 1; i < lenv; i++ {
			if _, has := mparent[v[i]]; !has {
				mparent[v[i]] = count
				mrecord[v[i]] = name
			}
		}
		parent[count] = count
		count++
	}

	for _, account := range accounts {
		origin := mparent[account[1]]
		for _, email := range account[2:] {
			union(mparent[email], origin)
		}
	}

	temp := map[int][]string{}
	for email, index := range mparent {
		index = find(index)
		temp[index] = append(temp[index], email)
	}

	for _, emails := range temp {
		sort.Strings(emails)
		account := append([]string{mrecord[emails[0]]}, emails...)
		res = append(res, account)
	}

	return res
}
